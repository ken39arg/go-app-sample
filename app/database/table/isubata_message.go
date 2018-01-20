package table

import (
	"fmt"

	"github.com/ken39arg/go-app-sample/app/database/row"
)

// SelectByChannelIDWithLastID
func (t *IsubataMessage) SelectByChannelIDWithLastID(channelID int64, lastID *int64) ([]*row.IsubataMessage, error) {
	var last int64
	if lastID != nil {
		last = *lastID
	}
	// query generateは where inを含むなど好みに応じて github.com/Masterminds/squirrel を使っている場合もあります
	query := fmt.Sprintf("SELECT %s FROM %s WHERE channel_id = ? AND id > ? ORDER BY id ASC", t.ColumnsString(), t.Name())
	rows, err := t.c.DB().Query(query, channel_id, last)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return t.ScanRows(rows)
}
