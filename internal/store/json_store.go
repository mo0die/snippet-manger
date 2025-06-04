package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"snippet-manger/internal/snippet"
	"sync"
)

const (
	snippetsFileName = "snippets.json"
	configDirName    = ".snippet-manger"
)

var mu sync.Mutex

func GetSnippetsFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("failed to get users homedir: " + err.Error())
	}

	configPath := filepath.Join(homeDir, configDirName)
	return filepath.Join(configPath, snippetsFileName), nil
}

func LoadSnippets() ([]snippet.Snippet, error) {
	mu.Lock()

	defer mu.Unlock()

	filePath, err := GetSnippetsFilePath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []snippet.Snippet{}, nil
	} else if err != nil {
		return nil, errors.New("failed to check snippet file existence: " + err.Error())
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("failed to read file")
	}

	var snippets []snippet.Snippet

	if err := json.Unmarshal(data, &snippets); err != nil {
		return nil, errors.New("failed to load snippets data")
	}
	return snippets, nil
}

func SaveSnippet(snippets []snippet.Snippet) error {
	mu.Lock()
	defer mu.Unlock()

	filePath, err := GetSnippetsFilePath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.New("failed to create config directory" + err.Error())
	}
	data, err := json.MarshalIndent(snippets, "", "  ")
	if err != nil {
		return errors.New("failed to Marshall snippets Data")
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return errors.New("failed to write snippets file" + err.Error())
	}
	return nil
}
