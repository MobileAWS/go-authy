package authy

import (
	"net/url"
	"testing"
	"time"
)

func Test_SendApprovalRequestValidation(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)
	user, err := api.RegisterUser(data.Email, data.CountryCode, data.PhoneNumber, url.Values{})

	_, err = api.SendApprovalRequest(user.ID, "", Details{}, Logos{}, 30, url.Values{})
	if err == nil {
		t.Error("Should have thrown a validation error")
	}
	if err.Error() != "message should not be empty" {
		t.Error("Should have thrown 'message should not be empty' error")
	}

	_, err = api.SendApprovalRequest(user.ID, "netflix and chill", Details{}, Logos{}, 0, url.Values{})
	if err == nil {
		t.Error("Should have thrown a validation error")
	}
	if err.Error() != "expiry time should be greater than zero" {
		t.Error("Should have thrown 'expiry time should be greater than zero' error")
	}

	_, err = api.SendApprovalRequest(user.ID, "netflix and chill", Details{}, Logos{"https://someurl.com/foo.png": ""}, 30, url.Values{})
	if err == nil {
		t.Error("Should have thrown a validation error")
	}
	if err.Error() != "logo resolution should not be empty" {
		t.Error("Should have thrown 'logo resolution should not be empty' error")
	}


	_, err = api.SendApprovalRequest(user.ID, "netflix and chill", Details{}, Logos{"https://someurl.com/foo.png": "foo"}, 30, url.Values{})
	if err == nil {
		t.Error("Should have thrown a validation error")
	}
	if err.Error() != "logo resolution should be either default, low, mid or high" {
		t.Error("Should have thrown 'logo resolution should be either default, low, mid or high' error")
	}

	_, err = api.SendApprovalRequest(user.ID, "netflix and chill", Details{}, Logos{"": "default"}, 30, url.Values{})
	if err == nil {
		t.Error("Should have thrown a validation error")
	}
	if err.Error() != "logo url should not be empty" {
		t.Error("Should have thrown 'logo url should not be empty' error")
	}
}

func Test_SendApprovalRequest(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)
	user, err := api.RegisterUser(data.Email, data.CountryCode, data.PhoneNumber, url.Values{})

	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, Logos{}, 30, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if !approvalRequest.Valid() {
		t.Error("Apprval request should be valid.")
	}
}

func Test_FindApprovalRequest(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)

	user, err := api.RegisterUser(data.Email, data.CountryCode, data.PhoneNumber, url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, Logos{}, 30, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if !approvalRequest.Valid() {
		t.Error("Apprval request should be valid.")
	}

	uuid := approvalRequest.UUID
	approvalRequest, err = api.FindApprovalRequest(uuid, url.Values{})

	if err != nil {
		t.Error("External error found", err)
	}

	if approvalRequest.Status != "pending" {
		t.Error("Approval request status is wrong")
	}

	if uuid != approvalRequest.UUID {
		t.Error("Approval request doesn't match.")
	}
}

func Test_WaitForApprovalRequest(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)

	user, err := api.RegisterUser(data.Email, data.CountryCode, data.PhoneNumber, url.Values{})
	approvalRequest, err := api.SendApprovalRequest(user.ID, "please approve this", Details{"data1": "value1"}, Logos{}, 30, url.Values{})

	if err != nil {
		t.Error("error found", err)
	}

	if !approvalRequest.Valid() {
		t.Error("Apprval request should be valid.")
	}

	now := time.Now()
	status, err := api.WaitForApprovalRequest(approvalRequest.UUID, 1*time.Second, url.Values{"user_ip": {"234.78.25.2"}})
	elapsedTime := time.Now().Sub(now)

	if err != nil {
		t.Error("error found", err)
	}

	if elapsedTime < 1 {
		t.Error("max duration not reached")
	}

	if status != OneTouchStatusExpired {
		t.Error("expired status expected")
	}
}
