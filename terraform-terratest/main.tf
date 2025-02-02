provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rgtest" {
  name     = var.resource_group_name
  location = var.location
}

output "resource_group_name" {
  value = azurerm_resource_group.rgtest.name
}
