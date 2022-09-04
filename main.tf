terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.20.3"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}
