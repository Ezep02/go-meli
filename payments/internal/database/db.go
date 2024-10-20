package database

import (
	"fmt"
	"log"

	"github.com/ezep02/payments/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB_Connection() (*gorm.DB, error) {
	// Cargar las variables de entorno automáticamente
	viper.AutomaticEnv()

	// Obtener la conexión a la base de datos desde la variable de entorno
	dbConn := viper.GetString("DATABASE_URL")

	if dbConn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Establecer conexión a la base de datos
	connection, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		fmt.Println("[DB]: Successful connection")
	}

	// Migrar los modelos
	connection.AutoMigrate(&models.Item{})

	return connection, nil
}
