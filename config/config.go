package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga el archivo .env y establece las variables de entorno
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// GetEnv obtiene el valor de una variable de entorno, devolviendo un valor por defecto si no existe
func GetEnv(key string, defaultValue string) string {

	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
