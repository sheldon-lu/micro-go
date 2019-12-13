package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"micros/kubernetes/handler"
	"micros/kubernetes/subscriber"

	kubernetes "micros/kubernetes/proto/kubernetes"

	"encoding/json"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubes.Clientset

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.kubernetes"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	kubernetes.RegisterKubernetesHandler(service.Server(), new(handler.Kubernetes))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), new(subscriber.Kubernetes))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.kubernetes", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Kubernetes() {
	k8sconfig := flag.String("k8sconfig", "k8sconfig", "kubernetes config file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err = kubes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connect k8s success")
		fmt.Println("############################################")
	}

	//获取POD
	pods, err := clientset.CoreV1().Pods("devops").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(pods)
	mj, _ := json.Marshal(pods.Items[1])
	fmt.Println(string(mj))

	fmt.Println(pods.Items[1].Name)
	fmt.Println(pods.Items[1].CreationTimestamp)
	fmt.Println(pods.Items[1].Spec.Containers)
	fmt.Println(pods.Items[1].Labels)
	fmt.Println(pods.Items[1].Namespace)
	fmt.Println(pods.Items[1].Status.HostIP)
	fmt.Println(pods.Items[1].Status.PodIP)
	fmt.Println(pods.Items[1].Status.StartTime)
	fmt.Println(pods.Items[1].Status.Phase)
	fmt.Println(pods.Items[1].Status.ContainerStatuses[0].RestartCount) //重启次数
	fmt.Println(pods.Items[1].Status.ContainerStatuses[0].Image)        //获取重启时间

	//获取NODE
	fmt.Println("##################")
	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	mjson, _ := json.Marshal(nodes.Items[0])
	fmt.Println(string(mjson))

	fmt.Println(nodes.Items[0].Name)
	fmt.Println(nodes.Items[0].CreationTimestamp) //加入集群时间
	fmt.Println(nodes.Items[0].Status.NodeInfo)
	fmt.Println(nodes.Items[0].Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
	fmt.Println(nodes.Items[0].Status.Allocatable.Memory().String())

}
