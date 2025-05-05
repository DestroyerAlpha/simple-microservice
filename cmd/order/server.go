package main

import (
	"log"
	"net"

	orderPb "github.com/DestroyerAlpha/simple-microservice/api/order"
	restaurantPb "github.com/DestroyerAlpha/simple-microservice/api/restaurant"
	"github.com/DestroyerAlpha/simple-microservice/order"
	"github.com/DestroyerAlpha/simple-microservice/pkg/config"
	otelPkg "github.com/DestroyerAlpha/simple-microservice/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", config.GetServerAddress(config.SERVER_ADDR, config.ORDER_SERVICE_PORT))
	if err != nil {
		log.Fatalf("failed to listen to tcp conn: %v", err)
	}
	var grpcServOpts = []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler(
			otelgrpc.WithPropagators(otel.GetTextMapPropagator()),
			otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
		)),
	}
	otelPkg.InitOtelTracer()
	var dialOpts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))),
	}
	restaurantConn, errRestConn := grpc.NewClient(config.GetServerAddress(config.RESTAURANT_SERVICE_ADDR, config.RESTAURANT_SERVICE_PORT), dialOpts...)
	if errRestConn != nil {
		log.Fatalf("failed to listen restaurant conn: %v", errRestConn)
	}
	defer func() {
		_ = restaurantConn.Close()
	}()
	restaurantClient := restaurantPb.NewRestaurantServiceClient(restaurantConn)
	srv := order.NewOrderService(restaurantClient)
	grpcServer := grpc.NewServer(grpcServOpts...)
	orderPb.RegisterOrderServiceServer(grpcServer, srv)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
