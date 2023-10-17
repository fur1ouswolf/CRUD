package handlers

import (
	"gorm.io/gorm"
	"log/slog"
)

type Handler struct {
	db     *gorm.DB
	Logger *slog.Logger
}

func NewHandler(db *gorm.DB, logger *slog.Logger) *Handler {
	return &Handler{
		db:     db,
		Logger: logger,
	}
}
