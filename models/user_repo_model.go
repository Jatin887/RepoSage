package models
import(
	"gorm.io/gorm"
)
type UserRepo struct {
    gorm.Model
    UserID int
    RepoID int
}