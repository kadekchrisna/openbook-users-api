package envutils

import "github.com/joho/godotenv"

// Load env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env")
	}
}
