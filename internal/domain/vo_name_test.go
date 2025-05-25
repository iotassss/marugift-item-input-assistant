package domain_test

import (
	"testing"

	"github.com/iotassss/marugift-item-input-assistant/internal/domain"
)

func TestName_NewName(t *testing.T) {
	_, err := domain.NewName("Valid Name")
	if err != nil {
		t.Fatalf("failed to create valid name: %v", err)
	}

	_, err = domain.NewName("")
	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}
}
