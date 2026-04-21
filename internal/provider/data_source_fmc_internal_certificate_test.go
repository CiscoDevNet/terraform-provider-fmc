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

func TestAccDataSourceFmcInternalCertificate(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_internal_certificate.test", "name", "my_internal_certificate"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_internal_certificate.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcInternalCertificatePrerequisitesConfig + testAccDataSourceFmcInternalCertificateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcInternalCertificatePrerequisitesConfig + testAccNamedDataSourceFmcInternalCertificateConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcInternalCertificatePrerequisitesConfig = `
local {
  cert = chomp(<<EOT
-----BEGIN CERTIFICATE-----
MIIErzCCApegAwIBAgIUU6GBJ2MENn84UcCE7nOSN0g7swIwDQYJKoZIhvcNAQEL
BQAwXDELMAkGA1UEBhMCUEwxDDAKBgNVBAgMA01QTDEMMAoGA1UEBwwDS1JLMRAw
DgYDVQQKDAdDaHJhYmFzMQ0wCwYDVQQLDARsYWJzMRAwDgYDVQQDDAdjcm5zLnBs
MB4XDTIyMDUyMDEyMjkwNFoXDTI3MDQyNDEyMjkwNFowYzELMAkGA1UEBhMCUEwx
DDAKBgNVBAgMA01QTDEMMAoGA1UEBwwDS1JLMRAwDgYDVQQKDAdDaHJhYmFzMQ0w
CwYDVQQLDARsYWJzMRcwFQYDVQQDDA5mbWMtMDEuY3Jucy5wbDCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBAL/4XvPRdq54Lw7oDsP2KYKYKf3R9gdl2Gez
Mxn8SLisocMIIxI+utisIkeNBomBLJqPZvY9+QOsbM/jl9VmK/Lrlzr4cN5rlI4I
jCY4ut+X/XSYnojVqQ9Agc8NVuWl2yYMrIskdUsZDA7p/dB/mwrD3bGlTTkfvsOn
lpbR52MvwAF8qKPXzScZq5ibVjOwSoiJKmk3dcGkmgq03q0OJK2JHE4PUAiAhgU1
mS0tzyMCap+AKysk3k0GmV2O6vdBepAX600tNQoGyqzD9u1JCBMRESW31E0tKMTx
4KlD/CGDYTivxL+ysUqhS1lsE02Tdyo8X/k2fysax06uAjyePZsCAwEAAaNiMGAw
DAYDVR0TAQH/BAIwADAOBgNVHQ8BAf8EBAMCA4gwHQYDVR0lBBYwFAYIKwYBBQUH
AwEGCCsGAQUFBwMCMCEGA1UdEQQaMBiCDmZtYy0wMS5jcm5zLnBsggZmbWMtMDEw
DQYJKoZIhvcNAQELBQADggIBAG0oLwM0y8PBfiCTc4ElcPww+hFDeE8OsBgJPjhu
lh6mzhSZJEJh7SfqSRwDqMjVbS2LK6AubeY9lw52jKy9y8U7UBCg3k4GsNBVd0yV
6VQlJ76dIe7ksRELCc515jR+5HdoFg2CEEU/xhA/GLVpC6YFSfZseD95gRXPa/bm
d9rTuT17T3butji4i4DeS+BD+DzCC0vKVPZUlP3F+02Wv8MoyIxJzGVIhuHPHIAY
40xgvXna2onComp+0JMUmuWYNnBQgiZg5jdEL6EPXuGnrwDShZ9xV7nhXcx3vfyo
HCNOeRTPi6m0m4kxDrwqwgm8XFDFpetkcuVNC/QARLXcVL1gMRXi16Q2h1WmXSxX
XoaxyFXlvzzzW8fT4zL92H8F8c1FK/YjLQ4KucT9RXlrFbfp8jyaLhILfFwcJmXv
JERCDWZ0SD5HiK9tr+0ljGR4Zr4pApmz/0SvPXPMwSPwg1M5zUdUs2rDT+PdgQhs
PGauMnRNqxxhnq7A5gF3j+L0RJMymFX7d9bv8BKlUzmFB93F5qNKNTVfPbvDKXLM
m/LE3E4cEQq/O/nZOGkYwqAHu2FckS4a4rJgBdatDu8EjzStACIoDRXD0n3YFq83
lnkGIuzLdKP2P/6l4By5UzEWKxgbQF0Kt450jMH2pzLIxpMFKXXtXkLlJLmuBHUW
E+AT
-----END CERTIFICATE-----
EOT
)

  cert_key = chomp(<<EOT
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/+F7z0XaueC8O
6A7D9imCmCn90fYHZdhnszMZ/Ei4rKHDCCMSPrrYrCJHjQaJgSyaj2b2PfkDrGzP
45fVZivy65c6+HDea5SOCIwmOLrfl/10mJ6I1akPQIHPDVblpdsmDKyLJHVLGQwO
6f3Qf5sKw92xpU05H77Dp5aW0edjL8ABfKij180nGauYm1YzsEqIiSppN3XBpJoK
tN6tDiStiRxOD1AIgIYFNZktLc8jAmqfgCsrJN5NBpldjur3QXqQF+tNLTUKBsqs
w/btSQgTERElt9RNLSjE8eCpQ/whg2E4r8S/srFKoUtZbBNNk3cqPF/5Nn8rGsdO
rgI8nj2bAgMBAAECggEBAKCSfcnvsRfqi2IqlP8wzjPgV29pyiMR+1tyNxJLRgDC
1WOvULaYJe3KvbmfqpIlrEt0G6lo2PtLpJCXTI9BMQ22Jf87hB1nWxEa8S+fUCBq
n6SBbEoEfa62RF4OSFs6yf4fp8uTKVExmDZ5XsahlewBPVS0cc2QUh6R9ZId7S+2
dB2hq5/zdndyRIHJDrt5Xega9mz/g+k0NxkhC7s14DydF/NEWnEn1L4Mk7SpbBvx
62UkluuYpHXCVIosJQTk64qh04wbj9rF85IbYq/oj2tQbmZpYuVsAj0qyaQhysdE
b41NRt8u5r84MBiLGt2/sdGl8WvHfh5dN8IwajRAdwECgYEA5EDE+V/6avE0+uM2
dSH5AuBxNh+M/HIwqJZioHiF8Z2Mptgd8p6jzr+mDCL0a1crhpJCMAbE3UH+CJDg
FogcShBCTGqGJitsDpoJYxxMrhIybHHKI3iaclO/+AYzj20ivzXigke1tqMruGjk
oXn720EVhHcDoadf9It+0HWtgMECgYEA1057IwukasL/ounD9TtVohRgoFSozDo9
veQtKMACD51n/NLYq9pmBhCDQrmiJyn3wIR1OhCV6g40kwUmPHLo2d0Wg5zlcw7i
P9nVnO0X45dr/KTyuAl9QqFKiNd21V0ZTw3nNvWtd9GiY/ceFTdqXjuDU+lrinbH
PrK+wflmuVsCgYBEiq6btBaexpACRvizc+Ay3PMiLZdIt4GlHoO/bHHMfGMgNwXH
NZw7GZHjUxvSn/qSDpmRknbLemWubNLd8UGNfBRsnhBqpd9tAxSOjpD4NL6vkfwB
atX2PvAGqtYQ29TzVlsOhDos/hLC7by6QMdAr+qmPJb6lChcZwdN0gioAQKBgD9h
HuZmGvRCirYOUlzyJncbwIXx5e6YhmxLqu9/9htanq3R825DUB0g50LRGmak+AV3
+HorP7YykE9nCKZqvRjE+Eet++0uyHM7UKJtOMcKYANzGvAJ+xGOIT6/DoAoc7bN
xO5sy6+lykPbDsP4GBu9MR+Was3LwUM3oue+3vpZAoGAICS8gnuYP14JrTxCcdPX
pdwKF6hN9/J4eiACAd4eP9Q+wOik0u21iEUYUDv2DQdPC7Z9akuxZ7MpI84GcnVD
hy0SpiNifyyZVKb1LkdJp8yTgQG7LAq+gSBRGEJfL1sn6ZweXLmp4QRaFvEy2G1D
YbHCkyYU6sHHy8N/Z/nqLbM=
-----END PRIVATE KEY-----
EOT
)
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcInternalCertificateConfig() string {
	config := `resource "fmc_internal_certificate" "test" {` + "\n"
	config += `	name = "my_internal_certificate"` + "\n"
	config += `	certificate = local.cert` + "\n"
	config += `	private_key = local.cert_key` + "\n"
	config += `	password = ""` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_internal_certificate" "test" {
			id = fmc_internal_certificate.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcInternalCertificateConfig() string {
	config := `resource "fmc_internal_certificate" "test" {` + "\n"
	config += `	name = "my_internal_certificate"` + "\n"
	config += `	certificate = local.cert` + "\n"
	config += `	private_key = local.cert_key` + "\n"
	config += `	password = ""` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_internal_certificate" "test" {
			name = fmc_internal_certificate.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
