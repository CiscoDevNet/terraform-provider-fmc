resource "fmc_port" "example" {
  port        = "443"
  name        = "tcp443"
  protocol    = "TCP"
  description = "Port TCP/443 (HTTPS)"
}
