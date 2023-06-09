package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (h *BaseHandler) HelloWorld(c echo.Context) error {
	//DB with globar var
	if err := h.db.Ping(); err != nil {
		fmt.Println("DB Error")
	}
	return c.String(http.StatusOK, "Hello Ibnu dari Echo")
}
