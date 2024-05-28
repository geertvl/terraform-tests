run "resource_group_name" {
    command = apply

    assert {
        condition = azurerm_resource_group.example.name == "rg-test-001"
        error_message = "Invalid resource group name"
    }
}
