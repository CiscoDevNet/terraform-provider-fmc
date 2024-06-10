resource "fmc_device" "example" {
  name                 = "device1"
  host_name            = "10.0.0.1"
  license_capabilities = ["BASE"]
  registration_key     = "key1"
  access_policy_id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  performance_tier     = "FTDv50"
}
