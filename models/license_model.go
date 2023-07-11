package models
import(
	"gorm.io/gorm"
)
type License struct {
	gorm.Model
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}