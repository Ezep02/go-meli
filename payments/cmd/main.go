package main

import (
	"log"
	"net/http"

	"github.com/ezep/go-meli/payments/iternal/handler"
	"github.com/ezep/go-meli/payments/iternal/repository"
	"github.com/ezep/go-meli/payments/iternal/service"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

// porque esta se encarga exlusivamente de los pagos con MP
func main() {
	// Configurar Viper
	viper.SetConfigName(".env") // Nombre del archivo de configuración (sin extensión)
	viper.SetConfigType("env")  // Tipo de archivo
	viper.AddConfigPath(".")    // Ruta al directorio de la configuración

	// Intentar leer el archivo de configuración
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	// Cargar variables de entorno
	viper.AutomaticEnv()

	// Definir valores predeterminados
	viper.SetDefault("PAYMENT_SERVICE_PORT", "3000")

	// Crear repositorio, servicio y handler
	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Inicializa el router de Chi
	r := chi.NewRouter()
	handler.PaymentRouter(r, paymentHandler)

	// Obtener el puerto desde Viper
	port := viper.GetString("PAYMENT_SERVICE_PORT")

	// Iniciar el servidor HTTP
	log.Printf("Starting server on port http://localhost:%s...", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
