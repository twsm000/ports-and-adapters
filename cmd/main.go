package main

import (
	"appstruct/internal/adapters/app/api"
	"appstruct/internal/adapters/core/arithmetic"
	"appstruct/internal/adapters/framework/left/grpc"
	"appstruct/internal/adapters/framework/right/db"
	"appstruct/internal/ports"
	"log"
	"os"
)

var (
	err            error
	apiPort        ports.APIPort
	arithmeticPort ports.ArithmeticPort
	dbPort         ports.DBPort
	grpcPort       ports.GRPCPort
	dbDriver       string
	dbDSN          string
)

func main() {
	// framework - driven by application
	dbPort = getNewDBPortInitialized()
	defer dbPort.Close()

	// domain
	arithmeticPort = arithmetic.NewAdapter()

	// application
	apiPort = api.NewAdapter(dbPort, arithmeticPort)
	grpcPort = grpc.NewAdapter(apiPort)

	// framework - driver of the application
	grpcPort.Run()
}

func getNewDBPortInitialized() ports.DBPort {
	dbDriver = os.Getenv("DB_DRIVER")
	dbDSN = os.Getenv("DB_DSN")
	adapter, err := db.NewAdapter(dbDriver, dbDSN)
	if err != nil {
		log.Fatalln("Error creating db adapter:", err)
	}

	return adapter
}
