package example

default allow = false

# Policy to ensure the resource group is in "West Europe"
allow {
    input.resource_changes[_].change.after.location == "West Europe"
}