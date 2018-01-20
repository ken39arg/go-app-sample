package row

import "github.com/ken39arg/go-app-sample/app/data"

func (r *IsubataUser) IsubataUserData() data.IsubataUserData {
	return data.IsubataUserData{
		Name:        r.Name,
		DisplayName: r.DisplayName,
		AvatarIcon:  r.AvatarIcon,
	}
}
