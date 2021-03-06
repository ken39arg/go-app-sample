// Code generated by gencontroller; DO NOT EDIT

package data

import (
	"unicode/utf8"

	"github.com/ken39arg/go-app-sample/app/errors"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

type IsubataUserData struct {

	// 名前
	Name string `json:"name" schema:"name,required"`
	// 表示名
	DisplayName string `json:"display_name" schema:"display_name,required"`
	// アイコンURL
	AvatarIcon string `json:"avatar_icon" schema:"avatar_icon,required"`
}

func (d *IsubataUserData) Validate() error {

	lenName := utf8.RuneCountInString(d.Name)

	if lenName < 1 {
		return errors.NewError(errors.InvalidParam, "Name length must be greater than 1", lenName)
	}

	if lenName > 20 {
		return errors.NewError(errors.InvalidParam, "Name length must be smaller than 20", lenName)
	}

	lenDisplayName := utf8.RuneCountInString(d.DisplayName)

	if lenDisplayName < 1 {
		return errors.NewError(errors.InvalidParam, "DisplayName length must be greater than 1", lenDisplayName)
	}

	if lenDisplayName > 15 {
		return errors.NewError(errors.InvalidParam, "DisplayName length must be smaller than 15", lenDisplayName)
	}

	lenAvatarIcon := utf8.RuneCountInString(d.AvatarIcon)

	if lenAvatarIcon < 8 {
		return errors.NewError(errors.InvalidParam, "AvatarIcon length must be greater than 8", lenAvatarIcon)
	}

	if lenAvatarIcon > 100 {
		return errors.NewError(errors.InvalidParam, "AvatarIcon length must be smaller than 100", lenAvatarIcon)
	}

	return nil
}

func (d *IsubataUserData) DecodeMsgpack(dec *msgpack.Decoder) error {

	var key int
	len, err := dec.DecodeMapLen()
	if err != nil {
		return err
	}
	if len != 3 {
		return errors.NewError(errors.SystemErr, "size is not match", len)
	}
	for i := 0; i < len; i++ {
		if key, err = dec.DecodeInt(); err != nil {
			return err
		}
		switch key {
		case 0:
			err = dec.Decode(&d.Name)
		case 1:
			err = dec.Decode(&d.DisplayName)
		case 2:
			err = dec.Decode(&d.AvatarIcon)
		default:
			err = errors.NewError(errors.SystemErr, "unknown key.", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (d IsubataUserData) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeMapLen(3); err != nil {
		return err
	}
	if err := enc.Encode(int(0), d.Name); err != nil {
		return err
	}
	if err := enc.Encode(int(1), d.DisplayName); err != nil {
		return err
	}
	if err := enc.Encode(int(2), d.AvatarIcon); err != nil {
		return err
	}
	return nil
}

type IsubataMessageData struct {
	IdentifiableLong

	// ユーザー
	User IsubataUserData `json:"user" schema:"user,required"`
	// メッセージ本文
	Content string `json:"content" schema:"content,required"`
	// 送信時間
	Date UnixTime `json:"date" schema:"date,required"`
}

func (d *IsubataMessageData) Validate() error {

	if err := d.IdentifiableLong.Validate(); err != nil {
		return err
	}

	if err := (&d.User).Validate(); err != nil {
		return errors.NewError(errors.InvalidParam, "User is invalid", err)
	}

	lenContent := utf8.RuneCountInString(d.Content)

	if lenContent < 1 {
		return errors.NewError(errors.InvalidParam, "Content length must be greater than 1", lenContent)
	}

	if lenContent > 100 {
		return errors.NewError(errors.InvalidParam, "Content length must be smaller than 100", lenContent)
	}

	return nil
}

func (d *IsubataMessageData) DecodeMsgpack(dec *msgpack.Decoder) error {

	var key int
	len, err := dec.DecodeMapLen()
	if err != nil {
		return err
	}
	if len != 4 {
		return errors.NewError(errors.SystemErr, "size is not match", len)
	}
	for i := 0; i < len; i++ {
		if key, err = dec.DecodeInt(); err != nil {
			return err
		}
		switch key {
		case 0:
			err = dec.Decode(&d.ID)
		case 1:
			err = dec.Decode(&d.User)
		case 2:
			err = dec.Decode(&d.Content)
		case 3:
			err = dec.Decode(&d.Date)
		default:
			err = errors.NewError(errors.SystemErr, "unknown key.", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (d IsubataMessageData) EncodeMsgpack(enc *msgpack.Encoder) error {
	if err := enc.EncodeMapLen(4); err != nil {
		return err
	}
	if err := enc.Encode(int(0), d.ID); err != nil {
		return err
	}
	if err := enc.Encode(int(1), d.User); err != nil {
		return err
	}
	if err := enc.Encode(int(2), d.Content); err != nil {
		return err
	}
	if err := enc.Encode(int(3), d.Date); err != nil {
		return err
	}
	return nil
}
