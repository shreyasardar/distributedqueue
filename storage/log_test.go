package storage

import (
	"gominimq/models"
	"os"
	"path/filepath"
	"testing"
)

func TestAppend(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "test.log")

	log, err := NewLog(filename)
	if err != nil {
		t.Fatalf("failed to create log: %v", err)
	}
	defer log.Close()

	msg := models.Message{
		ID:        1,
		Key:       "user1",
		Value:     "hello",
		Timestamp: 123456,
	}

	err = log.Append(msg)
	if err != nil {
		t.Fatalf("failed to append message: %v", err)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	expected := `{"ID":1,"Key":"user1","Value":"hello","Timestamp":123456}` + "\n"

	if string(content) != expected {
		t.Errorf("expected %q, got %q", expected, string(content))
	}
}

func TestReadAll(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "test.log")

	log, err := NewLog(filename)
	if err != nil {
		t.Fatalf("failed to create log: %v", err)
	}
	defer log.Close()

	messages := []models.Message{
		{
			ID:        1,
			Key:       "user1",
			Value:     "hello",
			Timestamp: 100,
		},
		{
			ID:        2,
			Key:       "user2",
			Value:     "world",
			Timestamp: 200,
		},
	}

	for _, msg := range messages {
		err := log.Append(msg)
		if err != nil {
			t.Fatalf("failed to append message: %v", err)
		}
	}

	result, err := log.ReadAll()
	if err != nil {
		t.Fatalf("failed to read messages: %v", err)
	}

	if len(result) != len(messages) {
		t.Fatalf("expected %d messages, got %d", len(messages), len(result))
	}

	for i := range messages {
		if result[i] != messages[i] {
			t.Errorf("message %d does not match: expected %+v, got %+v",
				i, messages[i], result[i])
		}
	}
}
