package service

import (
	"devops_go/config"
	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8s 初始化k8s 的 集群client
var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
}

// Init 初始化k8s
func (k *k8s) Init() {
	// 从flags中构建k8s配置
	config, err := clientcmd.BuildConfigFromFlags("", config.KubeConfig)
	if err != nil {
		panic("创建k8s配置失败, " + err.Error())
	}
	// 根据 cinfig 类型 的对象，new 一个 centset出来
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("创建k8s clientSet失败, " + err.Error())
	} else {
		logger.Info("创建k8s clientSet成功")
	}
	//把clientset 赋值给 k8s
	k.ClientSet = clientSet
}
