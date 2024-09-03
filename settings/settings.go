package settings

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Struct that contains settings
type Settings struct {
	TgSenderToken   string
	TgReceiverToken string
	PgIp            string
	PgPort          int
	PgDb            string
	PgPass          string
	PgUser          string
}

var instSettings *Settings
var once sync.Once

// GetSettings is a func to get settings instance
func GetSettings() *Settings {

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load .env file !")
		}
		instSettings = &Settings{}
		instSettings.TgReceiverToken = os.Getenv("TG_RECEIVER_TOKEN")
		instSettings.TgSenderToken = os.Getenv("TG_SENDER_TOKEN")
		instSettings.PgIp = os.Getenv("POSTGRES_HOST")
		instSettings.PgPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		instSettings.PgDb = os.Getenv("POSTGRES_DB")
		instSettings.PgPass = os.Getenv("POSTGRES_PASSWORD")
		instSettings.PgUser = os.Getenv("POSTGRES_USER")

	})

	return instSettings
}
