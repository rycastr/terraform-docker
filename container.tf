resource "docker_container" "this" {
  image = docker_image.this.name
  name  = "goapi"
  env = [
    "PORT=8080"
  ]
  ports {
    internal = 8080
    external = 8080
  }
}
