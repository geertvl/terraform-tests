package tests

import (
    "fmt"
    "testing"

    "github.com/gruntwork-io/terratest/modules/azure"
    "github.com/gruntwork-io/terratest/modules/random"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestResourceGroup(t *testing.T) {
	t.Parallel()

	resourceGroupName := fmt.Sprintf("test-rg-%s", random.UniqueId())

	terraformOptions := &terraform.Options{
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"location":            "West Europe",
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	outputResourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, resourceGroupName, outputResourceGroupName)

	// Verify the Resource Group exists
	exists := azure.ResourceGroupExists(t, resourceGroupName, "")
	assert.True(t, exists, "Resource Group does not exist")
}