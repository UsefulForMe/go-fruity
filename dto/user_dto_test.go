package dto

import "testing"

func Test_bad_request_dto(t *testing.T) {
	loginUserRequest := LoginUserRequest{
		PhoneNumber: "",
		IdToken:     "",
	}
	if err := loginUserRequest.Validate(); err != nil {
		if err.Code != 400 {
			t.Errorf("Expected error code 400, got %d", err.Code)
		}
	}
}

func Test_valid_request_dto(t *testing.T) {
	loginUserRequest := LoginUserRequest{
		PhoneNumber: "0338613062",
		IdToken:     "123456",
	}
	if err := loginUserRequest.Validate(); err != nil {
		t.Errorf("Expected error code to be nil, got %v", err)
	}
}
