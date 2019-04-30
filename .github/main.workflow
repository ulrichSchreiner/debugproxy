workflow "Compile" {
  on = "push"
  resolves = ["docker://golang:1"]
}

action "docker://golang:1" {
  uses = "docker://golang:1"
}
