package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

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
func launchK8sJob(clientset *kubernetes.Clientset, jobName *string, image *string, cmd *string) {
	jobs := clientset.BatchV1().Jobs("default")
	var backOffLimit int32 = 0

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      *jobName,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    *jobName,
							Image:   *image,
							Command: strings.Split(*cmd, " "),
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job.")
	}

	//print job details
	log.Println("Created K8s job successfully")
}

func List(clientset *kubernetes.Clientset, resource *string, namespace *string) {

	if *resource == "deploy" {
		listDeployments(clientset, namespace)
	} else {
		listPods(clientset, namespace)
	}
}

func listDeployments(clientset *kubernetes.Clientset, namespace *string) {
	deploycl := clientset.AppsV1().Deployments(*namespace)
	deployments, err := deploycl.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range deployments.Items {
		log.Println("Deploy Name: ", item.GetName())
		log.Println("NameSpace: ", item.GetNamespace())
		log.Println("UID: ", item.GetUID())
		log.Println("-----------------")
	}
}

func listPods(clientset *kubernetes.Clientset, namespace *string) {
	pods, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, item := range pods.Items {
		log.Println("Pod Name: ", item.GetName())
		log.Println("NameSpace: ", item.GetNamespace())
		for _, status := range item.Status.ContainerStatuses {
			log.Println("Status Ready: ", status.Ready)
		}
		for _, container := range item.Spec.Containers {
			log.Println("Container Image: ", container.Image)
		}
		log.Println("-----------------")
	}
}

func main() {
	//jobName := flag.String("jobname", "test-job", "The name of the job")
	//containerImage := flag.String("image", "ubuntu:latest", "Name of the container image")
	//entryCommand := flag.String("command", "ls", "The command to run inside the container")
	ns := flag.String("ns", "default", "The name of the namespace")
	resource := flag.String("rs", "deploy", "Resource is a string")

	flag.Parse()

	clientset := connectToK8s()
	//launchK8sJob(clientset, jobName, containerImage, entryCommand)
	List(clientset, resource, ns)
}

