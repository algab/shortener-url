package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/xid"
)

var host string

// Post Shortener
type Post struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Short string `json:"short"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	} else {
		host = os.Getenv("HOST")
	}
}

// CreateShort Route
func CreateShort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body Post
	_ = json.NewDecoder(r.Body).Decode(&body)
	guid := xid.New().String()
	body.ID = guid
	body.Short = fmt.Sprintf("%s/%s", host, guid)
	SetCache(body)
	json.NewEncoder(w).Encode(body)
}
