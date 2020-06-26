package store

import (
	"context"
	"net/http"
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
	r.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(r)
	return resp, err
}

func updateItem(i *model.Item) {

	rt := BrowserCompatibleRoundTripper{}
	conf := client.Config{
		Transport: rt,
	}

	c := client.New(conf)
	//c.transport =
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
	_, _ = c.Developers.UpdateTodo(ctx, params)

	//endpoint := restEndpoint + "todo/" + i.ID.String()
	//toPut, _ := json.Marshal(i)
	//req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(toPut))
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//req.Header.Add("js.fetch:mode", "cors")
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//fmt.Println(err)
	//log.Printf("Error PUTting item to backend\n")
	//return
	//}

	//defer resp.Body.Close()
}

func postItemToBackend(i model.Item) {
	rt := BrowserCompatibleRoundTripper{}
	conf := client.Config{
		Transport: rt,
	}
	c := client.New(conf)
	t := swaggermodel.Todo{
		Completed:    i.BackEndModel.Completed,
		ID:           i.BackEndModel.ID,
		Title:        i.BackEndModel.Title,
		CreationDate: strfmt.DateTime(time.Now()),
	}
	params := developers.NewAddTodoParams()
	params.Todo = &t
	ctx := context.TODO()
	_, _ = c.Developers.AddTodo(ctx, params)

	//endpoint := restEndpoint + "todo"
	//toPost, _ := json.Marshal(i)
	//req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(toPost))
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//req.Header.Add("js.fetch:mode", "cors")
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//defer resp.Body.Close()
	//// handle the response

	//if err != nil {
	//log.Printf("Error POSTing new item to backend\n")
	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//log.Printf("Error parsing response from backend...\n")
	//}
	//log.Printf("Response = %v\n", string(body))

}

func destroyItemOnBackend(i *model.Item) {

	rt := BrowserCompatibleRoundTripper{}
	conf := client.Config{
		Transport: rt,
	}
	c := client.New(conf)
	//t := swaggermodel.Todo{}
	params := developers.NewDeleteTodoParams()
	params.Todoid = i.BackEndModel.ID.String()

	ctx := context.TODO()
	_, _ = c.Developers.DeleteTodo(ctx, params)

	//resp := c.Developers.AddTodo(params)
	//endpoint := restEndpoint + "todo/" + i.ID.String()
	//req, err := http.NewRequest("DELETE", endpoint, nil)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//req.Header.Add("js.fetch:mode", "cors")
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//fmt.Println(err)
	//log.Printf("Error POSTing new item to backend\n")
	//return
	//}
	//defer resp.Body.Close()
}
