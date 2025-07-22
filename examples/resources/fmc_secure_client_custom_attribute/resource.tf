resource "fmc_secure_client_custom_attribute" "example" {
  name                         = "my_secure_client_custom_attribute"
  description                  = "My Secure Client Custom Attribute"
  attribute_type               = "USER_DEFINED_CUSTOM_ATTR"
  user_defined_attribute_name  = "my_user_defined_attribute"
  user_defined_attribute_value = "my_value"
}
