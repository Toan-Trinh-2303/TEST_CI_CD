package service

import (
	"context"
	"testing"

	"ci-cd-go-learn/internal/model"
)

func TestCreateUser_EmptyName(t *testing.T) {
	svc := NewUserService(nil)

	_, err := svc.CreateUser(context.Background(), model.CreateUserRequest{
		Name:  "",
		Email: "test@example.com",
	})
	if err == nil || err.Error() != "name is required" {
		t.Fatalf("expected 'name is required', got %v", err)
	}
}

func TestCreateUser_EmptyEmail(t *testing.T) {
	svc := NewUserService(nil)

	_, err := svc.CreateUser(context.Background(), model.CreateUserRequest{
		Name:  "Phong",
		Email: "",
	})
	if err == nil || err.Error() != "email is required" {
		t.Fatalf("expected 'email is required', got %v", err)
	}
}

func TestCreateUser_TrimsWhitespace(t *testing.T) {
	svc := NewUserService(nil)

	_, err := svc.CreateUser(context.Background(), model.CreateUserRequest{
		Name:  "   ",
		Email: "test@example.com",
	})
	if err == nil || err.Error() != "name is required" {
		t.Fatalf("expected 'name is required' after trim, got %v", err)
	}
}
