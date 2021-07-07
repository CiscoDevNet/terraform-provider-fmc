# FMC Provider for terraform

This project is a terraform provider which interacts with the FMC APIs making it manageable statefully from terraform.

## 0. Prerequisites

Go ahead and install terraform on your machine. You can get more instructions for that [here](https://learn.hashicorp.com/tutorials/terraform/install-cli) or directly download the relevant package [here](https://www.terraform.io/downloads.html).

## 1. Installation - Development

Follow the below steps to add this provider to your terraform plugins directory. This is only needed for development. For normal use, let the terraform download it from Hashicorp registry.

- Download the latest suitable release from [Releases](https://gitlab-sjc.cisco.com/tfprovider/fmc-terraform/releases).
- Follow the [document here](https://www.terraform.io/docs/cli/config/config-file.html#provider-installation) to install the same.

## 2. Usage

Note: Check examples for more detail.

- Create a new folder anywhere suitable for the project
- Create a file called `main.tf` and initialize the provider as shown

```hcl
terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
    }
  }
}
```

- Run `terraform init` in a terminal from the same folder. You should not see any errors and see the output similar to the one shown below:

```bash
Initializing the backend...

Initializing provider plugins...
- Finding latest version of ciscodevnet/fmc...
- Installing ciscodevnet/fmc v0.1.1...
- Installed ciscodevnet/fmc v0.1.1 (self-signed, key ID 6EC4A79DAB7CB6D0)
```

That's it! You have successfully installed the FMC terraform provider. Head on to examples to see what you can do with them!
Provider documentation is present [here](https://registry.terraform.io/providers/CiscoDevNet/fmc/latest/docs) for now.
