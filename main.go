package dead_simple_service_library

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"log"
)

type Settings struct {
	Http fiber.Config
	Port string `mapstructure:"PORT"`
}

func LoadSettings() (Settings, error) {
	// Set the file name of the configurations file
	viper.SetConfigName("local.env")
	viper.SetConfigType("env")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return Settings{}, err
	}

	settings := Settings{}

	err := viper.Unmarshal(&settings)

	return settings, err
}


func RunFromSettings(app *fiber.App, settings Settings) {
	log.Fatal(app.Listen(settings.Port))
}

func CreateDefaultApp(settings Settings) *fiber.App {
	app := fiber.New(settings.Http)
	return app
}
