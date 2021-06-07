# FMC Provider for terraform

This project is a terraform provider which interacts with the FMC APIs making it manageable statefully from terraform.

## 1. Installation

Go ahead and install terraform on your machine. You can get more instructions for that [here](https://learn.hashicorp.com/tutorials/terraform/install-cli) or directly download the relevant package [here](https://www.terraform.io/downloads.html).

Once you have that, follow the below steps to add this provider to your terraform plugins directory.

### 1a. Windows build 17063 or later

- Download the latest [`release.tgz`](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest/download/release.tgz) from [Releases](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest).
- Extract the `hashicorp.com` folder to `%APPDATA%\terraform.d\plugins\`
  - Open a command prompt where the file is downloaded
  - Run the below commands

  ```cmd
  md %APPDATA%\terraform.d\plugins\

  tar -C %APPDATA%\terraform.d\plugins --strip-components=1 -xzvf release.tgz release/
  ```

### 1b. Any Windows build

- Download the latest [`release.zip`](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest/download/release.zip) from [Releases](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest).
- Open the ZIP file and extract the hashicorp.com folder to `%APPDAT%\terraform.d\plugins\` directory
- Once done, the should look like `%APPDAT%\terraform.d\plugins\hashicorp.com\cisco\fmc\0.2\`

### 1c. macOS/Linux

- Download the latest [`release.tgz`](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest/download/release.tgz) from [Releases](https://wwwin-github.cisco.com//tfprovider/fmc-terraform/releases/latest).
- Extract the `hashicorp.com` folder to `~/.terraform.d/plugins/`
  - Open a Terminal where the file is downloaded
  - Run the below commands

  ```bash
  mkdir -p ~/.terraform.d/plugins/

  tar -C ~/.terraform.d/plugins/ --strip-components=1 -xzvf release.tgz release/
  ```

## 2. Usage

Note: Check examples for more detail.

- Create a new folder anywhere suitable for the project
- Create a file called `main.tf` and initialize the provider as shown

```hcl
terraform {
  required_providers {
    fmc = {
      source = "hashicorp.com/cisco/fmc"
    }
  }
}
```

- Run `terraform init` in a terminal from the same folder. You should not see any errors and see the output similar to the one shown below:

```bash
Initializing the backend...

Initializing provider plugins...
- Finding latest version of hashicorp.com/cisco/fmc...
- Installing hashicorp.com/cisco/fmc v0.2.0...
- Installed hashicorp.com/cisco/fmc v0.2.0 (unauthenticated)
```

That's it! You have successfully installed the FMC terraform provider. Head on to examples to see what you can do with them!
Provider documentation is present [here](https://wwwin-github.cisco.com/TFProvider/fmc-terraform/tree/doc-do-not-merge/docs) for now.

**Note** When refreshing, planning, applying or destroying from terraform, remember to add `--parallelism=1` to the command, otherwise FMC will block parallel operations on the API. Issue tracked [here](https://gitlab-sjc.cisco.com/tfprovider/fmc-terraform/issues/17).
Ex:
`terraform refresh --parallelism=1`
`terraform plan --parallelism=1`
`terraform apply --parallelism=1`
`terraform destroy --parallelism=1`
