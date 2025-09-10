// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-fmc"
)

// FmcProvider defines the provider implementation.
type FmcProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FmcProviderModel describes the provider data model.
type FmcProviderModel struct {
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	Token      types.String `tfsdk:"token"`
	URL        types.String `tfsdk:"url"`
	Insecure   types.Bool   `tfsdk:"insecure"`
	ReqTimeout types.String `tfsdk:"req_timeout"`
	Retries    types.Int64  `tfsdk:"retries"`
}

// FmcProviderData describes the data maintained by the provider.
type FmcProviderData struct {
	Client *fmc.Client
}

// Define provider constants
const (
	// maximum elements in single create bulk request
	bulkSizeCreate int = 1000
	// maximum payload size in bytes
	maxPayloadSize int = 2048000
	// maximum URL Param length. This is a rough estimate and does not account for the entire URL length.
	maxUrlParamLength int = 7000
)

// Metadata returns the provider type name.
func (p *FmcProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fmc"
	resp.Version = p.version
}

func (p *FmcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the FMC instance. This can also be set as the FMC_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the FMC instance. This can also be set as the FMC_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "API token for cdFMC instance. This can also be set as the FMC_TOKEN environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco FMC/cdFMC instance or SCC Firewall Manager Base URI (https://api.X.security.cisco.com/firewall) . This can also be set as the FMC_URL environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the FMC_INSECURE environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"req_timeout": schema.StringAttribute{
				MarkdownDescription: "Timeout for a single HTTPS request made to REST API before it is retried. This can also be set as the FMC_REQTIMEOUT environment variable. A string like `\"1s\"` means one second. Defaults to `\"5s\"`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the FMC_RETRIES environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
		},
	}
}

