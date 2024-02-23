package dead_simple_service_library

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
)

// Nasty hack for dealing with nested settings
// Addr: "0.0.0.0:8080" default: "0.0.0.0:<random_unused_port>"
type Settings struct {
	AppName string `mapstructure:"app_name"`
	Addr    string
}

func LoadSettings() (settings Settings, err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("local.http.env")
	viper.SetConfigType("env")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&settings)
	viper.Debug()

	return
}

func loadFiberConfigFromSettings(settings Settings) fiber.Config {
	return fiber.Config{
		AppName: settings.AppName,
	}
}

func RunFromSettings(app *fiber.App, settings Settings) {
	log.Fatal(app.Listen(settings.Addr))
}

func CreateDefaultApp(settings Settings) *fiber.App {
	fiberConfig := loadFiberConfigFromSettings(settings)
	log.Println(fiberConfig)
	app := fiber.New(fiberConfig)
	return app
}
