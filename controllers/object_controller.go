package controllers

import (
	"net/http"
	"sync"

	"gin-go-api/models"
	"gin-go-api/services"

	"github.com/gin-gonic/gin"
)

func GetObjects() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		objects := make([]models.Object, 0)

		getObjects(&objects)

		ctx.JSON(http.StatusOK, objects)
	}
}

func getObjects(objects *[]models.Object) {
	wg := sync.WaitGroup{}

	for len(*objects) < 25 {
		missingElements := 25 - len(*objects)
		wg.Add(missingElements)

		for i := 0; i < missingElements; i++ {
			go getObject(&wg, objects)
		}
		wg.Wait()
	}
}

func getObject(wg *sync.WaitGroup, objects *[]models.Object) {

	responseObject := services.GetObjectFromWebService()

	if !contains(*objects, responseObject) {
		*objects = append(*objects, responseObject)
	}

	wg.Done()
}

func contains(objects []models.Object, responseObject models.Object) bool {
	for _, element := range objects {
		if element.Id == responseObject.Id {
			return true
		}
	}
	return false
}
