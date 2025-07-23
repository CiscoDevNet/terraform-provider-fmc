resource "fmc_single_sign_on_server" "example" {
  name                                                     = "my_sso_server"
  identity_provider_entity_id                              = "https://idp.example.com/saml"
  sso_url                                                  = "https://idp.example.com/sso"
  logout_url                                               = "https://idp.example.com/logout"
  base_url                                                 = "https://fmc.example.com/sso"
  identity_provider_certificate_id                         = "123e4567-e89b-12d3-a456-426614174000"
  identity_provider_certificate_name                       = "my_idp_certificate"
  service_provider_certificate_id                          = "123e4567-e89b-12d3-a456-426614174001"
  request_signature_type                                   = "RSA-SHA256"
  request_timeout                                          = 300
  identity_provider_accessible_only_on_internal_network    = false
  request_identity_provider_reauthentication_at_each_login = false
}
