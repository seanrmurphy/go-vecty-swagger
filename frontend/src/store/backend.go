package store

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/store/model"

	"github.com/seanrmurphy/go-vecty-swagger/client"
	"github.com/seanrmurphy/go-vecty-swagger/client/developers"
	swaggermodel "github.com/seanrmurphy/go-vecty-swagger/models"
)

type BrowserCompatibleRoundTripper struct {
}

func (rt BrowserCompatibleRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Printf("In round tripper...\n")
	r.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(r)
	return resp, err
}

func createClient() *client.SimpleTodo {
	rt := BrowserCompatibleRoundTripper{}
	url, _ := url.Parse(restEndpoint)
	conf := client.Config{
		URL:       url,
		Transport: rt,
	}
	c := client.New(conf)
	return c
}

func updateItem(i *model.Item) {
	c := createClient()

	t := swaggermodel.Todo{
		Completed:    i.BackEndModel.Completed,
		ID:           i.BackEndModel.ID,
		Title:        i.BackEndModel.Title,
		CreationDate: strfmt.DateTime(time.Now()),
	}
	params := developers.NewUpdateTodoParams()
	params.Todo = &t
	params.Todoid = i.BackEndModel.ID.String()
	ctx := context.TODO()

	if _, err := c.Developers.UpdateTodo(ctx, params); err != nil {
		log.Printf("Error updating item on backend - error %v\n", err)
		return
	}
}

func postItemToBackend(i model.Item) {
	c := createClient()

	t := swaggermodel.Todo{
		Completed:    i.BackEndModel.Completed,
		ID:           i.BackEndModel.ID,
		Title:        i.BackEndModel.Title,
		CreationDate: strfmt.DateTime(time.Now()),
	}
	params := developers.NewAddTodoParams()
	params.Todo = &t
	ctx := context.TODO()

	if _, err := c.Developers.AddTodo(ctx, params); err != nil {
		log.Printf("Error pusting new item on backend - error %v\n", err)
		return
	}
}

func destroyItemOnBackend(i *model.Item) {
	c := createClient()

	params := developers.NewDeleteTodoParams()
	params.Todoid = i.BackEndModel.ID.String()

	ctx := context.TODO()

	if _, err := c.Developers.DeleteTodo(ctx, params); err != nil {
		log.Printf("Error deleting item on backend - error %v\n", err)
		return
	}
}
