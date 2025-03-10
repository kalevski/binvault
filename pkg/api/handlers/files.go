package handlers

import (
	"binvault/pkg/api/helpers"
	"binvault/pkg/services/files"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GET /bucket/:bucketName/files
func FileGetMany(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pagination := helpers.GetRequestPagination(r)
	bucketName := params.ByName("bucketName")
	files := files.FileGetMany(bucketName, pagination.Limit, pagination.Offset)
	helpers.SendJSON(w, http.StatusOK, files)
}

// POST /bucket/:bucketName/files
func FileCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Unable to parse multipart form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, "Error retrieving the file")
		return
	}
	defer file.Close()

	strict := r.FormValue("strict") == "true"
	mimetype := header.Header.Get("Content-Type")
	if !helpers.IsMimeTypeAllowed(mimetype) {
		helpers.SendError(w, http.StatusBadRequest, "File type is not allowed")
		return
	}

	content := make([]byte, header.Size)
	_, err = file.Read(content)
	if err != nil && err != io.EOF {
		helpers.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdFile, err := files.FileCreate(bucketName, header.Filename, content, strict)
	if err != nil {
		helpers.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SendJSON(w, http.StatusCreated, createdFile)
}

// GET /bucket/:bucketName/files/:fileId
func FileGetOne(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bucketName := params.ByName("bucketName")
	fileId := params.ByName("fileId")
	file, err := files.FileGetOne(bucketName, fileId)
	if err != nil {
		helpers.SendError(w, http.StatusNotFound, err.Error())
	}
	helpers.SendJSON(w, http.StatusOK, file)
}

// GET /bucket/:bucketName/files/:fileId/content
func FileGetContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// DELETE /bucket/:bucketName/files/:fileId
func FileDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
