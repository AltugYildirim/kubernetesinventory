package actions

<<<<<<< HEAD
import (
	"io/ioutil"
	"log"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

func ReadYamlFile(filename string) []runtime.Object {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(dataByte)
	yamlfiles := strings.Split(dataString, "---")
	var resources []runtime.Object

	for _, file := range yamlfiles {
		file = strings.Trim(file, " ")
		file = strings.Trim(file, "\n")
		if file == "\n" || file == "" {
			continue
		}

		decode := scheme.Codecs.UniversalDeserializer().Decode
		obj, _, err := decode([]byte(file), nil, nil)

		if err != nil {
			log.Fatal(err)
			continue
		}
		resources = append(resources, obj)
	}
	return resources
}
=======

import (
    "io/ioutil"
    "log"
    "strings"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/kubernetes/scheme"
)


func ReadYamlFile(filename string) []runtime.Object {
    dataByte, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }


    dataString := string(dataByte)
    yamlfiles := strings.Split(dataString, "---")
    var resources []runtime.Object


    for _, file := range yamlfiles {
        file = strings.Trim(file, " ")
        file = strings.Trim(file, "\n")
        if file == "\n" || file == "" {            
            continue
        }


        decode := scheme.Codecs.UniversalDeserializer().Decode
        obj, _, err := decode([]byte(file), nil, nil)


        if err != nil {
            log.Fatal(err)
            continue
        }
        resources = append(resources, obj)
    }
    return resources
}

>>>>>>> 653764f301568c602a930174ba086252666708d8
