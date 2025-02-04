resource "fmc_tunnel_zones" "example" {
  items = {
    TunnelZone1 = {
      description = "My Tunnel Zone object"
    }
  }
}
