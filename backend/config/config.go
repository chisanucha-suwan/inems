package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/lpernett/godotenv"
)

func Load(filenames ...string) {
	_, file, _, _ := runtime.Caller(0)

	for _, filename := range filenames {
		fmt.Println("loading env from ", filename)
		err := godotenv.Load(filepath.Join(filepath.Dir(file), filename))
		if err != nil {
			fmt.Println("Error loading .env file:", err)
		}
	}

	// load .env.local & .env into os environment
	godotenv.Load(filepath.Join(filepath.Dir(file), ".env.local"))
	godotenv.Load(filepath.Join(filepath.Dir(file), ".env"))
	godotenv.Load()

	// check environment variable
	env := os.Getenv("APP_ENVELOPMENT")
	if env == "development" {
		godotenv.Load(filepath.Join(filepath.Dir(file), ".env.development"))
	} else if env == "production" {
		godotenv.Load(filepath.Join(filepath.Dir(file), ".env.production"))
	} else {
		fmt.Println("WARNING: Using default .env, environment variable 'ENV' not set or invalid.")
	}
}
