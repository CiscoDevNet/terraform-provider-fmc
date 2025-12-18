resource "fmc_certificate_enrollment" "example" {
  name                                                               = "my_certificate_enrollment"
  description                                                        = "My certificate enrollment"
  enrollment_type                                                    = "SELF_SIGNED_CERTFICATE"
  validation_usage_ipsec_client                                      = true
  validation_usage_ssl_server                                        = true
  validation_usage_ssl_client                                        = true
  skip_ca_flag_check                                                 = false
  certificate_include_fqdn                                           = "NONE"
  certificate_common_name                                            = "ftd.example.com"
  certificate_organizational_unit                                    = "my_organizational_unit"
  certificate_organization                                           = "my_organization"
  certificate_locality                                               = "my_locality"
  certificate_state                                                  = "my_state"
  certificate_country_code                                           = "PL"
  certificate_email                                                  = "me@example.com"
  certificate_include_device_serial_number                           = true
  key_type                                                           = "RSA"
  key_name                                                           = "my_key"
  key_size                                                           = "CertKey_2048"
  ignore_ipsec_key_usage                                             = false
  crl_use_distribution_point_from_the_certificate                    = true
  crl_static_urls                                                    = ["http://example.com/crl.pem"]
  ocsp_url                                                           = "http://example.com/ocsp"
  revocation_evaluation_priority                                     = "CRL"
  consider_certificate_valid_if_revocation_information_not_reachable = false
}
