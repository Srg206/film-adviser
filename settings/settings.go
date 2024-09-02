package settings

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Struct that contains settings
type Settings struct {
	TgSenderToken   string
	TgReceiverToken string
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
	})

	return instSettings
}
