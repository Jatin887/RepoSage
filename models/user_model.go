package models
import(
	"gorm.io/gorm"
)
type User struct{
	gorm.Model
	ID int
	Name string
	Username string
	Email string
	Repos    []Repo `gorm:"many2many:user_repos;"`
} 