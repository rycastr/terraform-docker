resource "docker_image" "this" {
  name = "goapi"
  build {
    tag  = ["goapi:latest"]
    path = "services/api"
  }
}
