// Code generated by gencontroller; DO NOT EDIT

package request

import (
	"unicode/utf8"

	"github.com/ken39arg/go-app-sample/app/errors"
)

type IsubataMessagePostRequest struct {

	// 作成したRoomID
	ChannelID int64 `json:"channel_id" schema:"channel_id,required"`
	// 送信するメッセージ
	Message string `json:"message" schema:"message,required"`
}

func (req *IsubataMessagePostRequest) Validate() error {

	if req.ChannelID < 1 {
		return errors.NewError(errors.InvalidParam, "ChannelID must be greater than 1", req.ChannelID)
	}

	lenMessage := utf8.RuneCountInString(req.Message)

	if lenMessage < 1 {
		return errors.NewError(errors.InvalidParam, "Message length must be greater than 1", lenMessage)
	}

	if lenMessage > 100 {
		return errors.NewError(errors.InvalidParam, "Message length must be smaller than 100", lenMessage)
	}

	return nil
}

type IsubataMessagesRequest struct {

	// 作成したRoomID
	ChannelID int64 `json:"channel_id" schema:"channel_id,required"`
	// 取得済みのMessageID
	LastMessageID *int64 `json:"last_message_id,omitempty" schema:"last_message_id"`
}

func (req *IsubataMessagesRequest) Validate() error {

	if req.ChannelID < 1 {
		return errors.NewError(errors.InvalidParam, "ChannelID must be greater than 1", req.ChannelID)
	}

	if req.LastMessageID != nil {

		if *req.LastMessageID < 0 {
			return errors.NewError(errors.InvalidParam, "LastMessageID must be greater than 0", *req.LastMessageID)
		}

	}

	return nil
}
