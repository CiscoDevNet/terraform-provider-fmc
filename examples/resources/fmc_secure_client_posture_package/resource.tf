resource "fmc_secure_client_posture_package" "example" {
  name        = "my_secure_client_posture_package"
  description = "My Secure Client Posture Package"
  path        = "./secure_client_posture_package.pkg"
}
