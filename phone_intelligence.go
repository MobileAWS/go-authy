package authy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PhoneIntelligence wraps the phone information provided by Authy API
type PhoneIntelligence struct {
	HTTPResponse *http.Response
	Message      string
	Type         string
	Provider     string
	Ported       bool
	Success      bool
}

// NewPhoneIntelligence creates a new instance of PhoneIntelligence
func NewPhoneIntelligence(response *http.Response) (*PhoneIntelligence, error) {
	info := &PhoneIntelligence{HTTPResponse: response}
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		Logger.Println("Error reading from API:", err)
		return info, err
	}

	err = json.Unmarshal(body, &info)
	if err != nil {
		Logger.Println("Error parsing JSON:", err)
		return info, err
	}

	return info, nil
}
