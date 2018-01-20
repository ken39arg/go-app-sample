package api

import (
	"net/http"

	"github.com/ken39arg/go-app-sample/app/model"
)

type IsubataController struct{}

func (c *IsubataController) MessagePostHandler(w http.ResponseWriter, r *http.Request) {
	req, chst, err := NewIsubataMessagePostRequest(r)
	if err != nil {
		NewErrorResponse(r.Context(), err).WriteTo(w)
		return
	}

	res := NewIsubataMessagePostResponse()
	res.MessageID, err = model.NewIsubata(chst.M).MessagePost(chst.Player, req.ChannelID, req.Message)
	if err != nil {
		NewErrorResponse(r.Context(), err).WriteTo(w)
		return
	}
	res.WriteTo(w)
}

func (c *IsubataController) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	req, chst, err := NewIsubataMessagesRequest(r)
	if err != nil {
		NewErrorResponse(r.Context(), err).WriteTo(w)
		return
	}

	res := NewIsubataMessagesResponse()
	res.Messages, err = model.NewIsubata(chst.M).Messages(chst.Player, req.ChannelID, req.LastMessageID)
	if err != nil {
		NewErrorResponse(r.Context(), err).WriteTo(w)
		return
	}
	res.WriteTo(w)
}

/*
	server := http.NewServeMux()
	isubataHandler := NewIsubataServer(&IsubataController{})
	server.Handle("/api/isubata", http.StripPrefix("/api", isubataHandler))
*/
