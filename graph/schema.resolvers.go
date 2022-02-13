package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jaslife1/shippingcost-server/graph/generated"
	"github.com/jaslife1/shippingcost-server/graph/model"
	"github.com/jaslife1/shippingcost-server/utils"
)

func (r *queryResolver) AllProvinces(ctx context.Context) ([]*string, error) {
	return utils.GetAllProvinces()
}

func (r *queryResolver) AllCitiesOfProvince(ctx context.Context, province string) ([]*string, error) {
	return utils.GetAllCities(province)
}

func (r *queryResolver) CalculateShippingCost(ctx context.Context, senderAddress model.Address, receiverAddress model.Address) (int, error) {
	fmt.Println("Sender address: ", senderAddress.City)
	fmt.Println("Receiver address: ", receiverAddress.City)

	weight := 1.0

	val := utils.CalculateJnTShippingCost(*senderAddress.City, *receiverAddress.City, weight)

	return int(val), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Towns(ctx context.Context) ([]*string, error) {

	var ret = utils.GetAllTowns()
	fmt.Printf("Getting all towns: %+v\n", ret)
	return ret, nil
}
func (r *queryResolver) Test(ctx context.Context) (string, error) {
	fmt.Println("Test is called")
	t := "Hello World from the other side"
	return t, nil
}

type mutationResolver struct{ *Resolver }
