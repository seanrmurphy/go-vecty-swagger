package store

import (
	"context"
	"net/url"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/seanrmurphy/go-vecty-swagger/client"
	"github.com/seanrmurphy/go-vecty-swagger/client/developers"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/actions"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/dispatcher"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/store/model"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/store/storeutil"
	swaggermodel "github.com/seanrmurphy/go-vecty-swagger/models"
)

var (
	// Items represents all of the TODO items in the store.
	Items []*model.Item

	// Filter represents the active viewing filter for items.
	Filter = model.All

	// Listeners is the listeners that will be invoked when the store changes.
	Listeners = storeutil.NewListenerRegistry()

	restEndpoint string
)

func init() {
	dispatcher.Register(onAction)
}

func Initialize(s string) {

	//restEndpoint = e
	//endpoint := restEndpoint + "todo"
	restEndpoint = s

	rt := BrowserCompatibleRoundTripper{}
	url, _ := url.Parse(restEndpoint)
	//transport := http.transport.New(host, basePath, schemes)
	//transport.Transport = rt
	conf := client.Config{
		URL:       url,
		Transport: rt,
	}
	c := client.New(conf)

	p := developers.NewGetAllTodosParams()
	ctx := context.TODO()
	todos, _ := c.Developers.GetAllTodos(ctx, p)

	for _, t := range todos.Payload {
		i := model.Item{
			BackEndModel: *t,
		}
		Items = append(Items, &i)
		//log.Printf("todo = %v\n", t)
	}

	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//log.Printf("Error reading response...\n")
	//}
	//log.Printf("Response = %v\n", string(body))

	//parseResponse(body)

	dispatcher.Dispatch(&actions.ReplaceItems{
		Items: Items,
	})
}

// ActiveItemCount returns the current number of items that are not completed.
func ActiveItemCount() int {
	return count(false)
}

// CompletedItemCount returns the current number of items that are completed.
func CompletedItemCount() int {
	return count(true)
}

func count(completed bool) int {
	count := 0
	for _, i := range Items {
		if i.BackEndModel.Completed == completed {
			count++
		}
	}
	return count
}

func addItem(i model.Item) {

	i.BackEndModel.CreationDate = strfmt.DateTime(time.Now())
	i.BackEndModel.ID = strfmt.UUID(uuid.New().String())
	go postItemToBackend(i)
	Items = append(Items, &i)
}

func destroyItem(idx int) {
	go destroyItemOnBackend(Items[idx])
	copy(Items[idx:], Items[idx+1:])
	Items = Items[:len(Items)-1]
}

func onAction(action interface{}) {
	switch a := action.(type) {
	case *actions.ReplaceItems:
		Items = a.Items

	case *actions.AddItem:
		m := model.Item{
			BackEndModel: swaggermodel.Todo{
				Title:     &a.Title,
				Completed: false,
			},
		}
		addItem(m)
		//Items = append(Items, &model.Item{Title: a.Title, Completed: false})

	case *actions.DestroyItem:
		destroyItem(a.Index)

	case *actions.SetTitle:
		Items[a.Index].BackEndModel.Title = &a.Title
		go updateItem(Items[a.Index])

	case *actions.SetCompleted:
		Items[a.Index].BackEndModel.Completed = a.Completed
		go updateItem(Items[a.Index])

	case *actions.SetAllCompleted:
		for _, item := range Items {
			item.BackEndModel.Completed = a.Completed
		}

	case *actions.ClearCompleted:
		var activeItems []*model.Item
		for _, item := range Items {
			if !item.BackEndModel.Completed {
				activeItems = append(activeItems, item)
			}
		}
		Items = activeItems

	case *actions.SetFilter:
		Filter = a.Filter

	default:
		return // don't fire listeners
	}

	Listeners.Fire()
}
