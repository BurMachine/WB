package storage

import (
	"encoding/json"
)

type Storage struct {
	data map[string]*UserGrade
}

type UserGrade struct {
	UserId        string `json:"user_id" validate:"required"`
	PostpaidLimit int    `json:"postpaid_limit"`
	Spp           int    `json:"spp"`
	ShippingFee   int    `json:"shipping_fee"`
	ReturnFee     int    `json:"return_fee"`
}

func NewStorage() *Storage {
	data := make(map[string]*UserGrade)
	return &Storage{data: data}
}

func (s *Storage) Set(data []byte) error {

	usGr := new(UserGrade)
	err := json.Unmarshal(data, &usGr)
	if err != nil {
		return err
	}
	tmpUsGr, ok := s.data[usGr.UserId]
	if !ok {
		s.data[usGr.UserId] = usGr
	} else {
		if usGr.Spp != 0 {
			tmpUsGr.Spp = usGr.Spp
		}
		if usGr.ReturnFee != 0 {
			tmpUsGr.ReturnFee = usGr.ReturnFee
		}
		if usGr.PostpaidLimit != 0 {
			tmpUsGr.PostpaidLimit = usGr.PostpaidLimit
		}
		if usGr.ShippingFee != 0 {
			tmpUsGr.ShippingFee = usGr.ShippingFee
		}
		s.data[usGr.UserId] = tmpUsGr
	}

	return nil
}
