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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcSecurityIntelligenceURLFeed(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_security_intelligence_url_feed.test", "name", "my_si_url_feed"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_security_intelligence_url_feed.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_security_intelligence_url_feed.test", "feed_url", "https://example.com/path/to/feed.txt"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_security_intelligence_url_feed.test", "checksum_url", "https://example.com/path/to/checksum.md5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_security_intelligence_url_feed.test", "update_frequency", "120"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSecurityIntelligenceURLFeedConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcSecurityIntelligenceURLFeedConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSecurityIntelligenceURLFeedConfig() string {
	config := `resource "fmc_security_intelligence_url_feed" "test" {` + "\n"
	config += `	name = "my_si_url_feed"` + "\n"
	config += `	feed_url = "https://example.com/path/to/feed.txt"` + "\n"
	config += `	checksum_url = "https://example.com/path/to/checksum.md5"` + "\n"
	config += `	update_frequency = 120` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_security_intelligence_url_feed" "test" {
			id = fmc_security_intelligence_url_feed.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcSecurityIntelligenceURLFeedConfig() string {
	config := `resource "fmc_security_intelligence_url_feed" "test" {` + "\n"
	config += `	name = "my_si_url_feed"` + "\n"
	config += `	feed_url = "https://example.com/path/to/feed.txt"` + "\n"
	config += `	checksum_url = "https://example.com/path/to/checksum.md5"` + "\n"
	config += `	update_frequency = 120` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_security_intelligence_url_feed" "test" {
			name = fmc_security_intelligence_url_feed.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
