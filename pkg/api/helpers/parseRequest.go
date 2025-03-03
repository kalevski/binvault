package helpers

import (
	"binvault/pkg/auth"
	"net/http"
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

func GetUserID(r *http.Request) *map[string]any {
	token := GetRequestToken(r)
	if token == nil {
		return nil
	}
	claims := auth.GetClaims(*token)
	return &claims
}
