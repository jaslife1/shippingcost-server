package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jaslife1/shippingcost-server/graph/generated"
	"github.com/jaslife1/shippingcost-server/graph/model"
)

func (r *queryResolver) CalculateShippingCost(ctx context.Context, senderAddress model.Address, receiverAddress model.Address) (int, error) {
	fmt.Println("Sender address: ", senderAddress.Town)
	fmt.Println("Receiver address: ", receiverAddress.Town)

	return 0, nil
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
func (r *queryResolver) Test(ctx context.Context) (string, error) {
	fmt.Println("Test is called")
	t := "Hello World from the other side"
	return t, nil
}

type mutationResolver struct{ *Resolver }
