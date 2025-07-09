// バックエンドアプリケーションのテストファイル
package main

import (
	"testing"
)

func TestGreetUser(t *testing.T) {
	expected := "Hello, Test User! Welcome to our backend API."
	actual := GreetUser("Test User")

	if actual != expected {
		t.Errorf("GreetUser() = %v, want %v", actual, expected)
	}
}

func TestMultiply(t *testing.T) {
	expected := 15
	actual := Multiply(3, 5)

	if actual != expected {
		t.Errorf("Multiply(3, 5) = %v, want %v", actual, expected)
	}
}

func TestMultiplyZero(t *testing.T) {
	expected := 0
	actual := Multiply(5, 0)

	if actual != expected {
		t.Errorf("Multiply(5, 0) = %v, want %v", actual, expected)
	}
}
