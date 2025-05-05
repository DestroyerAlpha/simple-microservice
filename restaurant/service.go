package restaurant

import (
	"context"
	"fmt"

	"github.com/DestroyerAlpha/simple-microservice/api/restaurant"

	"github.com/google/uuid"
)

type Service struct {
	restaurant.UnimplementedRestaurantServiceServer

	restaurantList []string
	restaurants    map[string]*Restaurant
}

func NewRestaurantService() *Service {
	return &Service{
		restaurantList: []string{},
		restaurants:    make(map[string]*Restaurant),
	}
}

type Restaurant struct {
	Id           string
	Name         string
	Cuisine      string
	Ratings      map[string]float32
	TotalRatings map[string]int
	FoodItems    []string
}

func (s *Service) ListRestaurants(ctx context.Context, request *restaurant.ListRestaurantsRequest) (*restaurant.ListRestaurantsResponse, error) {
	var restaurantList []string
	for _, id := range s.restaurantList {
		rest := s.restaurants[id]
		if rest != nil && rest.Cuisine == request.GetCuisine() {
			restaurantList = append(restaurantList, id)
		}
	}
	return &restaurant.ListRestaurantsResponse{
		RestaurantIds: restaurantList,
	}, nil
}

func (s *Service) GetRestaurantDetails(ctx context.Context, request *restaurant.GetRestaurantDetailsRequest) (*restaurant.GetRestaurantDetailsResponse, error) {
	rest, ok := s.restaurants[request.GetRestaurantId()]
	if !ok {
		return nil, fmt.Errorf("restaurant not found")
	}
	return &restaurant.GetRestaurantDetailsResponse{
		RestaurantId: rest.Id,
		Name:         rest.Name,
		Cuisine:      rest.Cuisine,
		FoodItems:    rest.FoodItems,
	}, nil
}

func (s *Service) AddRestaurant(ctx context.Context, request *restaurant.AddRestaurantRequest) (*restaurant.AddRestaurantResponse, error) {
	id := uuid.NewString()
	ratings := make(map[string]float32)
	totalRatings := make(map[string]int)
	for _, foodItem := range request.GetFoodItems() {
		ratings[foodItem] = 0
		totalRatings[foodItem] = 0
	}
	rest := &Restaurant{
		Id:           id,
		Name:         request.GetName(),
		Cuisine:      request.GetCuisine(),
		Ratings:      ratings,
		TotalRatings: totalRatings,
		FoodItems:    request.GetFoodItems(),
	}
	s.restaurants[id] = rest
	s.restaurantList = append(s.restaurantList, id)
	return &restaurant.AddRestaurantResponse{
		RestaurantId: id,
		Status:       "success",
	}, nil
}

func (s *Service) AddRating(ctx context.Context, request *restaurant.AddRatingRequest) (*restaurant.AddRatingResponse, error) {
	id := request.GetRestaurantId()
	rest, ok := s.restaurants[id]
	if !ok {
		return nil, fmt.Errorf("restaurant not found")
	}
	oldRating := rest.Ratings[request.GetFoodItem()] * float32(rest.TotalRatings[request.GetFoodItem()])
	rest.TotalRatings[request.GetFoodItem()]++
	newRating := (oldRating + float32(request.GetRating())) / float32(rest.TotalRatings[request.GetFoodItem()])
	rest.Ratings[request.GetFoodItem()] = newRating
	return &restaurant.AddRatingResponse{
		Status: "success",
	}, nil
}
