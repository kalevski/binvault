package helpers

import (
	"binvault/pkg/env"
	"binvault/pkg/services/auth"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type RequestPagination struct {
	Limit  int
	Offset int
}

func GetRequestPagination(r *http.Request) *RequestPagination {
	query := r.URL.Query()
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		offset = 0
	}

	if (limit < 1) || (limit > 100) {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return &RequestPagination{limit, offset}
}

func GetRequestToken(r *http.Request) *string {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil
	}

	if len(token) < 7 || token[:7] != "Bearer " {
		return nil
	}
	tokenStr := token[7:]
	return &tokenStr
}

var SystemUserID = "SYSTEM_USER"

func GetUserID(r *http.Request) *string {
	token := GetRequestToken(r)
	if token == nil {
		return &SystemUserID
	}
	claims := auth.GetClaims(*token)
	id, ok := claims[env.GetVars().JWT_CLAIM_ID]
	if !ok {
		return &SystemUserID
	}

	idType := reflect.TypeOf(id)
	if idType.Kind() == reflect.String {
		return id.(*string)
	}

	if idType.Kind() == reflect.Float64 {
		str := fmt.Sprintf("%.0f", id)
		return &str
	}

	log.Printf("type [%s] is not handled", idType.Kind())
	return &SystemUserID
}

var allowedMimeTypes = []string{
	"image/jpeg",
	"image/png",
}

func IsMimeTypeAllowed(mimetype string) bool {
	for _, allowedType := range allowedMimeTypes {
		if mimetype == allowedType {
			return true
		}
	}
	return false
}
