package config

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var Bucket = os.Getenv("AWS_BUCKET")
var Endpoint = os.Getenv("AWS_ENDPOINT")
var Region = os.Getenv("AWS_REGION")

func init() {
	if !strings.HasSuffix(Endpoint, "/") {
		Endpoint += "/"
	}
}
