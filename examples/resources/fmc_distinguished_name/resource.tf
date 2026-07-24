resource "fmc_distinguished_name" "example" {
  name               = "my_distinguished_name"
  distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"
}
