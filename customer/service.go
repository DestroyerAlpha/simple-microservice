package customer

import (
	"context"
	"fmt"

	"github.com/DestroyerAlpha/simple-microservice/api/customer"
	"github.com/DestroyerAlpha/simple-microservice/api/order"
	"github.com/DestroyerAlpha/simple-microservice/api/ratings"
)

type Service struct {
	customer.UnimplementedCustomerServiceServer
	orderClient   order.OrderServiceClient
	ratingsClient ratings.RatingsServiceClient
}

func NewCustomerService(orderClient order.OrderServiceClient, ratingsClient ratings.RatingsServiceClient) *Service {
	return &Service{orderClient: orderClient, ratingsClient: ratingsClient}
}

func (s *Service) GetMenu(ctx context.Context, req *customer.GetMenuRequest) (*customer.GetMenuResponse, error) {
	resp, err := s.orderClient.GetMenu(ctx, &order.GetMenuRequest{
		Cuisine: req.GetCuisine(),
	})
	if err != nil {
		return nil, fmt.Errorf("error in getting menu: %v", err)
	}
	return &customer.GetMenuResponse{
		FoodItems: resp.GetFoodItems(),
	}, nil
}

func (s *Service) PlaceFoodOrder(ctx context.Context, req *customer.PlaceFoodOrderRequest) (*customer.PlaceFoodOrderResponse, error) {
	resp, err := s.orderClient.PlaceOrder(ctx, &order.PlaceOrderRequest{
		FoodItem: req.GetFoodItem(),
		Quantity: req.GetQuantity(),
		Cuisine:  req.GetCuisine(),
	})
	if err != nil || resp.GetStatus() != "success" {
		return nil, fmt.Errorf("error in placing order: %v, status: %s", err, resp.GetStatus())
	}
	return &customer.PlaceFoodOrderResponse{
		OrderId: resp.GetOrderId(),
		Status:  "success",
	}, nil
}

func (s *Service) ReviewFoodItem(ctx context.Context, req *customer.ReviewFoodItemRequest) (*customer.ReviewFoodItemResponse, error) {
	resp, err := s.ratingsClient.SubmitRating(ctx, &ratings.SubmitRatingRequest{
		OrderId:  req.GetOrderId(),
		FoodItem: req.GetFoodItem(),
		Rating:   req.GetRating(),
	})
	if err != nil || resp.GetStatus() != "success" {
		return nil, fmt.Errorf("error in submitting ratings: %v, status: %s", err, resp.GetStatus())
	}
	return &customer.ReviewFoodItemResponse{
		Status: "success",
	}, nil
}
