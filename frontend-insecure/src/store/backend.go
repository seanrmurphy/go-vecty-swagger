package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/seanrmurphy/go-fullstack/frontend/src/store/model"

	"github.com/seanrmurphy/go-vecty-swagger/client"
	"github.com/seanrmurphy/go-vecty-swagger/client/developers"
	swaggermodel "github.com/seanrmurphy/go-vecty-swagger/model"
)

func updateItem(i *model.Item) {

	c := client.New()
	t := swaggermodel.Todo{}
	params := developers.UpdateTodoParams{
		Todo:   t,
		TodoID: id,
	}
	resp := c.Developers.UpdateTodo(params)

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
	defer resp.Body.Close()
}

func postItemToBackend(i model.Item) {
	endpoint := restEndpoint + "todo"
	toPost, _ := json.Marshal(i)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(toPost))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// handle the response

	if err != nil {
		log.Printf("Error POSTing new item to backend\n")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error parsing response from backend...\n")
	}
	log.Printf("Response = %v\n", string(body))

}

func destroyItemOnBackend(i *model.Item) {
	endpoint := restEndpoint + "todo/" + i.ID.String()
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("js.fetch:mode", "cors")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		log.Printf("Error POSTing new item to backend\n")
		return
	}
	defer resp.Body.Close()
}
