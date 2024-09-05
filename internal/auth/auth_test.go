package auth

import (
	"errors"
	"net/http"
	"testing"
)

type apiKeyTest struct {
	headerKey string;
	headerValue string;
	expectedResponse string;
	errorMessage error;
}

var addTests = []apiKeyTest{
	apiKeyTest{"Authorization", "asdf", "", errors.New("malformed authorization header")},
	apiKeyTest{"something", "asdf", "", errors.New("no authorization header included")},
	apiKeyTest{"Authorization", "ApiKey asdfasdf", "asdfasdf", nil},

}

func TestGetAPIKey(t *testing.T) {

	for _, test := range addTests{
		testHeader := http.Header{}
		testHeader.Set(test.headerKey, test.headerValue)
		outputResponse, errorResponse := GetAPIKey(testHeader)
		if errorResponse == nil || test.errorMessage == nil {
			if outputResponse != test.expectedResponse || errorResponse != test.errorMessage {
				t.Errorf("Outputs %s, %v not equal to expected %s, %v", outputResponse, errorResponse, test.expectedResponse, test.errorMessage)
			}
		} else if outputResponse != test.expectedResponse || errorResponse.Error() != test.errorMessage.Error() {
			t.Errorf("Outputs %s, %s not equal to expected %s, %s", outputResponse, errorResponse, test.expectedResponse, test.errorMessage.Error())
		}
	}
}
