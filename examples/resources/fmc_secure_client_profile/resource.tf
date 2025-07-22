resource "fmc_secure_client_profile" "example" {
  name        = "my_secure_client_profile"
  description = "My Secure Client Profile"
  file_type   = "ANYCONNECT_VPN_PROFILE"
  path        = "./package.xml"
}
