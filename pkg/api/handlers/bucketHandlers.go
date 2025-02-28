package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/services"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /buckets
func BucketGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("BucketGetMany")
	pagination := helpers.GetRequestPagination(r)
	buckets := services.BucketGetMany(pagination.Limit, pagination.Offset)
	log.Println(buckets)
	helpers.JSONResponse(w, http.StatusOK, buckets)
}

// POST /buckets
func BucketCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// GET /buckets/:bucketName
func BucketGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /buckets/:bucketName
func BucketDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
