package main

import (
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"

	"./controller"
)

func main() {
	router := httprouter.New()

	router.GET("/", controller.Index)
	router.POST("/dfs/upload", controller.FileUpload)
	router.GET("/dfs/query/filehash/:fileHash", controller.SingleFileQuery)
	router.GET("/dfs/query/limitcount/:limitCount", controller.BatchFilesQuery)
	router.GET("/dfs/download/:fileHash", controller.FileDownload)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Failed to listen and serve, err: \n" + err.Error())
	}
}
