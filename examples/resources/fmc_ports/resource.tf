resource "fmc_ports" "example" {
  items = {
    my_ports = {
      port        = "443"
      protocol    = "TCP"
      description = "Port TCP/443 (HTTPS)"
    }
  }
}