func (p *FmcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config FmcProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//// User must provide (username and password) or (token) to the provider
	// Get username
	var username string
	if config.Username.IsNull() {
		username = os.Getenv("FMC_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	// Get password
	var password string
	if config.Password.IsNull() {
		password = os.Getenv("FMC_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	// Get token
	var token string
	if config.Token.IsNull() {
		token = os.Getenv("FMC_TOKEN")
	} else {
		token = config.Token.ValueString()
	}

	// Fail if the user didn't provider any credentials
	if password == "" && token == "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Please provide credentials for FMC (username and password) or cdFMC (token)",
		)
		return
	}

	// Fail if the user provided credentials for both FMC and cdFMC
	if password != "" && token != "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use both password and token at the same time",
		)
		return
	}

	// Fail if the user provided `password` but not `username`
	if password != "" && username == "" {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Username and password need to be provided for FMC",
		)
		return
	}

	// User must provide a URL to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("FMC_URL")
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	// Check if URL points to Security Cloud Control (SCC) base URI
	scc, err := regexp.MatchString("^https://api\\.\\S+?\\.security\\.cisco\\.com/firewall$", url)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to parse URL",
			fmt.Sprintf("Cannot parse the URL: %v", err),
		)
		return
	}

	// If it is a Security Cloud Control API router, append the path
	if scc {
		url += "/v1/cdfmc"
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("FMC_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var reqTimeout time.Duration
	if config.ReqTimeout.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as req_timeout",
		)
		return
	}

	var reqTimeoutStr string
	if config.ReqTimeout.IsNull() {
		reqTimeoutStr = os.Getenv("FMC_REQTIMEOUT")
		if reqTimeoutStr == "" {
			reqTimeoutStr = "5s"
		}
	} else {
		reqTimeoutStr = config.ReqTimeout.ValueString()
	}
	reqTimeout, err = time.ParseDuration(reqTimeoutStr)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			fmt.Sprintf("Cannot parse the req_timeout string: %v", err),
		)
		return
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("FMC_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	tflog.Debug(ctx, fmt.Sprint("Creating a new FMC client",
		"  url=", url,
		"  insecure=", insecure,
		"  req_timeout=", reqTimeout,
		"  retries=", retries,
	))

	// Create a new FMC or cdFMC client and set it to the provider client
	var c fmc.Client
	if password != "" {
		c, err = fmc.NewClient(url, username, password, fmc.Insecure(insecure), fmc.MaxRetries(int(retries)), fmc.RequestTimeout(reqTimeout))
	} else if token != "" {
		c, err = fmc.NewClientCDFMC(url, token, fmc.Insecure(insecure), fmc.MaxRetries(int(retries)), fmc.RequestTimeout(reqTimeout))
	} else {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Failed to determine target device type (FMC/cdFMC)",
		)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create fmc client:\n\n"+err.Error(),
		)
		return
	}

	data := FmcProviderData{Client: &c}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *FmcProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAccessCategoryResource,
		NewAccessControlPolicyResource,
		NewAccessRuleResource,
		NewAccessRulesResource,
		NewApplicationFilterResource,
		NewApplicationFiltersResource,
		NewASPathResource,
		NewASPathsResource,
		NewBFDTemplateResource,
		NewCertificateEnrollmentResource,
		NewCertificateMapResource,
		NewCertificateMapsResource,
		NewChassisResource,
		NewChassisEtherChannelInterfaceResource,
		NewChassisLogicalDeviceResource,
		NewChassisPhysicalInterfaceResource,
		NewChassisSubinterfaceResource,
		NewDeviceResource,
		NewDeviceBFDResource,
		NewDeviceBGPResource,
		NewDeviceBGPGeneralSettingsResource,
		NewDeviceBridgeGroupInterfaceResource,
		NewDeviceClusterResource,
		NewDeviceClusterHealthMonitorResource,
		NewDeviceDeployResource,
		NewDeviceEtherChannelInterfaceResource,
		NewDeviceHAPairResource,
		NewDeviceHAPairFailoverInterfaceMACAddressResource,
		NewDeviceHAPairMonitoringResource,
		NewDeviceIPv4StaticRouteResource,
		NewDeviceIPv6StaticRouteResource,
		NewDeviceLoopbackInterfaceResource,
		NewDevicePhysicalInterfaceResource,
		NewDeviceSubinterfaceResource,
		NewDeviceVNIInterfaceResource,
		NewDeviceVRFResource,
		NewDeviceVRFIPv4StaticRouteResource,
		NewDeviceVRFIPv6StaticRouteResource,
		NewDeviceVTEPPolicyResource,
		NewDeviceVTIInterfaceResource,
		NewDynamicAccessPolicyResource,
		NewDynamicObjectsResource,
		NewExpandedCommunityListResource,
		NewExpandedCommunityListsResource,
		NewExtendedACLResource,
		NewExtendedCommunityListResource,
		NewExtendedCommunityListsResource,
		NewExternalCertificateResource,
		NewFilePolicyResource,
		NewFQDNObjectResource,
		NewFQDNObjectsResource,
		NewFTDAutoNATRuleResource,
		NewFTDManualNATRuleResource,
		NewFTDNATPolicyResource,
		NewGeolocationResource,
		NewGroupPolicyResource,
		NewHostResource,
		NewHostsResource,
		NewICMPv4ObjectResource,
		NewICMPv4ObjectsResource,
		NewICMPv6ObjectResource,
		NewICMPv6ObjectsResource,
		NewIKEv1IPsecProposalsResource,
		NewIKEv1PoliciesResource,
		NewIKEv2IPsecProposalsResource,
		NewIKEv2PoliciesResource,
		NewInterfaceGroupResource,
		NewInternalCertificateResource,
		NewInternalCertificateAuthorityResource,
		NewIntrusionPolicyResource,
		NewIPv4AddressPoolResource,
		NewIPv4AddressPoolsResource,
		NewIPv4PrefixListResource,
		NewIPv4PrefixListsResource,
		NewIPv6AddressPoolResource,
		NewIPv6AddressPoolsResource,
		NewIPv6PrefixListResource,
		NewIPv6PrefixListsResource,
		NewNetworkResource,
		NewNetworkAnalysisPolicyResource,
		NewNetworkGroupResource,
		NewNetworkGroupsResource,
		NewNetworksResource,
		NewPolicyAssignmentResource,
		NewPolicyListResource,
		NewPortResource,
		NewPortGroupResource,
		NewPortGroupsResource,
		NewPortsResource,
		NewPrefilterPolicyResource,
		NewRadiusServerGroupResource,
		NewRangeResource,
		NewRangesResource,
		NewRealmADLDAPResource,
		NewRealmLocalResource,
		NewResourceProfilesResource,
		NewRouteMapResource,
		NewSecureClientCustomAttributeResource,
		NewSecureClientCustomizationResource,
		NewSecureClientExternalBrowserPackageResource,
		NewSecureClientImageResource,
		NewSecureClientPosturePackageResource,
		NewSecureClientProfileResource,
		NewSecurityIntelligenceNetworkFeedResource,
		NewSecurityIntelligenceNetworkFeedsResource,
		NewSecurityIntelligenceURLFeedResource,
		NewSecurityIntelligenceURLFeedsResource,
		NewSecurityZoneResource,
		NewSecurityZonesResource,
		NewServiceAccessResource,
		NewSGTResource,
		NewSGTsResource,
		NewSingleSignOnServerResource,
		NewSLAMonitorResource,
		NewSLAMonitorsResource,
		NewSmartLicenseResource,
		NewStandardACLResource,
		NewStandardCommunityListResource,
		NewStandardCommunityListsResource,
		NewTimeRangeResource,
		NewTimeRangesResource,
		NewTrustedCertificateAuthorityResource,
		NewTunnelZoneResource,
		NewTunnelZonesResource,
		NewURLResource,
		NewURLGroupResource,
		NewURLGroupsResource,
		NewURLsResource,
		NewVLANTagResource,
		NewVLANTagGroupResource,
		NewVLANTagGroupsResource,
		NewVLANTagsResource,
		NewVPNRAResource,
		NewVPNRAAddressAssignmentPolicyResource,
		NewVPNRACertificateMapResource,
		NewVPNRAConnectionProfilesResource,
		NewVPNRAIPSecCryptoMapResource,
		NewVPNRAIPSecIKEParametersResource,
		NewVPNRALDAPAttributeMapResource,
		NewVPNRALoadBalancingResource,
		NewVPNRASecureClientCustomizationResource,
		NewVPNS2SResource,
		NewVPNS2SAdvancedSettingsResource,
		NewVPNS2SEndpointsResource,
		NewVPNS2SIKESettingsResource,
		NewVPNS2SIPSECSettingsResource,
	}
}

