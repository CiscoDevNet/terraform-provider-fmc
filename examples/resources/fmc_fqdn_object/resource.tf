resource "fmc_fqdn_object" "example" {
  name           = "my_fqdn_object"
  fqdn           = "www.example.com"
  dns_resolution = "IPV4_AND_IPV6"
  description    = "My FQDN Object"
}
