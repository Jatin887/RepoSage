package models
import(
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Admin bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push bool `json:"push"`
	Triage bool `json:"triage"`
	Pull bool `json:"pull"`
}