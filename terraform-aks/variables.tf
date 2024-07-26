variable "resource_group_name" {
  type        = string
  description = "The name of the resource group in which to create the AKS cluster."
}

variable "location" {
  type        = string
  description = "The Azure region in which to create the resources."
}

variable "aks_cluster_name" {
  type        = string
  description = "The name of the AKS cluster."
}
