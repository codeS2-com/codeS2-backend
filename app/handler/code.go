package handler

import (
	"net/http"
	"database/sql"
	"github.com/codeS2-com/codeS2-backend/app/model"
)


func GetAllCode(DB *sql.DB, w http.ResponseWriter) {

	codes, err := model.FindAll(DB)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, codes)

	// // defer db.Close()
	// stmtIns, err := DB.Prepare("INSERT INTO code(title) VALUES( ? )") // ? = placeholder
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer stmtIns.Close()

	// res, _ := stmtIns.Exec("Olá Mundo")

	// id, err := res.LastInsertId()

	// // json := "seu id é:"+ string(id)
	// respondJSON(w,200,id)
}
