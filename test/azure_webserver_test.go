package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// Ensure we're using the correct terratest version
// go get github.com/gruntwork-io/terratest@v0.56.0

func TestAzureLinuxVMCreation(t *testing.T) {
	t.Parallel()

	// Update these variables with your values
	subscriptionID := "acc96d9d-e040-48a5-9191-cd80276cb66d"
	labelPrefix := "kutt0011"

	// Construct the terraform options with default retryable errors
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"labelPrefix": labelPrefix,
		},
	})

	// At the end of the test, run `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	// Test 1: Verify the VM exists
	vmExists := azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID)
	assert.True(t, vmExists, "VM should exist")

	// Test 2: Verify the NIC exists and is connected to the VM
	nicExists := azure.NetworkInterfaceExists(t, nicName, resourceGroupName, subscriptionID)
	assert.True(t, nicExists, "NIC should exist")

	// Get the VM details to verify NIC is attached
	vm := azure.GetVirtualMachine(t, vmName, resourceGroupName, subscriptionID)
	assert.NotNil(t, vm.NetworkProfile, "VM should have a network profile")
	assert.NotEmpty(t, *vm.NetworkProfile.NetworkInterfaces, "VM should have at least one NIC attached")

	// Verify the NIC ID matches
	attachedNicID := *(*vm.NetworkProfile.NetworkInterfaces)[0].ID
	assert.Contains(t, attachedNicID, nicName, "The attached NIC should match the expected NIC name")

	// Test 3: Verify the VM is running the correct Ubuntu version
	vmSize := azure.GetSizeOfVirtualMachine(t, vmName, resourceGroupName, subscriptionID)
	assert.Equal(t, "Standard_B1s", vmSize, "VM size should be Standard_B1s")

	// Verify Ubuntu version from image reference
	assert.NotNil(t, vm.StorageProfile, "VM should have a storage profile")
	assert.NotNil(t, vm.StorageProfile.ImageReference, "VM should have an image reference")
	
	imageRef := vm.StorageProfile.ImageReference
	assert.Equal(t, "Canonical", *imageRef.Publisher, "Publisher should be Canonical")
	assert.Equal(t, "0001-com-ubuntu-server-jammy", *imageRef.Offer, "Offer should be Ubuntu Server Jammy")
	assert.Equal(t, "22_04-lts-gen2", *imageRef.Sku, "SKU should be 22.04 LTS Gen2")
}