func (p *FmcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccessCategoryDataSource,
		NewAccessControlPolicyDataSource,
		NewAccessRuleDataSource,
		NewApplicationDataSource,
		NewApplicationBusinessRelevanceDataSource,
		NewApplicationBusinessRelevancesDataSource,
		NewApplicationCategoriesDataSource,
		NewApplicationCategoryDataSource,
		NewApplicationFilterDataSource,
		NewApplicationFiltersDataSource,
		NewApplicationRiskDataSource,
		NewApplicationRisksDataSource,
		NewApplicationTagDataSource,
		NewApplicationTagsDataSource,
		NewApplicationTypeDataSource,
		NewApplicationTypesDataSource,
		NewApplicationsDataSource,
		NewASPathDataSource,
		NewASPathsDataSource,
		NewBFDTemplateDataSource,
		NewCertificateEnrollmentDataSource,
		NewCertificateMapDataSource,
		NewCertificateMapsDataSource,
		NewChassisDataSource,
		NewChassisEtherChannelInterfaceDataSource,
		NewChassisLogicalDeviceDataSource,
		NewChassisPhysicalInterfaceDataSource,
		NewChassisSubinterfaceDataSource,
		NewContinentsDataSource,
		NewCountriesDataSource,
		NewDeviceDataSource,
		NewDeviceBFDDataSource,
		NewDeviceBGPDataSource,
		NewDeviceBGPGeneralSettingsDataSource,
		NewDeviceBridgeGroupInterfaceDataSource,
		NewDeviceClusterDataSource,
		NewDeviceClusterHealthMonitorDataSource,
		NewDeviceEtherChannelInterfaceDataSource,
		NewDeviceHAPairDataSource,
		NewDeviceHAPairFailoverInterfaceMACAddressDataSource,
		NewDeviceHAPairMonitoringDataSource,
		NewDeviceIPv4StaticRouteDataSource,
		NewDeviceIPv6StaticRouteDataSource,
		NewDeviceLoopbackInterfaceDataSource,
		NewDevicePhysicalInterfaceDataSource,
		NewDeviceSubinterfaceDataSource,
		NewDeviceVNIInterfaceDataSource,
		NewDeviceVRFDataSource,
		NewDeviceVRFIPv4StaticRouteDataSource,
		NewDeviceVRFIPv6StaticRouteDataSource,
		NewDeviceVTEPPolicyDataSource,
		NewDeviceVTIInterfaceDataSource,
		NewDomainsDataSource,
		NewDynamicAccessPolicyDataSource,
		NewDynamicObjectsDataSource,
		NewEndpointDeviceTypesDataSource,
		NewExpandedCommunityListDataSource,
		NewExpandedCommunityListsDataSource,
		NewExtendedACLDataSource,
		NewExtendedCommunityListDataSource,
		NewExtendedCommunityListsDataSource,
		NewExternalCertificateDataSource,
		NewFileCategoriesDataSource,
		NewFileCategoryDataSource,
		NewFilePolicyDataSource,
		NewFileTypeDataSource,
		NewFileTypesDataSource,
		NewFQDNObjectDataSource,
		NewFQDNObjectsDataSource,
		NewFTDAutoNATRuleDataSource,
		NewFTDManualNATRuleDataSource,
		NewFTDNATPolicyDataSource,
		NewGeolocationDataSource,
		NewGroupPolicyDataSource,
		NewHostDataSource,
		NewHostsDataSource,
		NewICMPv4ObjectDataSource,
		NewICMPv4ObjectsDataSource,
		NewICMPv6ObjectDataSource,
		NewICMPv6ObjectsDataSource,
		NewIKEv1IPsecProposalsDataSource,
		NewIKEv1PoliciesDataSource,
		NewIKEv2IPsecProposalsDataSource,
		NewIKEv2PoliciesDataSource,
		NewInterfaceGroupDataSource,
		NewInternalCertificateDataSource,
		NewInternalCertificateAuthorityDataSource,
		NewIntrusionPolicyDataSource,
		NewIPv4AddressPoolDataSource,
		NewIPv4AddressPoolsDataSource,
		NewIPv4PrefixListDataSource,
		NewIPv4PrefixListsDataSource,
		NewIPv6AddressPoolDataSource,
		NewIPv6AddressPoolsDataSource,
		NewIPv6PrefixListDataSource,
		NewIPv6PrefixListsDataSource,
		NewISESGTsDataSource,
		NewNetworkDataSource,
		NewNetworkAnalysisPolicyDataSource,
		NewNetworkGroupDataSource,
		NewNetworksDataSource,
		NewPolicyAssignmentDataSource,
		NewPolicyListDataSource,
		NewPortDataSource,
		NewPortGroupDataSource,
		NewPortGroupsDataSource,
		NewPortsDataSource,
		NewPrefilterPolicyDataSource,
		NewRadiusServerGroupDataSource,
		NewRangeDataSource,
		NewRangesDataSource,
		NewRealmADLDAPDataSource,
		NewRealmLocalDataSource,
		NewResourceProfilesDataSource,
		NewRouteMapDataSource,
		NewSecureClientCustomAttributeDataSource,
		NewSecureClientCustomizationDataSource,
		NewSecureClientExternalBrowserPackageDataSource,
		NewSecureClientImageDataSource,
		NewSecureClientPosturePackageDataSource,
		NewSecureClientProfileDataSource,
		NewSecurityIntelligenceDNSFeedDataSource,
		NewSecurityIntelligenceDNSFeedsDataSource,
		NewSecurityIntelligenceDNSListDataSource,
		NewSecurityIntelligenceDNSListsDataSource,
		NewSecurityIntelligenceNetworkFeedDataSource,
		NewSecurityIntelligenceNetworkFeedsDataSource,
		NewSecurityIntelligenceNetworkListDataSource,
		NewSecurityIntelligenceNetworkListsDataSource,
		NewSecurityIntelligenceURLFeedDataSource,
		NewSecurityIntelligenceURLFeedsDataSource,
		NewSecurityIntelligenceUrlListDataSource,
		NewSecurityIntelligenceUrlListsDataSource,
		NewSecurityZoneDataSource,
		NewSecurityZonesDataSource,
		NewServiceAccessDataSource,
		NewSGTDataSource,
		NewSGTsDataSource,
		NewSingleSignOnServerDataSource,
		NewSLAMonitorDataSource,
		NewSLAMonitorsDataSource,
		NewSNMPAlertDataSource,
		NewSNMPAlertsDataSource,
		NewStandardACLDataSource,
		NewStandardCommunityListDataSource,
		NewStandardCommunityListsDataSource,
		NewSyslogAlertDataSource,
		NewSyslogAlertsDataSource,
		NewTimeRangeDataSource,
		NewTimeRangesDataSource,
		NewTrustedCertificateAuthorityDataSource,
		NewTunnelZoneDataSource,
		NewTunnelZonesDataSource,
		NewURLDataSource,
		NewURLGroupDataSource,
		NewURLGroupsDataSource,
		NewURLsDataSource,
		NewVariableSetDataSource,
		NewVLANTagDataSource,
		NewVLANTagGroupDataSource,
		NewVLANTagGroupsDataSource,
		NewVLANTagsDataSource,
		NewVPNRADataSource,
		NewVPNRAAddressAssignmentPolicyDataSource,
		NewVPNRACertificateMapDataSource,
		NewVPNRAConnectionProfilesDataSource,
		NewVPNRAIPSecCryptoMapDataSource,
		NewVPNRAIPSecIKEParametersDataSource,
		NewVPNRALDAPAttributeMapDataSource,
		NewVPNRALoadBalancingDataSource,
		NewVPNRASecureClientCustomizationDataSource,
		NewVPNS2SDataSource,
		NewVPNS2SAdvancedSettingsDataSource,
		NewVPNS2SEndpointsDataSource,
		NewVPNS2SIKESettingsDataSource,
		NewVPNS2SIPSECSettingsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FmcProvider{
			version: version,
		}
	}
}

// End of section. //template:end provider
