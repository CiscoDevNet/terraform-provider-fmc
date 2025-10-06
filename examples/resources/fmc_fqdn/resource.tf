resource "fmc_fqdn" "example" {
  name           = "my_fqdn"
  fqdn           = "www.example.com"
  dns_resolution = "IPV4_AND_IPV6"
  description    = "My FQDN Object"
}
