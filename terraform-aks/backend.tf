terraform {
  backend "azurerm" {
    resource_group_name  = "esgrp"
    storage_account_name = "esgaccount"
    container_name       = "esgcontainer"
    key                  = "terraform.tfstate"
  }
}
