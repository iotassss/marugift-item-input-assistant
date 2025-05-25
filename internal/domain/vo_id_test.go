package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
)

func TestID_NewID(t *testing.T) {
	_, err := domain.NewID(0)
	if err == nil {
		t.Fatal("expected error for ID 0, got nil")
	}

	_, err = domain.NewID(1)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestID_NewDraftID(t *testing.T) {
	id := domain.NewDraftID()
	if id.Value() != 0 {
		t.Fatalf("expected ID value 0, got %d", id.Value())
	}
}

func TestID_Value(t *testing.T) {
	id, err := domain.NewID(1)
	if err != nil {
		t.Fatalf("failed to create ID: %v", err)
	}
	if id.Value() != 1 {
		t.Fatalf("expected ID value 1, got %d", id.Value())
	}

	id, err = domain.NewID(100)
	if err != nil {
		t.Fatalf("failed to create ID: %v", err)
	}
	if id.Value() != 100 {
		t.Fatalf("expected ID value 100, got %d", id.Value())
	}
}
