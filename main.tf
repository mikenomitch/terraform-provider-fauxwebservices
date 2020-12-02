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
  token = "IZlB19VWu2dBtA.atlasv1.1vz9RztBTM0fk0NhZTv1TWz2jkSdrvyzIdUnHBTht28fUwDbnB7RlQaWhJXuTL0v5DY"
}

resource "fauxwebservices_server" "test-server" {
  name = "my-demo-server"
}

resource "fauxwebservices_server" "test-server-2" {
  name = "my-demo-server-2"
}

resource "fauxwebservices_database" "test-database" {
  name = "my-demo-db"
}

resource "fauxwebservices_database" "test-database-2" {
  name = "my-demo-db-2"
}

resource "fauxwebservices_bucket" "test-bucket" {
  name = "my-demo-bucket"
}

// resource "fauxwebservices_vpc" "test-vpc" {
//   name = "my-demo-vpc"
// }
