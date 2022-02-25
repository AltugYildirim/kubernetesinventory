package main

import (
    "flag"
    "fmt"
    "github.com/altugyildirim/actions"
    
)
func main() {
    jobName := flag.String("jobname", "test-job", "The name of the job")
    containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
    entryCommand := flag.String("command", "ls", "The command to run inside the container")

    flag.Parse()

    fmt.Printf("Args : %s %s %s\n", *jobName, *containerImage, *entryCommand)
}
