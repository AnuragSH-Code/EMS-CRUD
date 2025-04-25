package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestValidateEmployeePayload(t *testing.T) {
	tests := []struct {
		name    string
		payload EmployeePayload
		wantErr bool
	}{
		{
			name: "valid",
			payload: EmployeePayload{
				Firstname:  strPtr("Nausheen"),
				Lastname:   strPtr("Javed"),
				Role:       strPtr("People"),
				Department: strPtr("HR"),
				Email:      strPtr("Nausheen@Javed.com"),
				ContactNo:  strPtr("1234567890"),
				Manager:    strPtr("Stuti Tyagi"),
			},
			wantErr: false,
		},
		{
			name:    "missing fields",
			payload: EmployeePayload{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateEmployeePayload(&tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("got error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadJSON_Malformed(t *testing.T) {
	req := httptest.NewRequest("POST", "/", bytes.NewReader([]byte("{invalid}")))
	var data map[string]interface{}
	err := readJSON(httptest.NewRecorder(), req, &data)
	if err == nil {
		t.Error("expected error for malformed JSON")
	}
}
