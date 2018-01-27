package http

import (
	"github.com/dodohq/backdo/delivery/http/admin"
	"github.com/dodohq/backdo/usecase"
	"github.com/julienschmidt/httprouter"
)

// Handler hanlder for http routes
type Handler struct {
	Router *httprouter.Router
}

// InitAdminHandler initialize admin endpoints
func (h *Handler) InitAdminHandler(au usecase.AdminUsecase) *Handler {
	adminHandler := &admin.Handler{AdminUsecase: au}
	h.Router.POST("/api/admin/login", adminHandler.Login)
	return h
}
