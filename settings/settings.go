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
	TgSaverToken    string
	TgReminderToken string
	PgIp            string
	PgPort          int
	PgDb            string
	PgPass          string
	PgUser          string
}

var instSettings *Settings
var once sync.Once

// GetSettings is a func to get settings instance (Settings is a Singleton)
func GetSettings() *Settings {

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Could not load .env file !")
		}
		instSettings = &Settings{}
		instSettings.TgReminderToken = os.Getenv("TG_REMINDER_TOKEN")
		instSettings.TgSaverToken = os.Getenv("TG_SAVER_TOKEN")
		instSettings.PgIp = os.Getenv("POSTGRES_HOST")
		instSettings.PgPort, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		instSettings.PgDb = os.Getenv("POSTGRES_DB")
		instSettings.PgPass = os.Getenv("POSTGRES_PASSWORD")
		instSettings.PgUser = os.Getenv("POSTGRES_USER")

	})

	return instSettings
}
