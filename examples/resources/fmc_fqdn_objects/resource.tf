resource "fmc_fqdn_objects" "example" {
  items = {
    my_fqdn_objects = {
      description    = "My FQDN 1"
      fqdn           = "www.example.com"
      dns_resolution = "IPV4_AND_IPV6"
    }
  }
}
