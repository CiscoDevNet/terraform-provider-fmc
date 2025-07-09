resource "fmc_security_intelligence_url_feeds" "example" {
  items = {
    my_si_url_feeds = {
      feed_url         = "https://example.com/path/to/feed.txt"
      checksum_url     = "https://example.com/path/to/checksum.md5"
      update_frequency = 120
    }
  }
}
