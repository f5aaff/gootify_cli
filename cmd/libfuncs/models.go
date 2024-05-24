package libfuncs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	baseURL = "http://localhost:3000/"
)

type Resource interface {
	GetType() string
	GetContent() interface{}
}

type BasicResource struct {
	Typ     string
	Content interface{}
}

func (br BasicResource) GetType() string {
	return br.Typ
}

func (br BasicResource) GetContent() interface{} {
	return br.Content
}

type Endpoint interface {
	GetEndPointStr() string
	Getitem() Resource
	GetAll() []Resource
}

type BasicEndPoint struct {
	EndpointStr string
}

func (bep *BasicEndPoint) GetEndPointStr() string {
	return bep.EndpointStr
}

func GetItem(resource Resource, ep string) {
	res, err := http.Get(baseURL + ep)
	if err != nil {
		fmt.Println("error getting " + ep + err.Error())
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading " + ep + "response" + err.Error())
		return
	}
	content := resource.GetContent()
	typ := resource.GetType()
	err = json.Unmarshal(body, content)
	resource = BasicResource{typ, content}
}
