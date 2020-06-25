package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/seanrmurphy/go-fullstack/frontend/src/actions"
	"github.com/seanrmurphy/go-fullstack/frontend/src/dispatcher"
	"github.com/seanrmurphy/go-fullstack/frontend/src/store/model"
	"github.com/seanrmurphy/go-fullstack/frontend/src/store/storeutil"
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

func parseResponse(resp []byte) {

	_ = json.Unmarshal(resp, &Items)

}

func Initialize(e string) {

	restEndpoint = e
	endpoint := restEndpoint + "todo"

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// handle the response

	if err != nil {
		log.Printf("Error talking to rest endpoint\n")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response...\n")
	}
	log.Printf("Response = %v\n", string(body))

	parseResponse(body)

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
	for _, item := range Items {
		if item.Completed == completed {
			count++
		}
	}
	return count
}

func addItem(i model.Item) {

	i.CreationDate = time.Now()
	i.ID = uuid.New()
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
		addItem(model.Item{Title: a.Title, Completed: false})
		//Items = append(Items, &model.Item{Title: a.Title, Completed: false})

	case *actions.DestroyItem:
		destroyItem(a.Index)

	case *actions.SetTitle:
		Items[a.Index].Title = a.Title
		go updateItem(Items[a.Index])

	case *actions.SetCompleted:
		Items[a.Index].Completed = a.Completed
		go updateItem(Items[a.Index])

	case *actions.SetAllCompleted:
		for _, item := range Items {
			item.Completed = a.Completed
		}

	case *actions.ClearCompleted:
		var activeItems []*model.Item
		for _, item := range Items {
			if !item.Completed {
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
