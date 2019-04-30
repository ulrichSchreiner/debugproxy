workflow "Compile" {
  on = "push"
  resolves = ["compile sources"]
}

action "compile sources" {
  uses = "docker://docker.io/golang:1"
  args = "go build -v"
}
