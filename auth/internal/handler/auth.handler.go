package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ezep02/auth/internal/models"
	"github.com/ezep02/auth/internal/service"
	"github.com/ezep02/auth/internal/utils"
)

type AuthHandler struct {
	ctx         context.Context
	AuthService *service.AuthService
}

func NewAuthHandler(AuthService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		ctx:         context.Background(),
		AuthService: AuthService,
	}
}

func (h *AuthHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	//encriptacion del password
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.Password = hash
	user.Name = strings.ToLower(user.Name)

	//se registra el usuario
	u, err := h.AuthService.SignUpService(h.ctx, &user)

	if err != nil {
		http.Error(w, "error creating user", http.StatusBadRequest)
		return
	}

	// si el registro fue exitoso, se crea un token
	tokenString, err := utils.GenerateToken(u.ID, u.Is_admin, u.Name, u.Email, u.Surname)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Establece la cookie con el token
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expira en 24 horas
		Domain:   "",                             // Usa el dominio actual por defecto
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // Cambiar a true si usas HTTPS
		Path:     "/",
	})

	//si todo va bien se devuelve el header con la respuesta
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ID":         u.ID,
		"Name":       u.Name,
		"email":      u.Email,
		"Is_admin":   u.Is_admin,
		"Created_At": u.CreatedAt,
		"surname":    u.Surname,
		"token":      tokenString,
	})
}

func (h *AuthHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.LoginUserReq

	// caputurar el body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	// ejecutar la peticion al repo
	userLogged, err := h.AuthService.SignInService(h.ctx, &user)
	if err != nil {
		http.Error(w, "Error loging user", http.StatusBadRequest)
		return
	}

	// comparar las contraseñas
	if err := utils.HashCompare(user.Password, userLogged.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLogged.Name = strings.ToLower(userLogged.Name)

	// Creación del token con conversión de uint a int
	tokenString, err := utils.GenerateToken(userLogged.ID, userLogged.IsAdmin, userLogged.Name, userLogged.Email, userLogged.Surname)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// obtengo el token y lo seteo en el navegador
	// Establece la cookie con el token
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expira en 24 horas
		Domain:   "",                             // Usa el dominio actual por defecto
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // Cambiar a true si usas HTTPS
		Path:     "/",
	})

	//si todo va bien se devuelve el header con la respuesta
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":         userLogged.ID,
		"name":       userLogged.Name,
		"email":      userLogged.Email,
		"is_admin":   userLogged.IsAdmin,
		"surname":    userLogged.Surname,
		"created_At": userLogged.CreatedAt,
		"token":      tokenString,
	})
}

func VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("auth_token")

	if err != nil {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}
	// Validar el token
	tokenString := cookie.Value

	user, err := utils.VerfiyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func LogoutSession(w http.ResponseWriter, r *http.Request) {

	c := http.Cookie{
		Name:     "auth_token",
		MaxAge:   -1,
		Path:     "/",  // Asegúrate de que el Path coincida con el de la cookie original
		HttpOnly: true, // Evita que la cookie sea accesible desde JavaScript
		Secure:   true, // Solo permite que se envíe por HTTPS
	}

	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
}
