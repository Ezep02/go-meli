package main

import (
	"log"
	"net/http"

	"github.com/ezep02/auth/internal/database"
	"github.com/ezep02/auth/internal/handler"
	"github.com/ezep02/auth/internal/repository"
	"github.com/ezep02/auth/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
)

func main() {
	// Cargar archivo .env desde un directorio específico
	viper.SetConfigName(".env") // Nombre del archivo sin extensión
	viper.SetConfigType("env")  // Tipo de archivo
	viper.AddConfigPath(".")    // Directorio actual

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[AUTH] Error reading .env file: %v \n", err)
		return
	}

	// Cargar variables de entorno
	viper.AutomaticEnv()

	// Obtener el puerto y la cadena de conexión de la base de datos
	authPort := viper.GetString("PORT")
	dbConnString := viper.GetString("DATABASE_URL")

	// Establecer conexión a la base de datos
	connection, connectionErr := database.DB_Connection(dbConnString)
	if connectionErr != nil {
		log.Printf("Error connecting to the database: %v", connectionErr)
		return
	}

	// Inicializar el repositorio, servicio y handler
	Auth_repo := repository.NewAuthRepository(connection)
	Auth_service := service.NewAuthService(Auth_repo)
	Auth_handler := handler.NewAuthHandler(Auth_service)

	// Inicializar el enrutador
	r := chi.NewRouter()

	// Configuración de CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	handler.AuthRouter(r, Auth_handler)

	// Iniciar el servidor
	log.Printf("Starting server on auth port http://localhost:%s...", authPort)
	if err := Start(":"+authPort, r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func Start(addr string, r *chi.Mux) error {
	return http.ListenAndServe(addr, r)
}
