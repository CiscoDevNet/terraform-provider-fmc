resource "fmc_external_authentication_radius" "example" {
  name                          = "my_radius_auth_object"
  description                   = "My RADIUS external authentication object"
  server_hostname               = "10.1.1.10"
  server_port                   = 1812
  server_key                    = "my_secret_key"
  timeout                       = 30
  retries                       = 3
  message_authenticator_enabled = true
  ui_access_role_mapping = [
    {
      map            = "User-Category=Analyst,jdoe"
      user_role_name = "Administrator"
      user_role_id   = "bee2eb18-e129-11df-a04a-42c66f0a3b36"
      user_role_type = "AuthRole"
    }
  ]
}
