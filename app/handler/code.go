package handler

import (
	"fmt"
	"net/http"
)

func GetAllCode(w http.ResponseWriter) {
	fmt.Fprint(w, "TESTE")
}
