resource "fmc_file_policy" "example" {
  name                         = "my_file_policy"
  description                  = "My file policy"
  first_time_file_analysis     = true
  custom_detection_list        = true
  clean_list                   = true
  threat_score                 = "DISABLED"
  inspect_archives             = false
  block_encrypted_archives     = false
  block_uninspectable_archives = false
  max_archive_depth            = 2
  file_rules = [
    {
      application_protocol  = "ANY"
      action                = "DETECT"
      direction_of_transfer = "ANY"
      file_categories = [
        {
          id   = "5"
          name = "PDF files"
        }
      ]
      file_types = [
        {
          id   = "19"
          name = "7Z"
        }
      ]
    }
  ]
}
