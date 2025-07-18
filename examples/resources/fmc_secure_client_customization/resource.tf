resource "fmc_secure_client_customization" "example" {
  name               = "my_secure_client_customization"
  description        = "My Secure Client Customization"
  customization_type = "LANGUAGE_LOCALIZATION"
  language           = "pl"
  path               = "./AnyConnect_pl-pl.po"
}
