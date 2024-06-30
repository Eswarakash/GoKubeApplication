package kubemodel

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type PodDetails struct {
	PodName       string      `json:"podname"`
	NameSpace     string      `json:"namespace"`
	NodeName      string      `json:"nodeName"`
	Status        string      `json:"status"`
	IP            string      `json:"ip"`
	ContainerName string      `json:"containerName"`
	Image         string      `json:"image"`
	Dod           metav1.Time `json:"deployedDate"`
}

type SvcDetails struct {
	SvcName    string             `json:"servicename"`
	IP         string             `json:"ip"`
	Port       int32              `json:"port"`
	TargetPort intstr.IntOrString `json:"targetport"`
	NodePort   int32              `json:"nodeport"`
	Protocal   v1.Protocol        `json:"protocol"`
	Type       v1.ServiceType     `json:"type"`
}
