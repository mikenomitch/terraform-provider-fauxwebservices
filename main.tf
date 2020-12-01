terraform {
  required_providers {
    fauxwebservices = {
      source  = "fauxwebservices"
      version = "0.0.1"
    }
  }
}

provider "fauxwebservices" {
  host = "https://tfe-zone-cc09c2e7.ngrok.io"
  token = "12345"
}

resource "fauxwebservices_server" "test-server" {
  name = "my-demo-server"
}

// resource "fauxwebservices_database" "test-db" {
//   name = "my-demo-db"
//   type = "t1"
// }

// resource "fauxwebservices_storage_bucket" "test-bucket" {
//   name = "my-demo-bucket"
//   size = "50Gb"
// }

// resource "fauxwebservices_vpc" "test-vpc" {
//   name = "my-demo-vpc"
// }
