package config

import (
	"strconv"
)

const (
	SERVER_ADDR = "0.0.0.0"

	CUSTOMER_SERVICE_ADDR   = "customer"
	ORDER_SERVICE_ADDR      = "order"
	RATINGS_SERVICE_ADDR    = "ratings"
	RESTAURANT_SERVICE_ADDR = "restaurant"
	CUSTOMER_SERVICE_PORT   = 50051
	ORDER_SERVICE_PORT      = 50052
	RATINGS_SERVICE_PORT    = 50053
	RESTAURANT_SERVICE_PORT = 50054
)

func GetServerAddress(serverAddress string, port int) string {
	return serverAddress + ":" + strconv.Itoa(port)
}
