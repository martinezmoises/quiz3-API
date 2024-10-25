// Filename: cmd/api/users.go

package main

import (
	"errors"
	"net/http"

	"github.com/martinezmoises/quiz3/internal/data"
	"github.com/martinezmoises/quiz3/internal/validator"
)

func (a *applicationDependencies) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user data.User
	v := validator.New()

	err := a.readJSON(w, r, &user)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	data.ValidateUser(v, &user)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.userModel.Insert(&user) // Use userModel here
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	a.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
}

func (a *applicationDependencies) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	user, err := a.userModel.Get(id) // Use userModel here
	if err != nil {
		if errors.Is(err, data.ErrUserNotFound) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	a.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
}

func (a *applicationDependencies) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	var user data.User
	user.ID = id

	err = a.readJSON(w, r, &user)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	data.ValidateUser(v, &user)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.userModel.Update(&user) // Use userModel here
	if err != nil {
		if errors.Is(err, data.ErrUserNotFound) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	a.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
}

func (a *applicationDependencies) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	err = a.userModel.Delete(id) // Use userModel here
	if err != nil {
		if errors.Is(err, data.ErrUserNotFound) {
			a.notFoundResponse(w, r)
		} else {
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
