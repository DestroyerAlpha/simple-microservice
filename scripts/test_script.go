package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DestroyerAlpha/simple-microservice/api/customer"
	"github.com/DestroyerAlpha/simple-microservice/api/restaurant"
	"github.com/DestroyerAlpha/simple-microservice/pkg/config"
	otelPkg "github.com/DestroyerAlpha/simple-microservice/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testAddRestaurant(client restaurant.RestaurantServiceClient) {
	ctx := context.Background()
	_, _ = client.AddRestaurant(ctx, &restaurant.AddRestaurantRequest{
		Name:    "Hooda Halal",
		Cuisine: "Halal",
		FoodItems: []string{
			"Chicken Shawarma",
			"Lamb Gyro",
			"Falafel Wrap",
			"Beef Kebab",
			"Vegetarian Platter",
			"Baklava",
			"Mint Lemonade",
		},
	})
	_, _ = client.AddRestaurant(ctx, &restaurant.AddRestaurantRequest{
		Name:    "Dawat",
		Cuisine: "Indian",
		FoodItems: []string{
			"Butter Chicken",
			"Paneer Tikka",
			"Biryani",
			"Naan",
			"Raita",
		},
	})
	_, _ = client.AddRestaurant(ctx, &restaurant.AddRestaurantRequest{
		Name:    "Khwab",
		Cuisine: "Indian",
		FoodItems: []string{
			"Paneer Butter Masala",
			"Chole Bhature",
			"Dal Makhani",
			"Palak Paneer",
			"Veg Biryani",
			"Roti",
			"Salad",
			"Chutney",
			"Sweet Lassi",
			"Ras Malai",
			"Jalebi",
			"Paneer Tikka",
		},
	})
	resp, err := client.ListRestaurants(ctx, &restaurant.ListRestaurantsRequest{
		Cuisine: "Indian",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func testPlaceCustomerOrder(client customer.CustomerServiceClient) {
	ctx := context.Background()
	const cuisine = "Indian"
	resp, err := client.GetMenu(ctx, &customer.GetMenuRequest{
		Cuisine: cuisine,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	placeOrderResp, err := client.PlaceFoodOrder(ctx, &customer.PlaceFoodOrderRequest{
		FoodItem: resp.GetFoodItems()[0],
		Cuisine:  cuisine,
		Quantity: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(placeOrderResp)
	_, err = client.ReviewFoodItem(ctx, &customer.ReviewFoodItemRequest{
		OrderId:  placeOrderResp.GetOrderId(),
		FoodItem: resp.GetFoodItems()[0],
		Rating:   3,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	otelPkg.InitOtelTracer()
	var dialOpts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler(otelgrpc.WithTracerProvider(otel.GetTracerProvider()))),
	}
	restaurantConn, errRestConn := grpc.NewClient(config.GetServerAddress(config.SERVER_ADDR, config.RESTAURANT_SERVICE_PORT), dialOpts...)
	if errRestConn != nil {
		log.Fatalf("failed to listen to restaurant conn: %v", errRestConn)
	}

	restaurantClient := restaurant.NewRestaurantServiceClient(restaurantConn)
	customerConn, errCustConn := grpc.NewClient(config.GetServerAddress(config.SERVER_ADDR, config.CUSTOMER_SERVICE_PORT), dialOpts...)
	if errCustConn != nil {
		log.Fatalf("failed to listen to restaurant conn: %v", errCustConn)
	}

	customerClient := customer.NewCustomerServiceClient(customerConn)
	testAddRestaurant(restaurantClient)
	testPlaceCustomerOrder(customerClient)
}
