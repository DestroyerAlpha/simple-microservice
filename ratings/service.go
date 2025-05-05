package ratings

import (
	"context"
	"fmt"

	"github.com/DestroyerAlpha/simple-microservice/api/order"
	"github.com/DestroyerAlpha/simple-microservice/api/ratings"
	"github.com/DestroyerAlpha/simple-microservice/api/restaurant"
)

type Service struct {
	ratings.UnimplementedRatingsServiceServer
	restaurantClient restaurant.RestaurantServiceClient
	orderClient      order.OrderServiceClient
}

func NewRatingsService(restaurantClient restaurant.RestaurantServiceClient, orderClient order.OrderServiceClient) *Service {
	return &Service{restaurantClient: restaurantClient, orderClient: orderClient}
}

func (s *Service) SubmitRating(ctx context.Context, req *ratings.SubmitRatingRequest) (*ratings.SubmitRatingResponse, error) {
	orderResp, err := s.orderClient.GetOrderDetails(ctx, &order.GetOrderDetailsRequest{
		OrderId: req.GetOrderId(),
	})
	if err != nil || orderResp.GetStatus() != "success" {
		return nil, fmt.Errorf("error in getting order details: %v, status: %s", err, orderResp.GetStatus())
	}
	restResp, err := s.restaurantClient.AddRating(ctx, &restaurant.AddRatingRequest{
		RestaurantId: orderResp.GetRestaurantId(),
		FoodItem:     req.GetFoodItem(),
		Rating:       req.GetRating(),
	})
	if err != nil || restResp.GetStatus() != "success" {
		return nil, fmt.Errorf("error in submitting ratings: %v, status: %s", err, restResp.GetStatus())
	}
	return &ratings.SubmitRatingResponse{
		Status: "success",
	}, nil
}
