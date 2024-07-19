resource "fmc_fqdn_object" "example" {
  name           = "fqdn_1"
  value          = "www.example.com"
  dns_resolution = "IPV4_AND_IPV6"
  description    = "My FQDN Object"
}
