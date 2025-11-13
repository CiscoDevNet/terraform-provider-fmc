resource "fmc_device" "example" {
  name                     = "my_device"
  host                     = "10.0.0.1"
  licenses                 = ["ESSENTIALS"]
  registration_key         = "key1"
  performance_tier         = "FTDv5"
  snort_engine             = "SNORT3"
  object_group_search      = true
  access_control_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
}
