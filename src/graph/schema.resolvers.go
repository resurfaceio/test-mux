package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strconv"

	"github.com/resurfaceio/test-mux/graph/generated"
	"github.com/resurfaceio/test-mux/graph/model"
	"github.com/resurfaceio/test-mux/internal/news"
	database "github.com/resurfaceio/test-mux/internal/pkg/db"
)

func (r *mutationResolver) AddNews(ctx context.Context, title string, body string) (*model.AllNews, error) {
	var news news.News
	news.Title = title
	news.Body = body
	newsID := news.Save()

	return &model.AllNews{News: &model.News{ID: strconv.FormatInt(newsID, 10), Title: news.Title, Body: news.Body}}, nil
}

func (r *mutationResolver) DeleteEverything(ctx context.Context) (*model.Ok, error) {
	database.Truncate()
	log.Print("Truncated database")

	return &model.Ok{Ok: true}, nil
}

func (r *queryResolver) AllNews(ctx context.Context) ([]*model.News, error) {
	var resultNews []*model.News
	dbNews := news.GetAll()
	for _, nnews := range dbNews {
		resultNews = append(resultNews, &model.News{ID: nnews.ID, Title: nnews.Title, Body: nnews.Body})
	}
	return resultNews, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
