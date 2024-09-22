package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ezep/go-meli/payments/iternal/service"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

type PaymentHandler struct {
	ctx            context.Context
	PaymentService *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		ctx:            context.Background(),
		PaymentService: service,
	}
}

func (pay_handler *PaymentHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	// Si el pago es aceptado, se crea la orden
	accessToken := "APP_USR-196506190136225-092022-41af146cb6426644ccd360b92edc7ef6-1432087693"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:       "Zapatillas Nike",
				Quantity:    1,
				UnitPrice:   100,
				ID:          "231",
				CurrencyID:  "ARS",
				Description: "las verdaderas",
			},
		},
		//callbacks URLs
		BackURLs: &preference.BackURLsRequest{
			Success: "http://localhost:3000/payment/success",
			Pending: "http://localhost:3000/payment/pending",
			Failure: "http://localhost:3000/payment/failure",
		},
		NotificationURL: "https://acd1-181-16-120-161.ngrok-free.app/payment/webhook", // envia a la url un estado del pago
	}

	client := preference.NewClient(cfg)

	resource, err := client.Create(pay_handler.ctx, request)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Convertir el resourse a JSON y devolverlo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}

func (pay_handler *PaymentHandler) Success(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")
}

func (pay_handler *PaymentHandler) Failure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Failure")
}

func (pay_handler *PaymentHandler) Pending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Pending")
}

func (pay_handler *PaymentHandler) WebHook(w http.ResponseWriter, r *http.Request) {

	// Leer el cuerpo de la solicitud
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error leyendo el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Imprimir el cuerpo crudo (opcional, para depuración)
	log.Printf("Cuerpo recibido (raw): %s", string(body))

	// Decodificar el cuerpo JSON en un mapa de interfaces
	var bodyData map[string]interface{}
	if err := json.Unmarshal(body, &bodyData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Acceder al campo "data.id"
	data, ok := bodyData["data"].(map[string]interface{}) // Asegurarse de que "data" es un mapa
	if !ok {
		http.Error(w, "Error: 'data' field is missing or invalid", http.StatusBadRequest)
		return
	}

	idStr, ok := data["id"].(string) // Asegurarse de que "id" es una cadena de texto
	if !ok {
		http.Error(w, "Error: 'id' field is missing or invalid", http.StatusBadRequest)
		return
	}

	// Convertir el ID de cadena a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Error: 'id' is not a valid integer", http.StatusBadRequest)
		return
	}

	// Imprimir el ID convertido a entero
	log.Printf("ID de la transacción (entero): %d", id)

	// Si el pago es aceptado, se crea la orden
	accessToken := "APP_USR-196506190136225-092022-41af146cb6426644ccd360b92edc7ef6-1432087693"

	// Configuración para el cliente de pagos (esto es solo un ejemplo)
	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(cfg)

	if bodyData["type"] == "payment" {
		log.Printf("[ PAGO ACREDITADO ] Payment status accepted")
		client.Get(pay_handler.ctx, id)
	} else {
		log.Printf("[ ATENCION ] no se ejecuto la funcion")
	}

	// Aquí utilizamos el ID convertido a entero para realizar una acción

	// Respuesta exitosa
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "webhook received"})
}
