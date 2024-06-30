package main

import (
	kubecontroller "GoKubeAPI/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", kubecontroller.GetPong)
	r.GET("/pods", kubecontroller.GetPodsList)
	r.GET("/services", kubecontroller.GetSvcList)
	r.GET("/namespace", kubecontroller.GetNamespace)
	r.POST("/deploy", kubecontroller.PostDeployment)
	r.POST("/createnamespace", kubecontroller.CreateNameSpace)
	r.POST("/createdeploy", kubecontroller.CreateDeployment)

	log.Fatal(http.ListenAndServe(":8080", r))
	r.Run()
}
