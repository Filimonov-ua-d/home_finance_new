package unmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	_ "github.com/lib/pq"
)

type GetMoneyRequest struct {
	Date time.Time `json:"date" db:"date" tformat:"02.01.2006"`
	//Sum  string `json:"sum" db:"sum"`
}

func (d *GetMoneyRequest) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	rawDate := v.(map[string]interface{})
	vv := rawDate["date"].(string)

	c := reflect.TypeOf(*d).Field(0).Tag
	g := c.Get("tformat")

	dd, err := time.ParseInLocation(g, vv, time.Local)
	fmt.Println("Unmarshal rsult: ", dd, err)
	d.Date = dd
	return err
}
