package main

import (
	"fmt"
	"regexp"
	"strings"
)

func validateID(id string) error {
	if id == "" {
		return fmt.Errorf("ID cannot be empty")
	}

	matched, _ := regexp.MatchString("^[a-z0-9-]+$", id)
	if !matched {
		return fmt.Errorf("ID must contain only lowercase letters, numbers, and hyphens")
	}

	return nil
}

func normalizeID(input string) string {
	normalized := strings.ToLower(input)
	normalized = strings.ReplaceAll(normalized, " ", "-")
	normalized = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(normalized, "")
	normalized = regexp.MustCompile(`-+`).ReplaceAllString(normalized, "-")
	normalized = strings.Trim(normalized, "-")
	return normalized
}

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}
