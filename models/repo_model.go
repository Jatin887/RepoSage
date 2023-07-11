package models
import(
	"gorm.io/gorm"
)

type AccessLevel string

const (
    PublicAccess  AccessLevel = "public"
    PrivateAccess AccessLevel = "private"
)

type Repo struct{
	gorm.Model
	OwnerID int
	OwnerName string
	OwnerEmail string
	RepoID int
	RepoName string
	Status AccessLevel
	Stars int
	Users       []User `gorm:"many2many:user_repos;"`
} 