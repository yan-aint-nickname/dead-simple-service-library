package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

// NOTE: I don't use struct to unmarshal in
//  because I kinda need a runtime access to the settings
//  and I can't predict what you might need in your application
//  you can use native viper to get access to the settings
func LoadSettingsHttp() (err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("local.http.env")
	viper.SetConfigType("env")
	viper.SetDefault("addr", "0.0.0.0:8080")
	viper.SetDefault("app_name", "Test App !Change Name!")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	return
}

func loadFiberConfigFromSettings() fiber.Config {
	return fiber.Config{
		AppName: viper.GetString("app_name"),
	}
}

func RunFromSettings(app *fiber.App) {
	log.Fatal(app.Listen(viper.GetString("addr")))
}

func CreateDefaultApp() *fiber.App {
	fiberConfig := loadFiberConfigFromSettings()
	return fiber.New(fiberConfig)
}
