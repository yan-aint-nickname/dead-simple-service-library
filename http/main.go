package http

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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

func loadSentry() error {
	if ok := viper.IsSet("sentry_dsn"); !ok {
		log.Fatal("sentry_dsn not set")
		os.Exit(1)
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("sentry_dsn"),
		EnableTracing: true,
		AttachStacktrace: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func getSentryMiddleware() gin.HandlerFunc {
	loadSentry()
	return sentrygin.New(sentrygin.Options{
		Repanic: true,
	})
}

func RunFromSettings(app *gin.Engine) {
	log.Fatal(app.Run(viper.GetString("addr")))
}

func CreateDefaultApp() *gin.Engine {
	app := gin.Default()

	app.Use(getSentryMiddleware())

	return app
}
