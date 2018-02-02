package http

import (
	"github.com/dodohq/backdo/delivery/http/admin"
	"github.com/dodohq/backdo/delivery/http/company"
	"github.com/dodohq/backdo/delivery/http/delivery"
	"github.com/dodohq/backdo/delivery/http/user"
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

// InitUserHandler initialize user endpoints
func (h *Handler) InitUserHandler(uu usecase.UserUsecase) *Handler {
	userHandler := &user.Handler{UserUsecase: uu}
	h.Router.GET("/api/user", middleware.AdminAuthy(userHandler.GetAllExistingAccounts))
	h.Router.GET("/api/user/forgot_password", userHandler.ForgotPassword)
	h.Router.POST("/api/user", middleware.AdminAuthy(userHandler.CreateNewAccount))
	h.Router.POST("/api/user/login", userHandler.Login)
	h.Router.POST("/api/user/reset_password", middleware.UserAuthy(userHandler.ResetPassword))
	h.Router.DELETE("/api/user/:id", middleware.AdminAuthy(userHandler.DeleteAccount))
	return h
}

// InitDeliveryHandler initialize delivery endpoints
func (h *Handler) InitDeliveryHandler(du usecase.DeliveryUsecase) *Handler {
	deliveryHandler := &delivery.Handler{DeliveryUsecase: du}
	h.Router.GET("/api/delivery", middleware.UserAuthy(deliveryHandler.GetAllDeliveries))
	h.Router.POST("/api/delivery", middleware.UserAuthy(deliveryHandler.CreateNewDelivery))
	h.Router.POST("/api/delivery/bulk", middleware.UserAuthy(deliveryHandler.BulkCreateDeliveries))
	h.Router.DELETE("/api/delivery/:id", middleware.UserAuthy(deliveryHandler.DeleteADelivery))
	return h
}
