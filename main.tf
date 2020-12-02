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
}

resource "fauxwebservices_server" "test-server" {
  name = "Server Alpha"
}

resource "fauxwebservices_database" "test-database" {
  name = "Prod DB"
}

// ==========
// Uncomment the resources below and rerun
// `terraform apply` to see updates.
// ==========

resource "fauxwebservices_server" "test-server-2" {
  name = "Server Bravo"
}

resource "fauxwebservices_server" "test-server-3" {
  name = "Server Charlie"
}

resource "fauxwebservices_database" "test-database-2" {
  name = "Replica DB"
}

resource "fauxwebservices_bucket" "test-bucket" {
  name = "Easy File Storage"
}

output "message" {
  value = "You ran a Terraform apply! View your example resources at: https://tfe-zone-cc09c2e7.ngrok.io/fws"
}
