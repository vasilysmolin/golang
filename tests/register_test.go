package tests

import (
	"bytes"
	"main/bootstrap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T, app *fiber.App) {

	// Тест регистрации
	registerPayload := []byte(`{"username": "testuser", "password": "testpassword"}`)
	registerRequest, _ := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(registerPayload))
	registerResponse := httptest.NewRecorder()
	registerResponse2 := httptest.NewRecorder()
	if 1 != 2 {
		t.Errorf("Ожидалось сообщение '%s', получено '%s'", "1", "2")
		if registerResponse != registerResponse2 {
			t.Errorf("Ожидалось сообщение '%s', получено '%s'", "1", "2")
		}
	}
	app.Test(registerRequest, 2)
	// var registerRes Response
	// json.NewDecoder(registerResponse.Body).Decode(&registerRes)
	//
	// expectedRegisterMessage := "Регистрация прошла успешно"
	//
	//	if registerRes.Message != expectedRegisterMessage {
	//	 t.Errorf("Ожидалось сообщение '%s', получено '%s'", expectedRegisterMessage, registerRes.Message)
	//	}
	//
	// // Тест аутентификации
	// authPayload := []byte(`{"username": "testuser", "password": "testpassword"}`)
	// authRequest, _ := http.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(authPayload))
	// authResponse := httptest.NewRecorder()
	// app.Test(authRequest, authResponse)
	//
	// var authRes Response
	// json.NewDecoder(authResponse.Body).Decode(&authRes)
	//
	// expectedAuthMessage := "Аутентификация прошла успешно"
	//
	//	if authRes.Message != expectedAuthMessage {
	//	 t.Errorf("Ожидалось сообщение '%s', получено '%s'", expectedAuthMessage, authRes.Message)
	//	}
}
