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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcVPNRASecureClientCustomization(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_secure_client_customization.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "gui_and_text_messages.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "icons_and_images.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "scripts.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "binaries.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "custom_installer_transforms.0.id", "12345678-1234-1234-1234-123456789012"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_secure_client_customization.test", "localized_installer_transforms.0.id", "12345678-1234-1234-1234-123456789012"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNRASecureClientCustomizationConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRASecureClientCustomizationConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcVPNRASecureClientCustomizationConfig_minimum() string {
	config := `resource "fmc_vpn_ra_secure_client_customization" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRASecureClientCustomizationConfig_all() string {
	config := `resource "fmc_vpn_ra_secure_client_customization" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	gui_and_text_messages = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `	icons_and_images = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `	scripts = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `	binaries = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `	custom_installer_transforms = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `	localized_installer_transforms = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456789012"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
