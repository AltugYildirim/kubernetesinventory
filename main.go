package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

type Collection struct {
	Name       string
	Namespaces string
	Replicas   *int32
	//UID string
	Image    string
	Version  string
	NodeName string
	HostIP   string
	Ready    bool
}

func connectToK8s() *kubernetes.Clientset {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatalln("failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create K8s clientset")
	}

	return clientset
}

func List(clientset *kubernetes.Clientset, resource *string, namespace *string) {

	if *resource == "application" {
		listApplicationsDetails(clientset, namespace)

	} else {
		listPods(clientset, namespace)
	}
}

func listApplicationsDetails(clientset *kubernetes.Clientset, namespace *string) {
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	deploycl := clientset.AppsV1().Deployments(*namespace)
	deployments, err := deploycl.List(context.TODO(), metav1.ListOptions{})
	var collectionAdding Collection

	if err != nil {
		panic(err)
	}
	for _, item := range deployments.Items {
		// log.Println("Deploy Name: ", item.GetName())
		// log.Println("NameSpace: ", item.GetNamespace())
		// log.Println("UID: ", item.GetUID())
		//log.Println("UID: ", item.Spec.Replicas)

		// log.Println("Image: ", item.Spec.Template.Spec.Containers[0].Image)
		// log.Println("Version: ", strings.Split(item.Spec.Template.Spec.Containers[0].Image, ":")[1])
		// log.Println("-----------------")

		collectionAdding.Name = item.GetName()
		collectionAdding.Namespaces = item.GetNamespace()
		collectionAdding.Replicas = item.Spec.Replicas
		collectionAdding.Image = item.Spec.Template.Spec.Containers[0].Image
		collectionAdding.Version = strings.Split(item.Spec.Template.Spec.Containers[0].Image, ":")[1]

		// collectionAdding = Collection{
		// 	Name:       item.GetName(),
		// 	Namespaces: item.GetNamespace(),
		// 	Image:      item.Spec.Template.Spec.Containers[0].Image,
		// 	Version:    strings.Split(item.Spec.Template.Spec.Containers[0].Image, ":")[1],
		// }

	}

	if err != nil {
		panic(err)
	}
	for _, item := range pods.Items {
		log.Println("Pod Name: ", item.GetName())
		log.Println("NameSpace: ", item.GetNamespace())
		log.Println("Image: ", item.Spec.NodeName)
		log.Println("Image: ", item.Status.HostIP)
		var readyStatus bool
		for _, status := range item.Status.ContainerStatuses {
			log.Println("Status Ready: ", status.Ready)
			readyStatus = status.Ready
		}
		for _, container := range item.Spec.Containers {
			log.Println("Container Image: ", container.Image)
		}
		log.Println("-----------------")
		collectionAdding.HostIP = item.Status.HostIP
		collectionAdding.NodeName = item.Spec.NodeName
		collectionAdding.Ready = readyStatus
		// collectionAdding = Collection{
		// 	NodeName: item.Spec.NodeName,
		// 	HostIP:   item.Status.HostIP,
		// 	Ready:    readyStatus,
		// }
	}
	var postBody []byte
	postBody, err = json.Marshal(collectionAdding)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(postBody))

}

func listPods(clientset *kubernetes.Clientset, namespace *string) {
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range pods.Items {
		log.Println("Pod Name: ", item.GetName())
		log.Println("NameSpace: ", item.GetNamespace())
		log.Println("Image: ", item.Spec.NodeName)
		log.Println("Image: ", item.Status.HostIP)
		// for _, status := range item.Status.ContainerStatuses {
		// 	log.Println("Status Ready: ", status.Ready)
		// }
		// for _, container := range item.Spec.Containers {
		// 	log.Println("Container Image: ", container.Image)
		// }
		log.Println("-----------------")
	}
}

func main() {
	//jobName := flag.String("jobname", "test-job", "The name of the job")
	//containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	//entryCommand := flag.String("command", "ls", "The command to run inside the container")
	ns := flag.String("ns", "default", "The name of the namespace")
	resource := flag.String("rs", "application", "Resource is a string")

	flag.Parse()

	clientset := connectToK8s()
	//launchK8sJob(clientset, jobName, containerImage, entryCommand)
	List(clientset, resource, ns)
}
