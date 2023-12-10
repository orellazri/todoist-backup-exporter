package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BackupResponse struct {
	Url     string `json:"url"`
	Version string `json:"version"`
}

func main() {
	// Parse environment variables
	token := os.Getenv("TODOIST_TOKEN")
	if len(token) == 0 {
		log.Fatal("TODOIST_TOKEN must be set")
	}

	destDir := os.Getenv("DEST_DIR")
	if len(destDir) == 0 {
		destDir = "/exports"
	}

	destFilenameIncludeDate := os.Getenv("DEST_FILENAME_INCLUDE_DATE") == "true"

	// Get list of backups
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.todoist.com/sync/v9/backups/get", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get backups list: %s", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var backupsResponse []BackupResponse
	if err := decoder.Decode(&backupsResponse); err != nil {
		log.Fatalf("Failed to decode response: %s", err)
	}

	if len(backupsResponse) == 0 {
		log.Fatal("No backups found")
	}

	// Download last backup
	backup := backupsResponse[0]
	log.Printf("Downloading backup version %q from %q", backup.Version, backup.Url)

	req, err = http.NewRequest("GET", backup.Url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("Failed to download backup: %s", err)
	}
	defer resp.Body.Close()

	// Save backup to file
	destFilename := "Todoist_Backup.zip"
	if destFilenameIncludeDate {
		replacer := strings.NewReplacer(" ", "_", ":", "-")
		backupVersionClean := replacer.Replace(backup.Version)
		destFilename = fmt.Sprintf("Todoist_Backup_%s.zip", backupVersionClean)
	}
	destFilepath := filepath.Join(destDir, destFilename)
	out, err := os.Create(destFilepath)
	if err != nil {
		log.Fatalf("Failed to create output file: %s", err)
	}
	defer out.Close()
	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Fatalf("Failed to write response to output file: %s", err)
	}

	log.Printf("Backup saved to %q", destFilepath)
}
