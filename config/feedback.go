package config

import (
	"encoding/json"
	"fmt"

	"github.com/shekohex/anonback/utils"
)

// Feedback ...
type Feedback struct {
	ID      string   `json:"id"`
	Message *Message `json:"message"`
	Reply   *Message `json:"reply"`
}

// Message the message object
type Message struct {
	Value     string `form:"value" json:"value" binding:"required"`
	CreatedAt string `json:"createdAt"`
}

// New create new feedback
func (f *Feedback) New(message *Message) *Feedback {
	f.ID = utils.RandUID(16)
	f.Message = message
	return f
}

// MarshalBinary -
func (f *Feedback) MarshalBinary() []byte {
	val, err := json.Marshal(f)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// UnmarshalBinary -
func (f *Feedback) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &f)
	if err != nil {
		return err
	}

	return nil
}
