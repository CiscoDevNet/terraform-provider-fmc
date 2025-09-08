resource "fmc_dynamic_access_policy" "example" {
  name                             = "my_dynamic_access_policy"
  description                      = "My Dynamic Access Policy"
  secure_client_posture_package_id = "12345678-1234-1234-1234-1234567890ab"
}
