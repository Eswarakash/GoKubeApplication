package kubeservice

import (
	kubemodel "GoKubeAPI/models"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientset *kubernetes.Clientset

func init() {

	var kubeConfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "path of config file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	clientSetinstance, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	clientset = clientSetinstance
}

func GetPodsService() []kubemodel.PodDetails {

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Printf("Pods %s: ", pods)

	var listPods []kubemodel.PodDetails
	//getting the name and namespace of the pods from the cluster
	for _, pod := range pods.Items {

		// listPods = append(listPods, kubepods.NameAndNameSpace{PodName: pod.Name, NameSpace: pod.Namespace})
		for _, container := range pod.Spec.Containers {
			listPod := kubemodel.PodDetails{
				PodName:       pod.Name,
				NameSpace:     pod.Namespace,
				NodeName:      pod.Spec.NodeName,
				Status:        string(pod.Status.Phase),
				IP:            pod.Status.PodIP,
				ContainerName: container.Name,
				Image:         container.Image,
				Dod:           pod.CreationTimestamp,
			}
			listPods = append(listPods, listPod)
		}
	}

	return listPods
}

func GetSevices() []kubemodel.SvcDetails {
	svcs, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	// fmt.Printf("services : %s \n", svcs)

	// fmt.Print("Items : \n", svcs.Items)
	var listSvcs []kubemodel.SvcDetails
	for _, svc := range svcs.Items {
		for _, ports := range svc.Spec.Ports {
			listSvc := kubemodel.SvcDetails{
				SvcName:    svc.Name,
				IP:         svc.Spec.ClusterIP,
				Port:       ports.Port,
				TargetPort: ports.TargetPort,
				NodePort:   ports.NodePort,
				Protocal:   ports.Protocol,
				Type:       svc.Spec.Type,
			}
			listSvcs = append(listSvcs, listSvc)
		}
	}
	return listSvcs
}

func Deploy(deploymentJson string) {

	var deployment appsv1.Deployment
	if err := json.Unmarshal([]byte(deploymentJson), &deployment); err != nil {
		panic(err.Error())
	}

	_, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully deployed")
}

func CreateDeployment(deploymentJson appsv1.Deployment) error {
	_, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), &deploymentJson, metav1.CreateOptions{})
	if err != nil {
		//panic(err.Error())
		return errors.New("Failed with error" + err.Error())
	}

	fmt.Println("Successfully deployed")
	return nil
}

func GetNamespace() []string {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var listNameSpaces []string
	for _, ns := range namespaces.Items {
		//fmt.Println("name : ", ns.Name)
		listNameSpaces = append(listNameSpaces, ns.Name)
	}

	return listNameSpaces
}

func CreateNamespace(nameSpace string) {

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: nameSpace,
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Created NampeSpace %s \n", nameSpace)
}
