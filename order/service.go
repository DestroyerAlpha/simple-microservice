package order

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/DestroyerAlpha/simple-microservice/api/order"
	"github.com/DestroyerAlpha/simple-microservice/api/restaurant"

	"github.com/google/uuid"
)

type Service struct {
	order.UnimplementedOrderServiceServer
	orders           []*Order
	restaurantClient restaurant.RestaurantServiceClient
}

func NewOrderService(restaurantClient restaurant.RestaurantServiceClient) *Service {
	return &Service{restaurantClient: restaurantClient, orders: make([]*Order, 0)}
}

type Order struct {
	Id           string
	CustomerId   string
	RestaurantId string
	FoodItem     string
	Quantity     int
}

func (s *Service) PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	log.Printf("Placing order for %s cuisine and %s food item", req.GetCuisine(), req.GetFoodItem())
	resp, err := s.restaurantClient.ListRestaurants(ctx, &restaurant.ListRestaurantsRequest{
		Cuisine: req.GetCuisine(),
	})
	if err != nil {
		return nil, fmt.Errorf("error in getting restaurants: %w", err)
	}
	for _, restaurantId := range resp.GetRestaurantIds() {
		restDetailsResp, errResp := s.restaurantClient.GetRestaurantDetails(ctx, &restaurant.GetRestaurantDetailsRequest{
			RestaurantId: restaurantId,
		})
		if errResp != nil {
			return nil, fmt.Errorf("error in getting restaurant details: %v", errResp)
		}
		for _, foodItem := range restDetailsResp.GetFoodItems() {
			log.Printf("Is food item %s equal to %s", foodItem, req.GetFoodItem())
			if strings.EqualFold(req.GetFoodItem(), foodItem) {
				orderId := uuid.NewString()
				s.orders = append(s.orders, &Order{
					Id:           orderId,
					RestaurantId: restaurantId,
					FoodItem:     foodItem,
					Quantity:     int(req.GetQuantity()),
				})
				log.Println("Returning success")
				return &order.PlaceOrderResponse{
					OrderId: orderId,
					Status:  "success",
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("could not place order as food item %s not found", req.GetFoodItem())
}

func (s *Service) GetMenu(ctx context.Context, req *order.GetMenuRequest) (*order.GetMenuResponse, error) {
	resp, err := s.restaurantClient.ListRestaurants(ctx, &restaurant.ListRestaurantsRequest{
		Cuisine: req.GetCuisine(),
	})
	if err != nil {
		return nil, fmt.Errorf("error in getting restaurants: %w", err)
	}
	var foodItems []string
	for _, restaurantId := range resp.GetRestaurantIds() {
		restDetailsResp, errResp := s.restaurantClient.GetRestaurantDetails(ctx, &restaurant.GetRestaurantDetailsRequest{
			RestaurantId: restaurantId,
		})
		if errResp != nil {
			return nil, fmt.Errorf("error in getting restaurant details: %v", errResp)
		}
		foodItems = append(foodItems, restDetailsResp.GetFoodItems()...)
	}
	return &order.GetMenuResponse{
		FoodItems: foodItems,
	}, nil
}

func (s *Service) GetOrderDetails(_ context.Context, req *order.GetOrderDetailsRequest) (*order.GetOrderDetailsResponse, error) {
	for _, ord := range s.orders {
		if ord.Id == req.GetOrderId() {
			return &order.GetOrderDetailsResponse{
				OrderId:      ord.Id,
				RestaurantId: ord.RestaurantId,
				FoodItem:     ord.FoodItem,
				Quantity:     int32(ord.Quantity),
				Status:       "success",
			}, nil
		}
	}
	return nil, fmt.Errorf("order not found")
}
