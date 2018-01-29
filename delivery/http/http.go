package http

import (
	"github.com/dodohq/backdo/delivery/http/admin"
	"github.com/dodohq/backdo/delivery/http/company"
	"github.com/dodohq/backdo/delivery/middleware"
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
	h.Router.GET("/api/admin", middleware.AdminAuthy(adminHandler.Dummy))
	return h
}

// InitCompanyHandler initialize company endpoints
func (h *Handler) InitCompanyHandler(cu usecase.CompanyUsecase) *Handler {
	compHanlder := &company.Handler{CompanyUsecase: cu}
	h.Router.GET("/api/company", middleware.AdminAuthy(compHanlder.GetAllCompanies))
	h.Router.POST("/api/company", middleware.AdminAuthy(compHanlder.OnboardNewCompany))
	h.Router.DELETE("/api/company/:id", middleware.AdminAuthy(compHanlder.DeleteACompany))
	return h
}
