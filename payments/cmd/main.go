package main

import (
	"log"
	"net/http"

	"github.com/ezep02/payments/internal/handler"
	"github.com/ezep02/payments/internal/repository"
	"github.com/ezep02/payments/internal/service"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")

	// Intentar leer el archivo de configuración
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("[PAYMENT] Error reading config file: %v\n", err)
	}

	// Cargar variables de entorno
	viper.AutomaticEnv()

	paymentPort := viper.GetString("PORT")

	// Crear repositorio, servicio y handler
	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Inicializa el router de Chi
	r := chi.NewRouter()
	handler.PaymentRouter(r, paymentHandler)

	log.Printf("Starting server on port http://localhost:%s...", paymentPort)

	// Iniciar el servidor HTTP
	if err := http.ListenAndServe(":"+paymentPort, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
