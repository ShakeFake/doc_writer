package order_map

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/orderedmap"
	"testing"
)

type Sample struct {
	Name string `json:"name"`
}

func TestOrderMap(t *testing.T) {
	om := orderedmap.New()
	om.Set("abc", &Sample{
		Name: "abc",
	})

	om.Set("b", "abc")

	result, err := json.Marshal(om)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

}
