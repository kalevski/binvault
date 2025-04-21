package api

import (
	"binvault/pkg/api/handlers"
	"binvault/pkg/api/helpers"
	"binvault/pkg/api/middlewares"

	"github.com/julienschmidt/httprouter"
)

func initRouter() *httprouter.Router {
	router := httprouter.New()

	routerMiddlewares := []helpers.Middleware{}
	routerMiddlewares = append(routerMiddlewares, helpers.LoggingMiddleware)
	routerMiddlewares = append(routerMiddlewares, middlewares.AuthMiddleware)

	router.GET("/api/buckets", helpers.ApplyMiddleware(handlers.BucketGetMany, routerMiddlewares))
	router.POST("/api/buckets", helpers.ApplyMiddleware(handlers.BucketCreate, routerMiddlewares))
	router.GET("/api/buckets/:bucketName", handlers.BucketGetOne)
	router.DELETE("/api/buckets/:bucketName", handlers.BucketDelete)

	router.GET("/api/buckets/:bucketName/files", handlers.FileGetMany)
	router.POST("/api/buckets/:bucketName/files", handlers.FileCreate)
	router.GET("/api/buckets/:bucketName/files/:fileId", handlers.FileGetOne)
	router.DELETE("/api/buckets/:bucketName/files/:fileId", handlers.FileDelete)

	return router
}
