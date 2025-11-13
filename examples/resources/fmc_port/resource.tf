resource "fmc_port" "example" {
  name        = "my_port"
  protocol    = "TCP"
  port        = "443"
  description = "Port TCP/443 (HTTPS)"
}
