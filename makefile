# Variables del servicio de pagos (payment-service)
PAYMENT_SERVICE_DIR=./payments/cmd

# Correr servicio de pagos
run-payment-service:
	@echo "Inicializando servidor de pagos"
	@cd $(PAYMENT_SERVICE_DIR) && go run main.go

# Compilar payment-service
build-payment-service:
	@echo "Construyendo servicio de pagos"
	@cd $(PAYMENT_SERVICE_DIR) && go build -o payment main.go
