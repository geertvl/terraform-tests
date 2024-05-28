provider "azurerm" {
  features {
    
  }
}

resource "azurerm_resource_group" "example" {
  name = "rg-test-001"
  location = "West Europe"
}

output "resource_group_name" {
  value = azurerm_resource_group.example.name
}