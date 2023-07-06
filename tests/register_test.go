package tests

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// package tests
//
//
// import (
// 	"bytes"
// 	"main/bootstrap"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )
//
// func TestRegister(t *testing.T, app *fiber.App) {
// // 	registerPayload := []byte(`{"username": "testuser", "password": "testpassword"}`)
// // 	registerRequest, _ := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(registerPayload))
// // 	registerResponse := httptest.NewRecorder()
// // 	registerResponse2 := httptest.NewRecorder()
// 	if 1 != 2 {
// 		t.Errorf("Ожидалось сообщение '%s', получено '%s'", "1", "2")
// 		if registerResponse != registerResponse2 {
// 			t.Errorf("Ожидалось сообщение '%s', получено '%s'", "1", "2")
// 		}
// 	}
// // 	app.Test(registerRequest, 2)
//
// }
