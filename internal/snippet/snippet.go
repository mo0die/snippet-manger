package snippet

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type Snippet struct {
	Name        string
	Content     string
	Description string
	Tags        []string
	CreatedAt   time.Time
	ID          int64
}

var nextID int64 = 0

func (s Snippet) String() string {
	tagStr := "No Tags"
	if len(s.Tags) > 0 {
		tagStr = strings.Join(s.Tags, ", ")
	}
	return fmt.Sprintf("Name: %s\nDescription: %s\nTags: %s\nCreated: %s\nContent: %s\nID: %d",
		s.Name,
		s.Description,
		tagStr,
		s.CreatedAt.Local().String(),
		s.Content,
		s.ID)
}

func SetNextId(snippets []Snippet) {
	var maxID int64 = 0
	for _, s := range snippets {
		if s.ID > maxID {
			maxID = s.ID
		}
	}
	fmt.Println(maxID)
	atomic.StoreInt64(&nextID, maxID)
}

func New(name string, content string, tags []string, description string) Snippet {
	return Snippet{
		name,
		content,
		description,
		tags,
		time.Now(),
		atomic.AddInt64(&nextID, 1),
	}
}

func ListSnippets(snippets []Snippet) {
	fmt.Println("==============")
	for _, snippet := range snippets {
		fmt.Println(snippet.String())
		fmt.Println("===============")
	}
}

func Get(id string, snippets []Snippet) (Snippet, error) {
	idConverted := convertToInt64(id)
	for _, s := range snippets {
		if idConverted == s.ID {
			return s, nil
		}
	}
	return Snippet{}, errors.New("cannot fine id")
}

func Delete(id string, snippets []Snippet) ([]Snippet, error) {
	idConverted := convertToInt64(id)
	foundId := -1

	for i, s := range snippets {
		if s.ID == idConverted {
			foundId = i
			break
		}
	}
	if foundId == -1 {
		return snippets, errors.New("snippet not found: " + id)
	}

	updateSnippets := slices.Delete(snippets, foundId, foundId+1)

	ListSnippets(updateSnippets)
	return updateSnippets, nil
}

func convertToInt64(num string) int64 {
	idConverted, err := strconv.ParseInt(num, 0, 64)
	if err != nil {
		fmt.Println("cannot convert value, please pass a int: " + err.Error())
	}
	return idConverted
}
