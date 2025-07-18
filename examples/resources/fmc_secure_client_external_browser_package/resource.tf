resource "fmc_secure_client_external_browser_package" "example" {
  name        = "my_secure_client_external_browser_package"
  description = "My Secure Client External Browser Package"
  path        = "./secure_client_external_browser_package.pkg"
}
