resource "fmc_ports" "example" {
  items = {
    my_ports = {
      protocol    = "TCP"
      port        = "443"
      description = "Port TCP/443 (HTTPS)"
    }
  }
}
