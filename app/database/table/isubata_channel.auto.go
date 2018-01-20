// Code generated by genschema; DO NOT EDIT

package table

import (
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/ken39arg/go-app-sample/app/database"
	"github.com/ken39arg/go-app-sample/app/database/row"
	"github.com/ken39arg/go-app-sample/app/enum"
	mtime "github.com/ken39arg/go-app-sample/app/time"
)

type IsubataChannel struct {
	c *database.DBClient
}

func NewIsubataChannel(c *database.DBClient) *IsubataChannel {
	return &IsubataChannel{c}
}

func (t *IsubataChannel) Name() string {
	return "isubata_channel"
}

func (t *IsubataChannel) Columns() []string {
	return []string{"id", "name", "updated_at", "created_at"}
}

func (t *IsubataChannel) ColumnsString() string {
	return "id,name,updated_at,created_at"
}

func (t *IsubataChannel) ScanRow(r *sql.Row) (*row.IsubataChannel, error) {
	v := row.IsubataChannel{}
	if err := r.Scan(&v.ID, &v.Name, &v.UpdatedAt, &v.CreatedAt); err != nil {
		return nil, err
	}
	return &v, nil
}

func (t *IsubataChannel) ScanRows(r *sql.Rows) ([]*row.IsubataChannel, error) {
	ret := []*row.IsubataChannel{}
	for r.Next() {
		v := row.IsubataChannel{}
		if err := r.Scan(&v.ID, &v.Name, &v.UpdatedAt, &v.CreatedAt); err != nil {
			return nil, err
		}
		ret = append(ret, &v)
	}
	if err := r.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *IsubataChannel) Refetch(v *row.IsubataChannel) error {
	query := "SELECT id,name,updated_at,created_at FROM isubata_channel WHERE id = ? LIMIT 1"
	return t.c.DB().QueryRow(query, v.ID).Scan(
		&v.ID,
		&v.Name,
		&v.UpdatedAt,
		&v.CreatedAt,
	)
}

func (t *IsubataChannel) RefetchForUpdate(v *row.IsubataChannel) error {
	if !t.c.InTransaction() {
		return fmt.Errorf("RefetchForUpdate should be called in transaction")
	}

	err := t.c.AddEndHook(func() { v.IsForUpdate = false })
	if err != nil {
		return err
	}
	v.IsForUpdate = true
	query := "SELECT id,name,updated_at,created_at FROM isubata_channel WHERE id = ? LIMIT 1 FOR UPDATE"
	return t.c.DB().QueryRow(query, v.ID).Scan(
		&v.ID,
		&v.Name,
		&v.UpdatedAt,
		&v.CreatedAt,
	)
}

func (t *IsubataChannel) Insert(v *row.IsubataChannel) error {

	v.UpdatedAt = mtime.Now()
	v.CreatedAt = mtime.Now()

	_, err := t.c.DB().Exec("INSERT INTO isubata_channel (id,name,updated_at,created_at) VALUES (?,?,?,?)", v.ID, v.Name, v.UpdatedAt, v.CreatedAt)
	return err
}

func (t *IsubataChannel) Update(v *row.IsubataChannel) error {

	v.UpdatedAt = mtime.Now()

	query := "UPDATE isubata_channel SET name = ?,updated_at = ? WHERE id = ?"
	_, err := t.c.DB().Exec(query, v.Name, v.UpdatedAt.Format(mtime.MySQLDatetime), v.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *IsubataChannel) eachRowsFast(builder sq.SelectBuilder, fn func([]*row.IsubataChannel) error) error {
	builder = builder.OrderBy("id ASC").Limit(1000)

	itr := func(id int64) (int64, error) {
		query, args, err := builder.Where(sq.Gt{"id": id}).ToSql()
		if err != nil {
			return 0, err
		}
		rows, err := t.c.DB().Query(query, args...)
		if err != nil {
			return 0, err
		}
		defer rows.Close()

		rs, err := t.ScanRows(rows)
		if err != nil {
			return 0, err
		}

		var lastID int64
		if l := len(rs); l > 0 {
			lastID = rs[l-1].ID
		} else {
			return 0, nil
		}
		return lastID, fn(rs)
	}

	var lastID int64
	var err error
	for {
		lastID, err = itr(lastID)
		if err != nil {
			return err
		} else if lastID == 0 {
			break
		}
	}

	return nil
}

func (t *IsubataChannel) selectWithCursorID(builder sq.SelectBuilder, id int64, col string, order enum.PagerOrderType, perPage uint64) (rs []*row.IsubataChannel, prev, next int64, err error) {
	if order == enum.PagerNewer {
		if id != 0 {
			builder = builder.Where(sq.Gt{col: id})
		}
		builder = builder.OrderBy(col + " ASC")
	} else {
		if id != 0 {
			builder = builder.Where(sq.Lt{col: id})
		}
		builder = builder.OrderBy(col + " DESC")
	}

	query, args, err := builder.Limit(perPage + 1).ToSql()
	if err != nil {
		return
	}

	rows, err := t.c.DB().Query(query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	rs, err = t.ScanRows(rows)
	if err != nil {
		return
	}

	if len(rs) == 0 {
		return
	}

	if id != 0 {
		prev = rs[0].ID
	}
	if l := len(rs); uint64(l) > perPage {
		next = rs[l-2].ID
		rs = rs[:l-1]
	}

	if order == enum.PagerNewer {
		prev, next = next, prev
		for l, r := 0, len(rs)-1; l < r; l, r = l+1, r-1 {
			rs[l], rs[r] = rs[r], rs[l]
		}
	}

	return
}

func (t *IsubataChannel) FindByName(name string) (*row.IsubataChannel, error) {
	query := "SELECT id,name,updated_at,created_at FROM isubata_channel WHERE name = ? LIMIT 1"
	return t.ScanRow(t.c.DB().QueryRow(query, name))
}

func (t *IsubataChannel) FindAllByNames(names []string) ([]*row.IsubataChannel, error) {
	l := len(names)
	if 0 == l {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT id,name,updated_at,created_at FROM isubata_channel WHERE name IN (?%s)", strings.Repeat(",?", l-1))
	args := make([]interface{}, l)
	for i, v := range names {
		args[i] = v
	}
	rows, err := t.c.DB().Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return t.ScanRows(rows)
}

func (t *IsubataChannel) FindByID(id int64) (*row.IsubataChannel, error) {
	query := "SELECT id,name,updated_at,created_at FROM isubata_channel WHERE id = ? LIMIT 1"
	return t.ScanRow(t.c.DB().QueryRow(query, id))
}

func (t *IsubataChannel) FindAllByIDs(ids []int64) ([]*row.IsubataChannel, error) {
	l := len(ids)
	if 0 == l {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT id,name,updated_at,created_at FROM isubata_channel WHERE id IN (?%s)", strings.Repeat(",?", l-1))
	args := make([]interface{}, l)
	for i, v := range ids {
		args[i] = v
	}
	rows, err := t.c.DB().Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return t.ScanRows(rows)
}