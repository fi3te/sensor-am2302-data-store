terraform {
  required_version = ">= 1.2"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~>5.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~>2.4.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "~>3.2.1"
    }
  }
}

provider "aws" {
  default_tags {
    tags = {
      Tool = "Terraform"
    }
  }
}
