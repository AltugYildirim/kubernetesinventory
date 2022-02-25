package main

import (
    //"flag"
    "github.com/altugyildirim/kubernetesinventory/actions"
    //"./actions"
)

func main() {


    //cmd := flag.String("cmd", "list", "Should be a string")
    //ns := flag.String("ns", "default", "Namespace is a string")
    //resource := flag.String("rs", "pods", "Resource is a string")
    //file := flag.String("file", "pods", "Resource is a string")


    //flag.Parse()
    actions.List("deploy","default")
    //switch *cmd {
    //case "deploy":
    //    actions.Deploy(file)
    //default:
    //    actions.List(resource, ns)
    //}
}

