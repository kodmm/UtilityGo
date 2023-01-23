package main

import(
	"time"
	"fmt"
	"encoding/json"
)

type CreatedColumns struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedTraceID string `json:"created_trace_id"`
}

type UpdatedColumns struct {
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedTraceID string `json:"updated_trace_id"`
	Revision int64 `json:"revision"`
}

type Customer struct {
	CustomerID string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	CreatedColumns
	UpdatedColumns
}

func main() {

	c := Customer{
		CustomerID: "001",
		CustomerName: "o'reily taro",
		CreatedColumns: CreatedColumns{
			CreatedAt: time.Now(),
			CreatedTraceID: "55034320- 573429eiifrvvm-befewifw",
		},
		UpdatedColumns: UpdatedColumns{
			UpdatedAt: time.Now(),
			UpdatedTraceID: "55034320- 573429eiifrvvm-befewifw",
			Revision: 1,
		},
	}

	b, _ := json.Marshal(c)
	fmt.Printf("%+v", string(b))
}