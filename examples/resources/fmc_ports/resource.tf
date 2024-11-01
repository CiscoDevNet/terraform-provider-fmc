resource "fmc_ports" "example" {
  items = {
    ports_1 = {
      port        = "443"
      protocol    = "TCP"
      description = "Port TCP/443 (HTTPS)"
    }
  }
}
