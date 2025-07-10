resource "fmc_security_intelligence_network_feed" "example" {
  name             = "my_si_network_feed"
  feed_url         = "https://example.com/path/to/feed.txt"
  checksum_url     = "https://example.com/path/to/checksum.md5"
  update_frequency = 120
}
