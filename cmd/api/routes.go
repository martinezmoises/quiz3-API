package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(a.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)

	// Healthcheck
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthCheckHandler)

	// User routes
	router.HandlerFunc(http.MethodPost, "/v1/users", a.createUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", a.getUserHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/users/:id", a.updateUserHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", a.deleteUserHandler)

	return a.recoverPanic(router)
}
