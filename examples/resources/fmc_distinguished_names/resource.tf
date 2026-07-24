resource "fmc_distinguished_names" "example" {
  items = {
    my_distinguished_names = {
      distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"
    }
  }
}
