package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/models"
	"binvault/pkg/services/buckets"
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
	buckets := buckets.BucketGetMany(pagination.Limit, pagination.Offset)
	helpers.SendJSON(w, http.StatusOK, buckets)
}

// POST /buckets
func BucketCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input BucketCreateInput
	if err := helpers.DecodeJSONBody(r, &input); err != nil {
		helpers.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := helpers.GetUserID(r)
	bucket, err := buckets.BucketCreate(input.Name, input.Visibility, *userId)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.SendJSON(w, http.StatusCreated, bucket)
}

// GET /buckets/:bucketName
func BucketGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	bucket, err := buckets.BucketGetOne(bucketName)
	if err != nil {
		helpers.SendError(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.SendJSON(w, http.StatusOK, bucket)
}

// DELETE /buckets/:bucketName
func BucketDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	err := buckets.BucketDelete(bucketName)
	if err != nil {
		helpers.SendError(w, http.StatusNotFound, err.Error())
		return
	}
	// schedule task for removing files
	helpers.SendJSON(w, http.StatusAccepted, &helpers.OperationResult{
		Success: true,
		Message: "bucket deleted successfully",
	}) //
}
