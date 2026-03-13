# Terratest Integration Tests

This folder contains automated integration tests for the Azure web server infrastructure using Terratest.

## Prerequisites

- Go 1.26.1 or later
- Azure CLI configured with valid credentials
- Terraform installed

## Setup

1. Initialize the Go module (already done):
```bash
go mod init github.com/Sighopss/cst8918-w24-A05-kutt0011
```

2. Download dependencies:
```bash
go mod tidy
```

## Running Tests

Run all tests with verbose output:
```bash
go test -v azure_webserver_test.go
```

Run tests with timeout:
```bash
go test -v -timeout 30m azure_webserver_test.go
```

## Test Coverage

The test suite validates:

1. **VM Creation**: Verifies the virtual machine is created successfully
2. **NIC Existence**: Confirms the network interface exists
3. **NIC Connection**: Validates the NIC is properly attached to the VM
4. **VM Size**: Checks the VM is using the correct size (Standard_B1s)
5. **Ubuntu Version**: Verifies the VM is running Ubuntu 22.04 LTS (Jammy)

## Test Configuration

Update these variables in `azure_webserver_test.go`:
- `subscriptionID`: Your Azure subscription ID
- `labelPrefix`: Your username prefix for resource naming

## Notes

- Tests will automatically run `terraform init`, `terraform apply`, and `terraform destroy`
- The test typically takes 5-10 minutes to complete
- Ensure you have sufficient Azure credits before running tests
