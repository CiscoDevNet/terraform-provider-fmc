resource "fmc_radius_external_authentication_object" "example" {
  name                          = "my_radius_auth_object"
  description                   = "My RADIUS external authentication object"
  server_address                = "10.1.1.10"
  server_port                   = "1812"
  key                           = "my_secret_key"
  timeout                       = 30
  retries                       = 3
  message_authenticator_enabled = true
}
