# FMC Provider for terraform

> **Note:** this Terraform provider is now publically available on the [Terraform Registry](https://registry.terraform.io/providers/CiscoDevNet/fmc/latest).

This project is a terraform provider which interacts with the FMC APIs making it manageable statefully from terraform.

## 0. Prerequisites

Go ahead and install terraform on your machine. You can get more instructions for that [here](https://learn.hashicorp.com/tutorials/terraform/install-cli) or directly download the relevant package [here](https://www.terraform.io/downloads.html).

## 1. Installation - Development

Follow the below steps to add this provider to your terraform plugins directory. This is only needed for development. For normal use, let the terraform download it from Hashicorp registry.

- Download the latest suitable release from [Releases](https://github.com/CiscoDevNet/terraform-provider-fmc/releases/latest).
- Follow the [document here](https://www.terraform.io/docs/cli/config/config-file.html#provider-installation) to install the same.
- Install [tfplugindocs](https://github.com/hashicorp/terraform-plugin-docs) and run "make generate" to regenerate provider documentation. 

## 2. Usage

Note: Check examples for more detail.

- Create a new folder anywhere suitable for the project
- Create a file called `main.tf` and initialize the provider as shown

```hcl
terraform {
  required_providers {
    fmc = {
      source = "CiscoDevNet/fmc"
      # version = "0.1.1"
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
- Installed ciscodevnet/fmc v0.1.1 (signed by a HashiCorp partner, key ID 6EC4A79DAB7CB6D0)
```

That's it! You have successfully installed the FMC terraform provider. Head on to examples to see what you can do with them!
Provider documentation is present [here](https://registry.terraform.io/providers/CiscoDevNet/fmc/latest/docs).

## 3. Troubleshooting
> **Note**: This is for debugging purposes only and it probably **will** break your terraform-provider-fmc, if using for the first time.
> It is recommended to take a backup of the terraform-provider-fmc binary before proceeding. You can use the following command to do so: `cp ~/.terraform.d/plugins/registry.terraform.io/CiscoDevNet/fmc/0.2/darwin_amd64/terraform-provider-fmc ~/.terraform.d/plugins/registry.terraform.io/CiscoDevNet/fmc/0.2/darwin_amd64/terraform-provider-fmc.bak`

The logs are by default are logged to `outputs/reqres`.

The Logger is disabled by default. To enable it, set the environment variable `LOGS` to `REQ_RES` OR `USER` OR `ALL` .

- If the `LOGS` are set to `REQ_RES`, the request and response of the API calls will be logged.
- If the `LOGS` are set to `USER`, the user defined logs will be logged.
- If the `LOGS` are set to `ALL`, both the above logs will be logged.

To test your changes and see the user-defined logs, run the following script:

```bash
#! /bin/bash
EXPORT LOGS=USER
clear

rm -rf vendor

go mod vendor
go mod tidy

go build -o terraform-provider-fmc_0.2_darwin_amd64

mkdir -p ~/.terraform.d/plugins/registry.terraform.io/CiscoDevNet/fmc/0.2/darwin_amd64

mv terraform-provider-fmc_0.2_darwin_amd64 ~/.terraform.d/plugins/registry.terraform.io/CiscoDevNet/fmc/0.2/darwin_amd64/terraform-provider-fmc
```

Using the above script, the new provider will be built and moved to the terraform plugins directory. Now, you can run the terraform commands as usual and see the logs in the outputs directory.

To use self-defined logs, use the following code:

```go
Log.debug("This is a debug log")
Log.info("This is an info log")
Log.warn("This is a warning log")
Log.error("This is an error log")
```
output will be something like:
```
[DEBUG USER] "This is a debug log"

================================================================================================================
[INFO USER] "This is an info log"

================================================================================================================
[WARN USER] "This is a warning log"

================================================================================================================
[ERROR USER] "This is an error log"

================================================================================================================
```
## Tutorials

The Terraform Playlist: https://www.youtube.com/playlist?list=PLyf18hdY22ESR91vJtdvY_4CNAPMR04zP 

Includes: 
- FMC: Terraform and FMC Create User Account (Part1)
- FMC: Terraform and FMC Access Policy Creation (Part 2)
- FMC: Terraform and FMC Access Policy and Rule Creation (Part 3)
- FMC: Terraform and FMC Access Policy Device Assignment (Part 4)
- FMC: Terraform and FMC Access Policy Device Deploy (Part 5)
- FMC: Terraform and FMC URL Objects
- FMC: Terraform and FMC URL Objects and URL Object Groups
- FMC: Terraform and FMC Network Objects
- FMC: Terraform and FMC Network Objects and Network Object Groups
- FMC: Terraform and FMC Host Objects
- FMC: Terraform and FMC Port Objects 
- FMC: Terraform and FMC Port Objects and Port Groups
- FMC: Terraform and FMC Network Range Objects
- FMC: Terraform and FMC ICMP Objects
- FMC: Terraform and FMC FQDN Objects
- FMC: Terraform and FMC NAT Policy
- FMC: Terraform and FMC NAT Policy and Manual NAT Rules 
- FMC: Terraform and FMC NAT Policy and Auto NAT Rules

############# POWER VIDEO #############

FMC: Terraform and FMC Pulling it Together (Part 1 - The Code!) https://youtu.be/TvcuXP3Yn-0 
In this video we will leverage Terraform to build out a dynamic lab that includes deployment of firepower policies and a web server using Ubuntu. We will create multiple objects to support access control policy and nat policy. Then we will create the access control policy, nat policy, assign to a device then push policy. We will delay by 10min the start time of the web server build in esxi to ensure all policies are in place which allows the web server to have controlled access to repos, youtube for the custom page. We will leverage cloud-config to add a user "Kali" to the box and their public key which allows SSH access. We will also update the box using apt-get, pull apache2, start the service and customize index.html.....part 1 covers the code used to automate the entire process. NO HANDS ON REQUIRED! 

FMC: Terraform and FMC Pulling it Together (Part 2 - The Push!) https://youtu.be/fE9ZAzKNMVU 
Pushing out configuration and validation! NO HANDS ON REQUIRED! 


FMC: Terraform and FMC Pulling it Together (Part 3 - The Destroy!) https://youtu.be/EOdwrnb0FfE 
DESTROYING all our hard work! NO HANDS ON REQUIRED!
