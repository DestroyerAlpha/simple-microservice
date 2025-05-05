package main

import (
	"log"
	"net"

	orderPb "github.com/DestroyerAlpha/simple-microservice/api/order"
	ratingsPb "github.com/DestroyerAlpha/simple-microservice/api/ratings"
	restaurantPb "github.com/DestroyerAlpha/simple-microservice/api/restaurant"
	"github.com/DestroyerAlpha/simple-microservice/pkg/config"
	otelPkg "github.com/DestroyerAlpha/simple-microservice/pkg/otel"
	"github.com/DestroyerAlpha/simple-microservice/ratings"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", config.GetServerAddress(config.SERVER_ADDR, config.RATINGS_SERVICE_PORT))
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
		log.Fatalf("failed to listen to restaurant conn: %v", errRestConn)
	}
	orderConn, errOrderConn := grpc.NewClient(config.GetServerAddress(config.ORDER_SERVICE_ADDR, config.ORDER_SERVICE_PORT), dialOpts...)
	if errOrderConn != nil {
		log.Fatalf("failed to listen to order conn: %v", errOrderConn)
	}
	defer func() {
		_ = restaurantConn.Close()
		_ = orderConn.Close()
	}()
	restaurantClient := restaurantPb.NewRestaurantServiceClient(restaurantConn)
	orderClient := orderPb.NewOrderServiceClient(orderConn)
	srv := ratings.NewRatingsService(restaurantClient, orderClient)
	grpcServer := grpc.NewServer(grpcServOpts...)
	ratingsPb.RegisterRatingsServiceServer(grpcServer, srv)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
