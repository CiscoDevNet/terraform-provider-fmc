resource "fmc_dns_server_group" "example" {
  name           = "my_dns_server_group"
  default_domain = "example.com"
  timeout        = 2
  retries        = 2
  dns_servers = [
    {
      ip = "10.10.10.1"
    }
  ]
}
