package actions


import (
    "log"
    "os"
    "path/filepath"


    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)


var ClientSet *kubernetes.Clientset


func init() {
    log.Println("Start Kubernetes Connection")
    kubeconfig := os.Getenv("KUBECONFIG")
    if kubeconfig == "" {
        kubeconfig = filepath.Join(
            os.Getenv("HOME"), ".kube", "config",
        )
    }
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatal(err)
    }
    ClientSet = kubernetes.NewForConfigOrDie(config)
    
<<<<<<< HEAD
}
=======
}
>>>>>>> 653764f301568c602a930174ba086252666708d8
