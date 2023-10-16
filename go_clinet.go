package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "E:\\Code\\Go\\devops_go\\k8s.yml")
	if err != nil {
		panic(err.Error())
	}
	// 根据 cinfig 类型 的对象，new 一个 centset出来
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//使用clientset 获取pod 列表
	podlist, err := clientset.CoreV1().Pods("dev").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range podlist.Items {
		fmt.Println(pod.Name, pod.Namespace)
	}
}
