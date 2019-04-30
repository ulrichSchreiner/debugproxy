workflow "Compile" {
  on = "push"
  resolves = ["docker://docker.io/golang:1"]
}

action "docker://golang:1" {
  uses = "docker://golang:1"
}
