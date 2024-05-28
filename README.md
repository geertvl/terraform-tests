# General

> Learn, inspire, automate, observe!

In this course we are going to learn and get inspired. 

> Test your intent, and not simply the implementation.

- Writing tests that validate the desired outcomes (e.g., a resource exists in the correct location) rather than how those outcomes are achieved.
- Using policies (like OPA) to define and enforce the desired state of your infrastructure, making your tests more resilient to changes in implementation.
- Ensuring that your tests focus on the high-level goals and requirements of your infrastructure, providing assurance that your infrastructure meets its intended purposes regardless of specific implementation details.

## Terraform vs. OpenTofu

- Newer versions of Terraform will be placed under the BUSL license. Owned by IBM now. 
  - BUSL license:
    - Commercial use restrictions. Barrier for businesses that rely on Terraform for critical infrastructure automation.
    - Uncertainty and lock-in: Restrict usage in the future or require a commercial license. Lock-in if a switch is necessary.
    - Impact on ecosystem. Community, contributors, and businesses will be impacted on the broader ecosystem. Extensions (Terratest, kitchen-terraform, ...) will not be able to provide all features.
    - Perception of open source: Affect the trust and goodwill built within the community.
- OpenTofu true open-source version.
  - Expand on existing concepts and offerings of Terraform.
  - Will introduce improvements and enhancements.
  - The future of the Terraform ecosystem.

## Tools

- Use ```terraform fmt``` to format your code.
- Use ```terraform validate``` to verify the syntax.
- Use ```terraform plan``` to verify the config file will work as expected.
- Use ```tflint``` the syntax and the structure.
  - Download tflint.exe
  - Add it to your PATH environment variable with ```.\set-to-path.ps1 -DirectoryToAdd "c:\tools"```
- Use ```terraform test``` for unit tests and validating outputs and state.
  - See ./terraform-test for an example and run in the root of this folder ```terraform test``` to run the tests.
- Use **Terratest** for true integration tests and real end-to-end tests.
  - Supports Az CLI, Docker tests, https calls, enforce policy (OPA), ... (even test the connectivity of your software).
  - TODO: opa and the rego language
  - You need to learn the Go language [A tour of GO](https://go.dev/tour/welcome/1).
- Use **Kitchen-Terraform** for declarative testing with a focus on compliance and security testing.
  - We do not cover this in the planned sessions. (I will write a Confluence page on that on how and why you would use it in our infrastructure context).
- Use cookiecutter to create your templates for projects, modules, ...
  - TODO: example
# Terratest

## Steps

- Create the module in your tests folder
  ```
    go mod init github.com/geertvl/terraform-terratest-opa/tests
  ```
- Get the needed modules
  ```
    go get github.com/gruntwork-io/terratest/modules/terraform
    go get github.com/gruntwork-io/terratest/modules/opa
    go get github.com/gruntwork-io/terratest/modules/azure
    go get github.com/stretchr/testify/assert

    [See all available modules here](https://github.com/gruntwork-io/terratest/tree/master/modules)
  ```
- Tidy up your Go module
  - After adding new dependencies
  - After removing dependencies
  - Clean up your dependencies that are not used.
  - Before committing changes
  ```
    go mod tidy
  ```

## Solutions

In Powershell we need to add the following line when you get this error.

> Az is not recognized as an internal or external command...

To resolve this you need to add:

```$env:AzureCLIPath = "C:\Program Files\Microsoft SDKs\Azure\CLI2\wbin"``` in Powershell.

If this does not help you need to execute:
```where.exe az``` to find the location where Az CLI is installed. 

```go test -v -timeout 30m```

# Cookiecutter

- [Install](https://cookiecutter.readthedocs.io/en/stable/README.html#installation) 
- Run from the root ```pipx run cookiecutter .\terraform-cookiecutter\```
- More documentation [Read here](https://cookiecutter.readthedocs.io/en/stable/index.htm)

