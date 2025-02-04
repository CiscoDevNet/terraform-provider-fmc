resource "fmc_device_deploy" "example" {
  ignore_warning  = false
  device_id_list  = ["2fe9063e-8bd5-11ef-9475-e4aeac78cf37"]
  deployment_note = "Terraform initiated deployment"
}
