package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	. "platform-backend/pkg/log"
)

func main() {
	InitLogger()
	defer Logger.Sync()
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//var kubeconfig *string
	//kubeconfig = flag.String("kubeconfig", "./config", "path of kubeconfig file")
	//flag.Parse()
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	svcs, err := clientset.CoreV1().Services("default").List(metav1.ListOptions{})
	if err != nil {
		Logger.Error(err.Error())
		panic(err.Error())
	}
	fmt.Println(svcs.Items)
	totalGpuNum := int64(0)
	for idx := range nodes.Items {
		quantity := nodes.Items[idx].Status.Allocatable["nvidia.com/gpu"]
		gpuNum, _ := quantity.AsInt64()
		totalGpuNum += gpuNum
	}

	Logger.Infof("get gpu num succeed, GPU NUM: %d", totalGpuNum)
}