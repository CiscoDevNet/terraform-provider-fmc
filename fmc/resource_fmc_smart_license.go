package fmc

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var smartLicenseType string = "SmartLicense"

func resourceFmcSmartLicense() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Smart License in FMC\n" +
			"\n" +
			"## Example\n" +
			"Examples are shown below: \n\n" +
			"Enable Evaluation Mode\n" +
			"```hcl\n" +
			"resource \"fmc_smart_license\" \"license\" {\n" +
			"  registration_type = \"EVALUATION\"\n" +
			"}\n" +
			"```\n\n" +
			"Runs successfully as long as the registration status is REGISTERED\n" +
			"```hcl\n" +
			"resource \"fmc_smart_license\" \"license\" {\n" +
			"  registration_type = \"REGISTER\"\n" +
			"  token             = \"X2M3YmJlY...\"\n" +
			"}\n" +
			"```\n\n" +
			"Force to de-register and register with the token provided\n" +
			"```hcl\n" +
			"resource \"fmc_smart_license\" \"license\" {\n" +
			"  registration_type = \"REGISTER\"\n" +
			"  token             = \"X2M3YmJlY...\"\n" +
			"  force             = true\n" +
			"}\n" +
			"```\n\n" +
			"De-register on terraform destroy\n" +
			"```hcl\n" +
			"resource \"fmc_smart_license\" \"license\" {\n" +
			"  registration_type = \"REGISTER\"\n" +
			"  token             = \"X2M3YmJlY...\"\n" +
			"  retain            = false\n" +
			"}\n" +
			"```\n\n",
		CreateContext: resourceFmcSmartLicenseCreate,
		ReadContext:   resourceFmcSmartLicenseRead,
		UpdateContext: resourceFmcSmartLicenseUpdate,
		DeleteContext: resourceFmcSmartLicenseDelete,
		Schema: map[string]*schema.Schema{
			"registration_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of registration",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := strings.ToUpper(val.(string))
					// allowedValues := []string{"REGISTER", "DEREGISTER", "EVALUATION"}
					allowedValues := []string{"REGISTER", "EVALUATION"}
					for _, allowed := range allowedValues {
						if v == allowed {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be in %v, got: %q", key, allowedValues, v))
					return
				},
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Smart license token of the FMC",
				Sensitive:   true,
			},
			"registration_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"force": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Force to unregister and register",
			},
			"retain": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Set to false if you want to deregister on destroy",
			},
		},
	}
}

func resourceFmcSmartLicenseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics
	var err error

	status, err := c.GetFmcSmartLicense(ctx)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to get smart license status",
			Detail:   err.Error(),
		})
		return diags
	}

	if d.Get("registration_type") == "EVALUATION" && status.RegistrationStatus == "EVALUATION" {
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

		return resourceFmcSmartLicenseRead(ctx, d, m)
	}

	if d.Get("token") == "" {
		_, err = c.CreateFmcSmartLicense(ctx, &SmartLicenseRequest{
			Type:             smartLicenseType,
			RegistrationType: d.Get("registration_type").(string),
		})
	} else {
		_, err = c.CreateFmcSmartLicense(ctx, &SmartLicenseRequest{
			Type:             smartLicenseType,
			RegistrationType: d.Get("registration_type").(string),
			Token:            d.Get("token").(string),
		})
	}

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceFmcSmartLicenseRead(ctx, d, m)
}

func resourceFmcSmartLicenseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	item, err := c.GetFmcSmartLicense(ctx)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("registration_status", item.RegistrationStatus); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read smart license",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourceFmcSmartLicenseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics

	if d.Get("registration_type") == "REGISTER" && d.Get("token") == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "invalid resource block",
			Detail:   "token required when registration_type = \"REGISTER\"",
		})
		return diags
	}

	if d.HasChanges("registration_type", "token", "force") {
		if d.Get("force") == true {
			var err error

			status, err := c.GetFmcSmartLicense(ctx)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to get smart license status",
					Detail:   err.Error(),
				})
				return diags
			}
			if status.RegistrationStatus != "UNREGISTERED" {
				err = c.DeleteFmcSmartLicense(ctx)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "unable to update smart license",
						Detail:   err.Error(),
					})
					return diags
				}
			}

			if d.Get("token") == "" {
				_, err = c.CreateFmcSmartLicense(ctx, &SmartLicenseRequest{
					Type:             smartLicenseType,
					RegistrationType: d.Get("registration_type").(string),
				})
			} else {
				_, err = c.CreateFmcSmartLicense(ctx, &SmartLicenseRequest{
					Type:             smartLicenseType,
					RegistrationType: d.Get("registration_type").(string),
					Token:            d.Get("token").(string),
				})
			}

			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to update smart license",
					Detail:   err.Error(),
				})
				return diags
			}

			status, err = c.GetFmcSmartLicense(ctx)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to get smart license status",
					Detail:   err.Error(),
				})
				return diags
			}

			if status.RegistrationStatus == "UNREGISTERED" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "unable to update license status",
					Detail:   "the token might be invalid",
				})
				return diags
			}
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "smart license not working as expected?",
				Detail:   "try setting force = true",
			})
			return diags
		}
	}
	return resourceFmcSmartLicenseRead(ctx, d, m)
}

func resourceFmcSmartLicenseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.Get("retain") == false {
		err := c.DeleteFmcSmartLicense(ctx)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to delete smart license",
				Detail:   err.Error(),
			})
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
