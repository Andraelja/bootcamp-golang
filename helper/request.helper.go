package helper

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// ambil ID dari url, kemudian ubah jadi integer
func ParseID(r *http.Request, prefix string)(int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	return strconv.Atoi(idStr)
}

// ubah json jadi struct go
func DecodeJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}