//package cache
//
//import (
//	"encoding/json"
//	"log"
//)
//
//type jsonStruct struct {
//	Uuid  string `json:"uuid"`
//	Date  string `json:"date"`
//	Event string `json:"event"`
//}
//type Cache struct {
//	m map[string]jsonStruct
//}
//
//func (c Cache) Packing(ma interface{}) {
//
//}
//
//func Serialization(j jsonStruct) (string, error) {
//	res, err := json.Marshal(j)
//	if err != nil {
//		log.Println("Error with marshaling:\n", err)
//		return "", err
//	}
//	log.Println(res)
//	return string(res), nil
//}
