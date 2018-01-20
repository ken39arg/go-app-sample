package row

import "github.com/ken39arg/go-app-sample/app/data"

func (r *IsubataMessage) IsubataMessageData() data.IsubataMessageData {
	return data.IsubataMessageData{
		Identifiable: data.Identifiable{ID: r.ID},
		User:         r.IsubataUser.IsubataUserData(),
		Content:      r.Content,
		Date:         data.UnixTime(r.CreatedAt),
	}
}
