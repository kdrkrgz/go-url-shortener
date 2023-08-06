package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	mode := os.Getenv("Mode")
	if mode != "Production" {
		fmt.Println("Loading .env file...")
		err := godotenv.Load("../env.example")
		if err != nil {
			panic("Error loading .env file")
		}

	}
}
