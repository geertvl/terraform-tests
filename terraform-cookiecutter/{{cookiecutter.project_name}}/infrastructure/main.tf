provider azurerm {
    features {}
}

resource "azurerm_resource_group" "example" {
  name = "{{cookiecutter.resource_group_name}}"
  location = "{{cookiecutter.location}}"
}

output "resource_group_name" {
  value = azurerm_resource_group.example.name
}