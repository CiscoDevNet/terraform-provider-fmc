resource "fmc_standard_access_list" "example" {
  name        = "my_standard_acl"
  description = "My standard ACL"
  entries = [
    {
      action = "DENY"
      objects = [
        {
        }
      ]
      literals = [
        {
          value = "10.1.1.0/24"
        }
      ]
    }
  ]
}
