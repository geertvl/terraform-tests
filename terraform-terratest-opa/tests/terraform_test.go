package tests

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "testing"

    "github.com/gruntwork-io/terratest/modules/azure"
    "github.com/gruntwork-io/terratest/modules/opa"
    "github.com/gruntwork-io/terratest/modules/random"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformAzureResourceGroup(t *testing.T) {
    t.Parallel()

	resourceGroupName := fmt.Sprintf("test-rg-%s", random.UniqueId())

    terraformOptions := &terraform.Options{
        TerraformDir: "../",

        Vars: map[string]interface{}{
            "resource_group_name": resourceGroupName,
            "location":            "East US",
        },

        // Path to .tfvars file
        VarFiles: []string{"./terraform.tfvars"},

        // Disable colors in Terraform commands so its easier to parse stdout/stderr
        NoColor: true,
    }

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	tfShowOutput, err := terraform.RunTerraformCommandAndGetStdoutE(t, terraformOptions, "show", "-json")
    if err != nil {
        t.Fatalf("Failed to run terraform show: %v", err)
    }

	outputPath := filepath.Join(t.TempDir(), "terraform.json")
    err = ioutil.WriteFile(outputPath, []byte(tfShowOutput), 0644)
    if err != nil {
        t.Fatalf("Failed to write terraform show output to file: %v", err)
    }

	policyPath := "../policy.rego"
	opaOpts := &opa.EvalOptions{
		RulePath: policyPath,
		FailMode: opa.FailUndefined,
	}

    terraform.OPAEval(t, terraformOptions, opaOpts, "data.example.allow")

    // Verify the Resource Group exists
    exists := azure.ResourceGroupExists(t, resourceGroupName, "")
    assert.True(t, exists, "Resource Group does not exist")
}