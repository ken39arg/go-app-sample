// Code generated by gencontroller; DO NOT EDIT

package response

import (
	"encoding/json"
	"net/http"

	"github.com/ken39arg/go-app-sample/app/data"
	"github.com/ken39arg/go-app-sample/app/errors"
)

type IsubataMessagePostResponse struct {

	// MessageID
	MessageID int64 `json:"message_id"`
}

func (res *IsubataMessagePostResponse) WriteTo(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}

func (res *IsubataMessagePostResponse) Validate() error {

	if res.MessageID < 1 {
		return errors.NewError(errors.InvalidParam, "MessageID must be greater than 1", res.MessageID)
	}

	return nil
}

type IsubataMessagesResponse struct {

	// メッセージ一覧
	Messages []data.IsubataMessageData `json:"messages"`
}

func (res *IsubataMessagesResponse) WriteTo(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}

func (res *IsubataMessagesResponse) Validate() error {

	if res.Messages == nil {
		return errors.NewError(errors.InvalidParam, "Messages is required", res.Messages)
	}
	for _, a := range res.Messages {
		if err := (&a).Validate(); err != nil {
			return errors.NewError(errors.InvalidParam, "Messages is invalid", err)
		}
	}
	return nil
}
