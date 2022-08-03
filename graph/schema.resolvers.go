package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sort"
	"github.com/joheee/johedotcom/graph/generated"
	"github.com/joheee/johedotcom/graph/model"
)

// CreateExperience is the resolver for the createExperience field.
func (r *mutationResolver) CreateExperience(ctx context.Context, input model.NewExperience) (*model.Experience, error) {
	dummyExperience := &model.Experience{
		ID: input.InputID,
		Title: input.InputTitle,
		YearBegin: input.InputYearBegin,
		YearEnd: input.InputYearEnd,
		Description: input.InputDescription,
		PointList: input.InputPointList,
	}
	r.experiences = append(r.experiences, dummyExperience)
	return dummyExperience, nil
}

// UpdateExperience is the resolver for the updateExperience field.
func (r *mutationResolver) UpdateExperience(ctx context.Context, input model.NewExperience) (*model.Experience, error) {
	var tempItem *model.Experience
	for i, item := range r.experiences {
		if item.ID == input.InputID{
			item = &model.Experience{
				ID: input.InputID,
				Title: input.InputTitle,
				YearBegin: input.InputYearBegin,
				YearEnd: input.InputYearEnd,
				Description: input.InputDescription,
				PointList: input.InputPointList,
			}
			tempItem = item
			r.experiences[i] = r.experiences[len(r.experiences) - 1]
			r.experiences[len(r.experiences) - 1] = &model.Experience{}
			r.experiences = r.experiences[:len(r.experiences) - 1]
			r.experiences = append(r.experiences, tempItem)
		}
	}
	return tempItem, nil
}

// DeleteExperience is the resolver for the deleteExperience field.
func (r *mutationResolver) DeleteExperience(ctx context.Context, input model.InputDeleteExperience) (*model.Experience, error) {
	var tempItem *model.Experience
	for i, item := range r.experiences{
		if item.ID == input.InputID {
			tempItem = item
			r.experiences[i] = r.experiences[len(r.experiences) - 1]
			r.experiences[len(r.experiences) - 1] = &model.Experience{}
			r.experiences = r.experiences[:len(r.experiences) - 1]
		}
	}
	return tempItem, nil
}

// Experiences is the resolver for the experiences field.
func (r *queryResolver) Experiences(ctx context.Context) ([]*model.Experience, error) {
	sort.Slice(r.experiences, func(i int, j int) bool {return r.experiences[i].ID < r.experiences[j].ID})
	return r.experiences, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
