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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcPolicyList(t *testing.T) {
	if os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_list.test", "name", "my_policy_list"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_policy_list.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_list.test", "action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_list.test", "match_community_exactly", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_list.test", "metric", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_list.test", "tag", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcPolicyListPrerequisitesConfig + testAccDataSourceFmcPolicyListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcPolicyListPrerequisitesConfig + testAccNamedDataSourceFmcPolicyListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcPolicyListPrerequisitesConfig = `
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name

resource "fmc_security_zone" "test" {
  name           = "my_policy_list_security_zone"
  interface_type = "ROUTED"
}

resource "fmc_standard_acl" "test" {
  name        = "my_policy_list_std_acl"
  entries = [
    { action = "DENY", literals = [{ value = "10.1.1.0/24" }] }
  ]
}

resource "fmc_as_path" "test" {
  name = 263
  entries = [
    { action = "PERMIT", regular_expression = "107" },
    { action = "PERMIT", regular_expression = "108" },
  ]
}

resource "fmc_standard_community_list" "test" {
  name = "my_policy_list_std_community_list"
  entries = [{
    action       = "PERMIT"
    communities  = "100 200 300"
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcPolicyListConfig() string {
	config := `resource "fmc_policy_list" "test" {` + "\n"
	config += `	name = "my_policy_list"` + "\n"
	config += `	action = "PERMIT"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = fmc_security_zone.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	interface_names = [var.interface_name]` + "\n"
	config += `	address_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	next_hop_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	route_source_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	as_paths = [{` + "\n"
	config += `		id = fmc_as_path.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	community_lists = [{` + "\n"
	config += `		id = fmc_standard_community_list.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	match_community_exactly = true` + "\n"
	config += `	metric = 100` + "\n"
	config += `	tag = 100` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_list" "test" {
			id = fmc_policy_list.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcPolicyListConfig() string {
	config := `resource "fmc_policy_list" "test" {` + "\n"
	config += `	name = "my_policy_list"` + "\n"
	config += `	action = "PERMIT"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = fmc_security_zone.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	interface_names = [var.interface_name]` + "\n"
	config += `	address_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	next_hop_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	route_source_standard_access_lists = [{` + "\n"
	config += `		id = fmc_standard_acl.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	as_paths = [{` + "\n"
	config += `		id = fmc_as_path.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	community_lists = [{` + "\n"
	config += `		id = fmc_standard_community_list.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	match_community_exactly = true` + "\n"
	config += `	metric = 100` + "\n"
	config += `	tag = 100` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_list" "test" {
			name = fmc_policy_list.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
