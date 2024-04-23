terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.66.0"
    }
  }

  required_version = ">= 0.14"
}

variable "project_id" {
  description = "project id"
}

variable "region" {
  description = "region"
}

variable "zone" {
  description = "zone"
}

provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
  credentials = file("gcp_key.json")
}

# VPC
resource "google_compute_network" "tf-vpc" {
  name                    = "${var.project_id}-tf-vpc"
  auto_create_subnetworks = "false"
}

resource "random_id" "bucket_prefix" {
  byte_length = 8
}

resource "google_storage_bucket" "tf-state-bucket-demo" {
  name     = "tf-state-bucket-demo-${random_id.bucket_prefix.hex}"
  force_destroy = false
  location      = "US"
  storage_class = "STANDARD"
  versioning {
    enabled = true
  }
}

#output "tf_state_bucket" {
#  value = google_storage_bucket.tf-state-bucket.name
#}