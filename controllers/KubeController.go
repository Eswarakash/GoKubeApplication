package kubecontroller

import (
	kubeservice "GoKubeAPI/services"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetPodsList(c *gin.Context) {

	c.JSON(200, kubeservice.GetPodsService())
}

func GetSvcList(c *gin.Context) {
	c.JSON(200, kubeservice.GetSevices())
}

func GetNamespace(c *gin.Context) {
	c.JSON(http.StatusOK, kubeservice.GetNamespace())
}

func PostDeployment(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Printf(" Payload : %s", body)

	kubeservice.Deploy(string(body))
	c.JSON(200, gin.H{"Status": "Done"})
}

func CreateDeployment(c *gin.Context) {
	var deployment appsv1.Deployment
	c.Bind(&deployment)
	deploymentStatus := kubeservice.CreateDeployment(deployment)
	if deploymentStatus != nil {
		c.JSON(http.StatusInternalServerError, deploymentStatus.Error())
	} else {
		c.JSON(200, gin.H{"Status": "deployed successfully"})
	}

}

func CreateNameSpace(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	kubeservice.CreateNamespace(string(body))
	c.JSON(200, gin.H{"Status": "Success"})
}
