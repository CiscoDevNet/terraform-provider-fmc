variable "fmc_username" {
    type = string
    default = "admin"
}

variable "fmc_password" {
    type = string
    default = "Cisco@123"
}

variable "fmc_host" {
    type = string
    default = "20.204.26.117"
}

variable "fmc_insecure_skip_verify" {
    type = bool
    default = true
}