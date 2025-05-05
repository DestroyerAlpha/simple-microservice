package main

import (
	"log"
	"net"

	restaurantPb "github.com/DestroyerAlpha/simple-microservice/api/restaurant"
	"github.com/DestroyerAlpha/simple-microservice/pkg/config"
	otelPkg "github.com/DestroyerAlpha/simple-microservice/pkg/otel"
	"github.com/DestroyerAlpha/simple-microservice/restaurant"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.GetServerAddress(config.SERVER_ADDR, config.RESTAURANT_SERVICE_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var grpcServOpts = []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler(
			otelgrpc.WithPropagators(otel.GetTextMapPropagator()),
			otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
		)),
	}
	otelPkg.InitOtelTracer()
	srv := restaurant.NewRestaurantService()
	grpcServer := grpc.NewServer(grpcServOpts...)
	restaurantPb.RegisterRestaurantServiceServer(grpcServer, srv)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
