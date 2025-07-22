resource "fmc_secure_client_image" "example" {
  name        = "my_secure_client_image"
  description = "My Secure Client Image"
  path        = "./secure_client_image.pkg"
}
