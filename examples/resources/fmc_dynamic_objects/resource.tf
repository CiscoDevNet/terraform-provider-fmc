resource "fmc_dynamic_objects" "example" {
  items = {
    dynamic_object_1 = {
      description = "My Dynamic Object 1"
      object_type = "IP"
      mappings    = ["10.0.0.1"]
    }
  }
}
