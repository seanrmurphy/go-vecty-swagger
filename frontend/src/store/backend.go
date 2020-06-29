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
	i int
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
		Completed: i.BackEndModel.Completed,
		ID:        i.BackEndModel.ID,
		Title:     i.BackEndModel.Title,
		//CreationDate: time.Now(),
	}
	params := developers.NewUpdateTodoParams()
	params.Todo = &t
	params.Todoid = i.BackEndModel.ID.String()
	ctx := context.TODO()

	// should check return value here...
	_, _ = c.Developers.UpdateTodo(ctx, params)

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

	// should perform check on responses here...
	_, _ = c.Developers.AddTodo(ctx, params)

}

func destroyItemOnBackend(i *model.Item) {
	c := createClient()

	params := developers.NewDeleteTodoParams()
	params.Todoid = i.BackEndModel.ID.String()

	ctx := context.TODO()

	// should check the return value here...
	_, _ = c.Developers.DeleteTodo(ctx, params)

}
