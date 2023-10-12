package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/bcb12/gestionFct/api/middleware/auth"
	"github.com/bcb12/gestionFct/internal/config"
	"github.com/bcb12/gestionFct/internal/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIHandler struct {
	Ctx *config.AppContext
}

func NewRouter(ctx *config.AppContext) *chi.Mux {
	r := chi.NewRouter()
	apiHandler := &APIHandler{Ctx: ctx}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Public routes
	r.Route("/", func(r chi.Router) {
		r.Post("/api/login", apiHandler.loginHandler)
		r.Post("/api/signin", apiHandler.signInHandler)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/api/test", apiHandler.homeHandler)
	})

	return r
}

func (ah *APIHandler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := helpers.ReadJSON(w, r, &requestPayload)
	if err != nil {
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
	}

	// Validate the user against the database
	user, err := auth.GetByEmail(ah.Ctx.DB, requestPayload.Email)

	if err != nil {
		log.Println(err)
		helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		helpers.ErrorJSON(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	payload := helpers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	helpers.WriteJSON(w, http.StatusAccepted, payload)
	//w.Write([]byte("Login"))
}

func (ah *APIHandler) signInHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}

func (ah *APIHandler) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home"))
}
