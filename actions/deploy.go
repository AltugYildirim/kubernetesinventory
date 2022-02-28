package actions

<<<<<<< HEAD
import (
	"log"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/util/retry"
)

func Deploy(file *string) {
	resources := ReadYamlFile(*file)
	for _, resource := range resources {
		switch rs := resource.(type) {
		case *appsv1.Deployment:
			applyDeployment(rs)
		default:
			log.Println("Only Deployment are allow to be Deploy.")
		}
	}
}

func applyDeployment(deployment *appsv1.Deployment) {
	namespace := deployment.GetNamespace()
	if namespace == "" {
		namespace = "default"
	}
	deployCl := ClientSet.AppsV1().Deployments(namespace)

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		log.Printf("Start %s Deployment", deployment.GetName())
		_, err := deployCl.Create(deployment)
		return err
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Deploy %s Crated", deployment.GetName())
}
=======

import (
    "log"
    appsv1 "k8s.io/api/apps/v1"
    "k8s.io/client-go/util/retry"
)


func Deploy(file *string) {
    resources := ReadYamlFile(*file)
    for _, resource := range resources {
        switch rs := resource.(type) {
        case *appsv1.Deployment:
            applyDeployment(rs)
        default:
            log.Println("Only Deployment are allow to be Deploy.")
        }
    }
}


func applyDeployment(deployment *appsv1.Deployment) {
    namespace := deployment.GetNamespace()
    if namespace == "" {
        namespace = "default"
    }
    deployCl := ClientSet.AppsV1().Deployments(namespace)


    err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
        log.Printf("Start %s Deployment", deployment.GetName())
        _, err := deployCl.Create(deployment)
        return err
    })
    if err != nil {
        panic(err)
    }
    log.Printf("Deploy %s Crated", deployment.GetName())
}

>>>>>>> 653764f301568c602a930174ba086252666708d8
