package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Type      string `json:"type,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	// Payloadを具体的な構造体に展開せず、json.RawMessageとして保持
	Payload json.RawMessage `json:"payload,omitempty"`
}

// message のペイロードを扱う構造体
type Message struct {
	ID        string  `json:"id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Message   string  `json:"message,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

// sensor のペイロードを扱う構造体
type Sensor struct {
	ID        string `json:"id,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	Result    string `json:"result,omitempty"`
	ProductID string `json:"product_id,omitempty"`
}

func main() {
	resMessage := `{
		"type": "message",
		"time": 1639269962,
		"payload": {
			"id": "6fc241563784abd",
			"user_id": "ABC12345",
			"message": "もうすぐ到着",
			"latitude": 86.242224,
			"longitude": 131.5353553
		}
	}`

	// 一度の json.RawMessageのフィールドとしてデコードし
	// Type の値によってもう一度デコードする
	var r Response
	err := json.Unmarshal([]byte(resMessage), &r)
	if err != nil {
		log.Fatal(err)
	}

	switch r.Type {
	case "message":
		var m Message
		_ = json.Unmarshal(r.Payload, &m)
		fmt.Println(m)
	case "sensor":
		var s Sensor
		_ = json.Unmarshal(r.Payload, &s)
		fmt.Println(s)
	default:
		fmt.Println(r)
		// ...
	}

}
