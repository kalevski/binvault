package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /bucket/:bucketName/files
func FileGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pagination := helpers.GetRequestPagination(r)
	bucketName := params.ByName("bucketName")
	files := services.FileGetMany(bucketName, pagination.Limit, pagination.Offset)
	helpers.JSONResponse(w, http.StatusOK, files)
}

// POST /bucket/:bucketName/files
func FileCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "unable to parse multipart form")
		return
	}

	// Get the file from the form
	file, header, err := r.FormFile("file")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Error retrieving the file")
		return
	}
	defer file.Close()

	// Call the service to handle the file upload
	createdFile, err := services.FileCreate(bucketName, *header, file)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, createdFile)

}

// GET /bucket/:bucketName/files/:fileId
func FileGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	fileId := params.ByName("fileId")
	file, err := services.FileGetOne(bucketName, fileId)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusNotFound, err.Error())
	}
	helpers.JSONResponse(w, http.StatusOK, file)
}

// GET /bucket/:bucketName/files/:fileId/content
func FileGetContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /bucket/:bucketName/files/:fileId
func FileDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
