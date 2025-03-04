package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/models"
	"binvault/pkg/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BucketCreateInput struct {
	Name       string            `json:"name" validate:"required,slug,max=30"`
	Visibility models.Visibility `json:"visibility" validate:"required,oneof=public private"`
}

// GET /buckets
func BucketGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pagination := helpers.GetRequestPagination(r)
	buckets := services.BucketGetMany(pagination.Limit, pagination.Offset)
	helpers.JSONResponse(w, http.StatusOK, buckets)
}

// POST /buckets
func BucketCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input BucketCreateInput
	if err := helpers.DecodeJSONBody(r, &input); err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := helpers.GetUserID(r)
	bucket, err := services.BucketCreate(input.Name, input.Visibility, *userId)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.JSONResponse(w, http.StatusCreated, bucket)
}

// GET /buckets/:bucketName
func BucketGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	bucket, err := services.BucketGetOne(bucketName)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.JSONResponse(w, http.StatusOK, bucket)
}

// DELETE /buckets/:bucketName
func BucketDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
