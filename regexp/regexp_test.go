package utils

import "testing"

func TestCheckTypeRange(t *testing.T) {
	checkType := "http"
	validTypeRange := []string{"HTTP", "HTTPS", "WEBSOCKET", "AJP", "GRPC", "GRPCS"}
	isExist := CheckTypeRange(validTypeRange, checkType)
	if isExist {
		t.Logf("%s is exist!", checkType)
	} else {
		t.Errorf("%s is not exist!", checkType)
	}
}
