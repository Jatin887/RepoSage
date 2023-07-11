package initializers
import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(env *Env) {
	var err error

	dsn := env.DBURL
	DB ,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal("Failed to Connect to Database")
	}
}