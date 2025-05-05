package main

import (
	"log"
	"net"

	customerPb "github.com/DestroyerAlpha/simple-microservice/api/customer"
	"github.com/DestroyerAlpha/simple-microservice/api/order"
	"github.com/DestroyerAlpha/simple-microservice/api/ratings"
	"github.com/DestroyerAlpha/simple-microservice/customer"
	"github.com/DestroyerAlpha/simple-microservice/pkg/config"
	otelPkg "github.com/DestroyerAlpha/simple-microservice/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.GetServerAddress(config.SERVER_ADDR, config.CUSTOMER_SERVICE_PORT))
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
	orderConn, errOrderConn := grpc.NewClient(config.GetServerAddress(config.ORDER_SERVICE_ADDR, config.ORDER_SERVICE_PORT), dialOpts...)
	if errOrderConn != nil {
		log.Fatalf("failed to listen to order conn: %v", errOrderConn)
	}
	ratingsConn, errRatingsConn := grpc.NewClient(config.GetServerAddress(config.RATINGS_SERVICE_ADDR, config.RATINGS_SERVICE_PORT), dialOpts...)
	if errRatingsConn != nil {
		log.Fatalf("failed to listen to ratings conn: %v", errRatingsConn)
	}
	defer func() {
		_ = orderConn.Close()
		_ = ratingsConn.Close()
	}()

	orderClient := order.NewOrderServiceClient(orderConn)
	ratingsClient := ratings.NewRatingsServiceClient(ratingsConn)
	srv := customer.NewCustomerService(orderClient, ratingsClient)
	grpcServer := grpc.NewServer(grpcServOpts...)
	customerPb.RegisterCustomerServiceServer(grpcServer, srv)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
