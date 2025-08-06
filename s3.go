package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

const (
	S3Bucket  = "s3://cinema-web"
	LocalFile = "data.json"
)

func downloadFromS3(s3Bucket string) error {
	cmd := exec.Command("aws", "s3", "cp", s3Bucket+"/data.json", LocalFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download from S3: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Downloaded data from %s\n", s3Bucket)
	return nil
}

func uploadToS3(s3Bucket string) error {
	cmd := exec.Command("aws", "s3", "cp", LocalFile, s3Bucket+"/data.json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Uploaded data to %s\n", s3Bucket)
	return nil
}

func loadData() (*CinemaData, error) {
	file, err := os.ReadFile(LocalFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", LocalFile, err)
	}

	var data CinemaData
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &data, nil
}

func saveData(data *CinemaData) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	if err := os.WriteFile(LocalFile, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %v", LocalFile, err)
	}

	return nil
}
