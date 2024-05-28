variable "resource_group_name" {
  type = string
}

variable "location" {
  type = string
  default = "West Europe"
}

variable "client_id" {
  description = "The Client ID for the Service Principal."
  type        = string
}

variable "client_secret" {
  description = "The Client Secret for the Service Principal."
  type        = string
}

variable "subscription_id" {
  description = "The Subscription ID for the Azure account."
  type        = string
}

variable "tenant_id" {
  description = "The Tenant ID for the Azure account."
  type        = string
}