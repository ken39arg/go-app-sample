package model

import (
	"database/sql"

	"github.com/ken39arg/go-app-sample/app/data"
	"github.com/ken39arg/go-app-sample/app/database/row"
	"github.com/ken39arg/go-app-sample/app/database/table"
	"github.com/ken39arg/go-app-sample/app/errors"
)

type Isubata struct {
	*Models
}

func NewIsubata(m *Models) *Isubata {
	return &Isubata{Models: m}
}

func (m *Isubata) MessagePost(player *row.Player, channelID int64, message string) (int64, error) {
	txn, err := m.DB().Master().Begin()
	if err != nil {
		return 0, errors.NewError(errors.SystemErr, "Transaction begin failed", err)
	}
	defer txn.End()

	channel, err := table.NewIsubataChannel(m.DB().Master()).FindByID(channelID)
	switch {
	case err == sql.ErrNoRows:
		return 0, errors.NewError(errors.NotFound, "channnel not found", channelID)
	case err != nil:
		return 0, errors.NewError(errors.SystemErr, "isubataChannel select failed", err)
	}

	id, err := m.Katsubushi().Get()
	if err != nil {
		return 0, errors.NewError(errors.SystemErr, "katsubushi get failed", err)
	}
	err = table.NewIsubataMessage(m.DB().Master()).Insert(&row.IsubataMessage{
		ID:        id,
		ChannelID: channel.ID,
		Content:   message,
	})
	if err != nil {
		return 0, errors.NewError(errors.SystemErr, "isubataMessage insert failed", err)
	}

	if err = txn.Commit(); err != nil {
		return 0, errors.NewError(errors.SystemErr, "Transaction commit failed", err)
	}
	return id, nil
}

func (m *Isubata) Messages(player *row.Player, channelID int64, lastMessageID *int64) ([]data.IsubataMessageData, error) {
	channel, err := table.NewIsubataChannel(m.DB().Slave()).FindByID(channelID)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.NewError(errors.NotFound, "channnel not found", channelID)
	case err != nil:
		return nil, errors.NewError(errors.SystemErr, "isubataChannel select failed", err)
	}

	messages, err := table.NewIsubataMessage(m.DB().Slave()).SelectByChannelIDWithLastID(channelID, lastMessageID)
	if err != nil {
		return nil, errors.NewError(errors.SystemErr, "isubataMessage select failed", err)
	}
	res := make([]data.IsubataMessageData, len(messages))
	for i, message := range messages {
		res[i] = message.ToIsubataMessageData()
	}
	return res, nil
}
