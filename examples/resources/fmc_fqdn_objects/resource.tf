resource "fmc_fqdn_objects" "example" {
  items = {
    fqdn_1 = {
      description    = "My FQDN 1"
      fqdn           = "www.example.com"
      dns_resolution = "IPV4_AND_IPV6"
    }
  }
}
