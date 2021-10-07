package database

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	GsheetService *sheets.Service
)

func GsheetConnect() {
	// Fetch the service account key JSON file contents
	pathX, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	dir := path.Dir(pathX)

	exPath := filepath.FromSlash(dir + "/gcpkey.json")
	abs, err := filepath.Abs(exPath)
	if err == nil {
		fmt.Println("Absolute:", abs)
	}

	ctx := context.Background()
	b, err := ioutil.ReadFile(abs)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	SetUpService(srv)

}

func SetUpService(srv *sheets.Service) {
	GsheetService = srv
}

func GetService() *sheets.Service {
	return GsheetService
}
