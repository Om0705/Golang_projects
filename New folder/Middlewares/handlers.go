package handlers

import(
	"database/sql"
	"g9ithub.com/joho/godotenv"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB{
	err := gotdotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sql.Open("postgress", os.Getenv("POSTGRESS_URL"))
} 