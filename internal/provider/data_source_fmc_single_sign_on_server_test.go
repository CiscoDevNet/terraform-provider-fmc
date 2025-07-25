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

func TestAccDataSourceFmcSingleSignOnServer(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "name", "my_sso_server"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_single_sign_on_server.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "identity_provider_entity_id_url", "https://idp.example.com/saml"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "sso_url", "https://idp.example.com/sso"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "logout_url", "https://idp.example.com/logout"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "base_url", "https://fmc.example.com/sso"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "request_signature_type", "RSA-SHA256"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "request_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "identity_provider_accessible_only_on_internal_network", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_single_sign_on_server.test", "request_identity_provider_reauthentication_at_each_login", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSingleSignOnServerPrerequisitesConfig + testAccDataSourceFmcSingleSignOnServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcSingleSignOnServerPrerequisitesConfig + testAccNamedDataSourceFmcSingleSignOnServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcSingleSignOnServerPrerequisitesConfig = `
locals {
  certificate = <<EOT
    3082 10df 0201 0330 8210 9506 092a 8648
    86f7 0d01 0701 a082 1086 0482 1082 3082
    107e 3082 066a 0609 2a86 4886 f70d 0107
    06a0 8206 5b30 8206 5702 0100 3082 0650
    0609 2a86 4886 f70d 0107 0130 5f06 092a
    8648 86f7 0d01 050d 3052 3031 0609 2a86
    4886 f70d 0105 0c30 2404 10a7 3517 16e2
    58d2 e3de 7f9c e59a b235 9b02 0208 0030
    0c06 082a 8648 86f7 0d02 0905 0030 1d06
    0960 8648 0165 0304 012a 0410 2456 3c59
    7371 7f3a 1677 0a5b 6e7c 448c 8082 05e0
    83ad 993a 36e7 e754 59da 286d 1e8f 5e96
    576c 54f6 ebb2 be61 12b7 1407 149d a968
    13e9 65db fce8 589d 098b ea2b 5a97 3735
    8665 7fc8 04c2 679e e537 9684 ab9a b527
    6f09 ba90 2964 0817 b68d 8ad9 cddb d238
    6fe1 f174 152b 854c 8707 55de d454 90d2
    0611 b8cf 4fc0 33bc 17d6 45c9 ee35 4ae0
    9b0f fe47 4837 113d bede 9f25 e1b9 28b2
    3b85 a0b3 f623 94bd 3393 c4cd c473 2da3
    57cc f6d0 f0a8 c5f8 6fab 4855 8d73 91e5
    6edf 754b 2395 c200 4db9 05ab 73af c3b0
    1fec 55a6 d540 886f 1b27 e7bb 18c9 151b
    62c4 55e1 1e90 5880 8abb 2fe7 f9f4 63bd
    e101 7ee0 9f8e 47e6 c250 674b 54a9 c542
    c9fe 449a 2042 c4b0 3ccc 9c12 fc62 6966
    a498 9860 f786 9ac3 4ea7 4481 6d81 7356
    60c9 a884 1570 2eb2 aeb2 c2a3 34e6 c584
    532a f5ce 6293 a07d 8016 9465 c83c e6d9
    68e8 9fbf 82f3 c37d d3cd 70d0 6ece 7f78
    c20a 9098 d02b b7cf b7f5 5fbb c911 ce02
    f984 9080 d5b3 1ae5 067d 6d39 11e3 7e64
    b961 ea34 76f8 b1fd c39f eada 7f0d 284e
    f563 6680 be6a 8d2b a9f4 923d 4c12 6dd7
    5c7b 6a89 c477 6f0b 2cf5 b8ef 8e32 11d0
    8709 bfb4 060e ae3a ab7d a2bb acca e05d
    5127 1b10 f8b3 f9a8 3171 e796 08f1 f484
    cb25 58f0 a4dc be7d 5f4f 5f14 24ec 185e
    a464 2a04 73a7 15a6 d3e2 1c94 041d 3826
    786f 3839 7e92 fab9 576d 3d9d 5211 3981
    379f 63dd 41e8 5905 4d7f 71d9 3ec0 db81
    a847 4468 5e97 633e 2b4b 9e16 5967 84ed
    fdb9 72b8 97a9 318f d681 606c 6dd7 2b5e
    1719 6701 837a 2d17 d4cc 3ec4 6b06 6d6f
    eb11 ec0f ca3d d24c c0ad 9c45 6159 2503
    1002 9721 d3e9 3563 6cd4 0a33 b573 086d
    c2d4 a259 8655 2d57 ed9a f462 b947 62ab
    44ca 0829 9f23 ebcd 48f2 ee19 aa58 db47
    5c1b 3395 fba3 5389 6e16 64a2 2a41 37a1
    313e 30e3 d49c 8433 f113 3191 e15c 641b
    8108 dd55 7da0 a89d 77df 4908 b1e2 e767
    a12c e4a3 a6dd 4293 f40f c4c3 1f0a 6305
    7d62 0d92 567b 22f3 8647 0b11 d15d c12d
    e42c 47ea 1895 b184 6478 4e96 d2bd 72ee
    5b84 04e4 2f3a d1a5 9b8b ef14 2c3e 0392
    11e5 6253 dc95 e024 069e c65d 5e26 af90
    76ff f6fb c7d2 8d62 aed7 2bc4 0e8d 4421
    7e92 78d9 e7cc 4a36 0e0e f485 2cd4 a3d4
    1679 b997 ae16 b3fd a113 4da1 226c 2f8f
    a6bc b250 5875 edfa feb6 3579 d552 12da
    e73e 8269 f4fc 6a23 1d8e 0d68 8a58 ff41
    cf72 4f9b e69f 9be4 5b0e 7b60 b620 f11e
    daeb eb77 1dd2 d3df a43d a3db e823 985a
    d380 fa06 86cc 1bf5 2abd cc15 9de2 3ff8
    60ba b50d cde2 7143 a21c 7d67 8917 e91f
    ac42 9816 b233 0d67 49aa c832 4800 2b51
    27e8 d63b d33b 94bc 1c85 154d 9d8f 17bf
    118c 02fa f11e e041 0d9a 348b 1df6 193a
    2fff dd7b 3f18 773b 3282 59a1 2df4 bb5a
    313a 687e f627 1c41 baf6 361e 9238 d7f8
    8b5d e5a0 c0e1 8b2f c9c1 5d7c 2df3 4a2a
    b8aa dff9 566b 204e 29b8 426c a842 b0d3
    b4be ce99 acb7 604c ba8b dc24 fb9c 1548
    5536 18d5 b174 27f1 4e35 1bd9 eb8b 7cc1
    df57 10d3 e888 2d02 361b cc2f d933 e289
    053f b7b2 499e 2be8 743e 1cc4 f5d9 3c45
    f3aa ba01 703b 48d6 ff9c 69fd 9ec4 7ce9
    e0c9 3f25 e28f 7349 a558 b77b b5e0 9e2b
    b6fc d079 4495 50de ed1b d109 69bd 97e0
    9525 6be2 98d7 0ec9 ec43 807c 15e6 6052
    0d37 a8f1 f7e9 d687 6d62 ea3f c5d3 2fd8
    70d3 8f0c 51b1 7e76 184d 5991 2ad4 3ec3
    c38a 510e fea0 6e52 64be efd9 8610 6ce0
    b212 9f79 1f5d b271 5b34 8832 ae80 faea
    95b7 a323 1d94 75be 4d4a 3faa ddb5 695b
    863b 526f ea18 f8dc 8ff7 18bc 2854 e9cf
    fca7 0069 6891 f011 bb87 f52f 0b63 aa3e
    a783 58d6 bb25 639a b215 b5a8 2abb 5aeb
    72dc 4bf8 922f ac39 9fd2 4e25 5305 99a9
    a547 00a7 e209 b194 bedb 7f73 379a 3f2a
    9617 f918 e2a4 dd67 f941 2028 fcfa 84c1
    5fd5 f22a 5d93 e21b ecb1 317e aeb0 ffd7
    f04a 398b dd5d 01bc 6a7c 0858 c94d c36b
    ba07 38a9 b467 363a d13e b8d6 a24f a1ff
    3885 85c7 6b8a 4b1b c71a 9fee 6d7e a26b
    3f59 80c2 6538 7273 84a8 0b94 b94d 6d71
    aef6 66a6 9b9d a9a2 9d1f 24a0 6f6b 16a7
    6f99 3c7d 1d1a 5032 16e2 318c 7fa3 ec3e
    0327 462e 9cd6 999d 8be5 0a7f 5403 f64f
    2ade 63c0 2487 0159 124b d0be 2294 9efb
    17a5 8bb3 1a61 5253 8948 ba17 2356 560d
    d883 a4d1 5dd0 8f96 2462 b4d7 9ef3 5f54
    ee82 f23f 94f4 72ac 66b8 598a 7300 00ce
    28aa 40f1 de67 cc17 64d3 a440 e923 b485
    0ca0 dbb3 6741 7b62 b193 3768 7e82 a82f
    3082 0a0c 0609 2a86 4886 f70d 0107 01a0
    8209 fd04 8209 f930 8209 f530 8209 f106
    0b2a 8648 86f7 0d01 0c0a 0102 a082 09b9
    3082 09b5 305f 0609 2a86 4886 f70d 0105
    0d30 5230 3106 092a 8648 86f7 0d01 050c
    3024 0410 27d2 1c19 3f38 8318 1c5d a153
    1833 60ef 0202 0800 300c 0608 2a86 4886
    f70d 0209 0500 301d 0609 6086 4801 6503
    0401 2a04 10bd 6fcb 487d 07cf dd74 43c9
    5e0b b242 4004 8209 5038 5cab 99ef efbe
    f081 d5f0 78a4 51aa 5168 5c07 2202 545f
    a7a3 c96f 5bd7 14ad d998 5ec0 f98d d6e5
    79f1 842a 30a4 0fe3 098f c1f0 bd0b 04b2
    2c6a 233c 487f 39ac 9fd0 a5b8 a227 eb43
    1b03 f268 adae 7761 3a13 5552 b2d8 6cb5
    b9f5 7642 142c b797 2d32 e355 4794 7505
    749a fe14 b831 ad53 9fb1 6f16 1b72 6c29
    e14d 559a 38b5 97cb 1000 da0c 6bca c149
    4413 409f 7f29 10ad 6c3a 8365 5230 332e
    9550 8062 df3f b8dc bb05 bfa4 660a 0636
    d90a 1d38 3c57 9862 2d3b e48e 47ef 6efc
    2348 95b5 db8b 1eea 72c6 cf4f a72b 6665
    33c3 d204 a2cd 93a1 c61f c55a 1a67 d830
    72de a20c 6b28 639c 1b87 1637 6722 eba9
    a124 db2f 076d 40f1 c10e 0465 8575 1410
    4849 8c26 2470 beff b594 295b 4958 4fd5
    d352 20a5 7552 8da6 8580 923a e481 40b7
    bef7 aa60 ef27 65b2 49c1 c2e7 68d1 5228
    ff8e 8ddc ebc6 5782 c150 3b5e 472d 6b7d
    5792 76dc a0eb 7c0b 5dc3 e395 0762 ccdf
    d49e f03b 48d3 1379 b542 11c2 8c00 2122
    de92 b304 c983 0d82 b62b 7aa6 e43a b966
    8ffe 582b 84d3 bd78 f1f6 a237 05f5 def5
    234d 2bd8 6430 b263 a412 5116 a0ea 0268
    5c9d bd3d 901a 27fe 33da 9679 fa50 2e72
    e15d 4c5e 7955 c903 cbc1 752d 44f7 3786
    60ca aaa8 c291 df1d 50b5 c030 8aa6 0ce7
    402f ace9 305b 3e4b fd3f 7827 fb3b ffb9
    1e5e 3720 0eb8 2972 e66b eeec 7a57 e938
    bb68 8bf1 ed8f 1ddd 3a53 95ca 9090 7bf3
    9cdd 56b3 fffb 65f4 2bf7 1401 4968 38cf
    5b8a 5967 54e0 828d 1e0a 5510 a59e 1918
    3ba4 3be3 937b c87a bd5b 2ad5 eb06 f8d0
    25f9 d13d e62a d845 36e2 5e22 270a 6791
    765f c5d4 a9fc 9b5d 7040 a5b9 e8ad 786d
    8f18 6bde 232c 84d2 8d5d dde5 ce47 fb8b
    d9f8 e45d 61bd e388 eb34 15b2 95d6 b538
    047f b861 bcf4 7ca8 0b54 77f5 291c daa1
    e6fd 8196 91e5 36ef 2f90 c175 e2b1 5c92
    2a27 db37 a04f 2088 0928 ce75 f375 47dc
    84de 6880 6215 0525 d4e5 a9fe c544 4fd0
    916b d796 cacd e48b 1c89 3e16 f942 4137
    936d f8d4 d660 d590 a04b 8f04 e369 2417
    45ff ee5a 086d 5888 279b 952d 7a95 9166
    f1e2 0218 f997 3631 5948 439e 884d fee7
    a627 e998 9ac4 4034 e6ec 910b de60 a37e
    bae6 7f62 8104 7a37 935a 58d0 cb3e c7ce
    aa39 4805 cc19 a86e 5e3f c211 ff17 13fd
    4232 f5ac 62f3 fd10 a1fc 8420 1330 e63a
    47c0 9f2a 4fd5 2b2e 8378 7c8e c6d8 35cc
    f17a 5b66 2511 208f fd5b 11e5 0ea0 34e5
    4502 dbc9 3375 23ca ffd2 9eaf 4688 b72a
    9f6b a89e 1ec5 64ae 25f4 42ed 5194 557c
    2051 c94b 4a14 28d5 786b 418b 081d 6c09
    d480 692e ba08 5697 7a44 90b6 1abc 6299
    3ba6 d6af dbb9 5ab8 8f2b 1c4e b99c 83e9
    d057 038e f9b2 0986 3ec7 9544 44a6 1a3b
    f34f 1047 f970 c24b 8a45 aa1c 06d0 d481
    814b 10fc c2e4 8bf7 ed92 566c 9370 a2fc
    85a0 e089 48ec 0d6f 0405 eace 0bf9 5046
    bd99 6d2a 5a28 38dd 7b32 4eb6 9f9c cadd
    3d51 4625 d9df ef66 ed2d e4b1 5ac0 a2cf
    ccd7 b8b1 b630 27f5 8e0f 4dcb 104a 4f2e
    7dc5 362c 54a9 7500 ec78 4a7b 3610 d850
    ca15 adad 71ab a892 c7aa fcf8 47c1 8c10
    51aa 65a0 a0a8 2709 56ee 9d46 31e7 414b
    403c 4398 4ca3 b8f3 69e7 1e36 0dc0 b822
    38ed aaa3 1e89 c3ad 3e15 7c7b 04a9 1035
    4b69 6aa0 c97b a513 6b1f 9fdd 1dd3 ae8f
    9dd2 cfa9 5f5d 0474 e8df 1123 1327 14dd
    6ba2 c961 ed45 4522 c5fc f411 c40f 136a
    b2f2 407e d873 9d4c 0371 1c36 624e 1684
    af02 7e05 9a2e 6087 4126 e1ef f10f fbba
    eb40 3991 bed1 38f3 df35 31a3 46ad b45f
    7efd f349 9ed9 873b 240c bb2b b252 8713
    e411 fdf4 9516 1856 5364 4840 9777 a726
    d986 59dc 09e6 2841 6382 448e 8cf5 72ba
    fb66 6b90 dde1 2f3b 397d 7ea9 609b e5cf
    5cc4 0a56 e72e df43 a296 68c7 bc62 5766
    7311 5619 2e2b 696d 4c05 d2a9 38fb 5f07
    8485 d5f7 b8a3 70ca be72 0896 7390 9d35
    88d7 83b8 cf33 e1ac 3c29 9642 03db bf72
    3156 29d1 a263 0ac7 2db3 344a 8148 72f4
    2ceb 0ff3 d072 76af 6e29 1145 2b6b f3e2
    7dac 264c 2b41 d4c6 225f 9d5c 71a8 4b87
    bc86 edae dca1 d472 1869 f316 104b 3f94
    6fe3 e08b 34aa 3953 1f7c 2e23 8343 1efd
    04cc a041 2c83 64ca b4db b894 fa3a 7d12
    0f9f a55a 820f 877c becd cd68 b94e 4641
    9fe8 a75e 0b37 d01f f0a7 8b01 88e2 6719
    db71 5dc9 2122 22f7 821e 9d03 a879 2f1a
    b4b6 abea 6fa6 1bfb ea8f 7b72 b200 228c
    0ed6 a41d 9924 bdc1 cacf bf23 db02 15b7
    74ba 1245 03bb ff44 27a3 d5c6 cf91 bb2b
    bf49 49aa 3f43 77c3 6a08 e233 960c f49d
    c2a4 9cd2 3589 d230 8dc5 9888 bd93 38a7
    4527 7849 2222 f6e2 c0b9 3ccd f3da 3e4b
    c242 86d6 be50 5c24 3c6d f251 e7ee bb2b
    6f65 8c55 1933 041c 05f1 58dd b444 6619
    6aa5 03dd 3936 6e1f 3d2a 31ca b912 d6ae
    3aa8 1126 b0df 0103 7eca 4761 f652 f5cc
    d95b d823 7578 8225 e29a 75fe e7b2 55f2
    e79c bb15 e12c 5374 1dde ff3d 3797 f15c
    bed3 27b4 0371 8095 9cbb c547 10e7 c194
    8d77 a7d4 2de8 37aa 6386 6820 bee6 8717
    4457 4fd5 6ef7 d6f8 3b41 9d18 7bc2 08a7
    82b8 53c6 3034 6714 b251 4336 2550 01b3
    f5b3 4486 9a90 68b7 6b5c 5f5d 8458 be97
    102a 6983 b8f1 b6b3 d6f1 77a3 758e 2e17
    bf40 3aae c137 4d3c 6b01 caab c513 8063
    f715 75c7 dae4 7204 79e8 4743 b1e3 c6de
    2aa5 8252 5fbb 5809 9a29 cd9e 4a53 0d87
    9c58 9c76 892d f43c 1e58 e6b1 9719 3ecd
    8476 63b3 8e2d 3507 eea0 4657 72d2 ae38
    5cce 53ba b062 20e6 bf4d 8b2a b879 d69a
    591e 4337 f270 7d83 6f51 f529 09da 6bfc
    cfd0 7ba6 3c88 b416 571a b89f 07f8 3de6
    7600 d89c cadb 6c6e 9f6f 5f0a 0e47 a5c4
    81ab 2c29 5a8f 949f 2e42 8434 6b0e fa48
    ffb3 5dca deeb 257f 5519 2ff0 1537 f86a
    fbc5 775d 591e 4ab9 9b98 841d 761b bce4
    519b 606b 3bc7 cf64 1ce1 7751 df26 ee7e
    97c9 911a d731 08c3 cf53 9031 077b 4742
    cb36 bcfc dc98 96f1 4dfc 82f4 9b79 ec68
    b6c3 0846 c190 10c9 dad5 803f a5fd 0e3e
    d0e6 d81c 13f3 1f97 7f36 2efb 9c7c 465d
    b367 c3d5 4c10 deeb 893b c993 a41c 0005
    d03f 18df 15d4 39ff 7fbb cd7a 7e89 6d93
    364e a2e1 cb83 e572 bfa5 256d 1990 0a9a
    581b eec8 dcc6 8263 38c8 8568 8bd8 e6d2
    f499 c9c6 4f13 5c05 3014 787a 3de6 dc79
    eb2d 305c 0303 49a8 90b8 2a9b f88c c2f7
    97f2 5342 2b8f 6f89 4b2a b530 d45a c5ee
    153e 43ce e46d 5cf6 d52e 14bd a59b c6d7
    becc c7bd 607a 228e 91d7 b577 5914 b572
    f36e d830 3cb9 8ebc a61f 08ba 6080 c430
    3b45 7031 e92a 9c5b c460 79d9 767b 8472
    7522 ee68 7294 645f 9d10 4c29 9c21 83f3
    5513 2f38 f48a 3ea4 5ca1 386e 2405 6435
    a3f7 a552 6596 cc11 e90d 104f 9112 5a88
    19db 28e5 717a d02a 48b8 aad8 4853 01b2
    eb41 4c18 4cc4 7389 3b5e c41c c383 eda2
    0c70 be6c f9ce c5e6 c196 ddf0 8c90 0062
    3468 7d00 3f6a 5ffa 2dfb 2608 fa2d 4cfc
    f3b7 b21a d281 20ba 3614 bc1b 8047 16b8
    eebd 533b bf7d 3fe3 b8de 2c34 89b5 94fd
    bc2a 3c83 6498 b149 5c67 dd09 b57a 6943
    300a aba8 2bc5 65ee c7c3 3528 0451 5a0c
    5976 ec38 5f78 3d39 f131 2530 2306 092a
    8648 86f7 0d01 0915 3116 0414 e62a f1f9
    9114 fd72 0fcc 18e6 7c62 b0a7 9799 acb7
    3041 3031 300d 0609 6086 4801 6503 0402
    0105 0004 203f 9bcf b843 2c7f bed5 f63d
    0c33 8297 8dc5 5ee8 ff8f 8dc4 a6a6 8c76
    299d f868 2c04 08ba 6b51 98fa 105d 8302
    0208 00
    EOT
  }

resource "fmc_certificate_enrollment" "test" {
  name                          = "sso_certificate_enrollment"
  enrollment_type               = "PKCS12"
  pkcs12_certificate            = base64encode(local.certificate)
  pkcs12_certificate_passphrase = "cisco123"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSingleSignOnServerConfig() string {
	config := `resource "fmc_single_sign_on_server" "test" {` + "\n"
	config += `	name = "my_sso_server"` + "\n"
	config += `	identity_provider_entity_id_url = "https://idp.example.com/saml"` + "\n"
	config += `	sso_url = "https://idp.example.com/sso"` + "\n"
	config += `	logout_url = "https://idp.example.com/logout"` + "\n"
	config += `	base_url = "https://fmc.example.com/sso"` + "\n"
	config += `	identity_provider_certificate_id = fmc_certificate_enrollment.test.id` + "\n"
	config += `	identity_provider_certificate_name = fmc_certificate_enrollment.test.name` + "\n"
	config += `	request_signature_type = "RSA-SHA256"` + "\n"
	config += `	request_timeout = 300` + "\n"
	config += `	identity_provider_accessible_only_on_internal_network = false` + "\n"
	config += `	request_identity_provider_reauthentication_at_each_login = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_single_sign_on_server" "test" {
			id = fmc_single_sign_on_server.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcSingleSignOnServerConfig() string {
	config := `resource "fmc_single_sign_on_server" "test" {` + "\n"
	config += `	name = "my_sso_server"` + "\n"
	config += `	identity_provider_entity_id_url = "https://idp.example.com/saml"` + "\n"
	config += `	sso_url = "https://idp.example.com/sso"` + "\n"
	config += `	logout_url = "https://idp.example.com/logout"` + "\n"
	config += `	base_url = "https://fmc.example.com/sso"` + "\n"
	config += `	identity_provider_certificate_id = fmc_certificate_enrollment.test.id` + "\n"
	config += `	identity_provider_certificate_name = fmc_certificate_enrollment.test.name` + "\n"
	config += `	request_signature_type = "RSA-SHA256"` + "\n"
	config += `	request_timeout = 300` + "\n"
	config += `	identity_provider_accessible_only_on_internal_network = false` + "\n"
	config += `	request_identity_provider_reauthentication_at_each_login = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_single_sign_on_server" "test" {
			name = fmc_single_sign_on_server.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
