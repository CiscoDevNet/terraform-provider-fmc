resource "fmc_port" "example" {
  port        = "443"
  name        = "my_port"
  protocol    = "TCP"
  description = "Port TCP/443 (HTTPS)"
}
