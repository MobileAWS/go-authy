package authy

import (
	"net/url"
	"testing"
)

func Test_TransformParams(t *testing.T) {
	str := transformParams(url.Values{
		"foo": []string{"unicorn", "rainbow"},
		"bar": []string{"pony"},
	})

	if str != "bar=pony&foo[]=unicorn&foo[]=rainbow" {
		t.Error("Unable to propery parse url params into a sorted string")
	}
}

func Test_VerifySignatureGet(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)

	result, err := api.VerifySignature(
		"DMuAaWa7hoYwficR8YAar18VOphbTWeRIYTi1UXVzSo=",
		"http://c9b4c941.ngrok.io/authy/authy-php/src/callback.php",
		"GET",
		url.Values{
			"device_uuid":     []string{"cea50e20-3aeb-0133-f92a-34363b620e52"},
			"callback_action": []string{"approval_request_status"},
			"uuid":            []string{"cdbabf40-1c65-0133-d113-34363b620e52"},
			"status":          []string{"approved"},
			"approval_request[transaction][details][Email Address]": []string{"jdoe@example.com"},
			"approval_request[transaction][device_details]":         []string{},
			"approval_request[transaction][device_geolocation]":     []string{""},
			"approval_request[transaction][device_signing_time]":    []string{"946641599"},
			"approval_request[transaction][encrypted]":              []string{"false"},
			"approval_request[transaction][flagged]":                []string{"false"},
			"approval_request[transaction][hidden_details][ip]":     []string{"1.1.1.1"},
			"approval_request[transaction][message]":                []string{"Request to Login"},
			"approval_request[transaction][reason]":                 []string{""},
			"approval_request[transaction][requester_details]":      []string{""},
			"approval_request[transaction][status]":                 []string{"approved"},
			"approval_request[transaction][uuid]":                   []string{"cdbabf40-1c65-0133-d113-34363b620e52"},
			"approval_request[transaction][created_at_time]":        []string{"946641599"},
			"approval_request[transaction][customer_uuid]":          []string{"25e026b36343-a29f-3310-bea3-02e05aec"},
			"approval_request[logos]":                               []string{""},
			"approval_request[expiration_timestamp]":                []string{"946641599"},
			"signature":                                             []string{`rzqf\/n08coE0Vi7IjbzAbt0IYMprJGAUx18kSJWE37K0mhvCGwepkm\/pSDXuSs+5kSUFK80L9RT7\/BZ7YwojSt5WhPnpRSImm5qKlvsNnGOPYCKVcFJxXCNJhtaztL\/2BjOMzdC5yNHH5uJIDGBhlb5fLVErsvauvxXWo\/Cj2STfITdSPULFz6XcbM1BDIriW7kP0GkELfUqE1iEuONEdhKYmPGolh3\/U4t8i0NYkQSPhbOGG1DZEsxhnxtelyBNOGK9sFojTsAg7dWesRYnyDkjTHZ1MvggdZwXo4qxphrY2Ve7+o04EHPZW9RPvakwl9yQ6rVsspVF\/xZT14BsgA==`},
			"authy_id":                                              []string{"1234"},
		},
		"1486660308",
	)

	if err != nil {
		t.Errorf("Signature Verification failed with error: %s", err.Error())
	}

	if !result {
		t.Error("Signatures doesn't match")
	}
}

func Test_VerifySignaturePost(t *testing.T) {
	api := NewAuthyAPI(data.APIKey)

	result, err := api.VerifySignature(
		"0Fl2yGPwgjhQCuiIyQXTaL56BL8ptyjso1kVgau5q1s=",
		"http://c9b4c941.ngrok.io/authy/authy-php/src/callback.php",
		"POST",
		url.Values{
			"device_uuid":     []string{"cea50e20-3aeb-0133-f92a-34363b620e52"},
			"callback_action": []string{"approval_request_status"},
			"uuid":            []string{"cdbabf40-1c65-0133-d113-34363b620e52"},
			"status":          []string{"approved"},
			"approval_request[transaction][details][Email Address]": []string{"jdoe@example.com"},
			"approval_request[transaction][device_details]":         []string{},
			"approval_request[transaction][device_geolocation]":     []string{""},
			"approval_request[transaction][device_signing_time]":    []string{"946641599"},
			"approval_request[transaction][encrypted]":              []string{"false"},
			"approval_request[transaction][flagged]":                []string{"false"},
			"approval_request[transaction][hidden_details][ip]":     []string{"1.1.1.1"},
			"approval_request[transaction][message]":                []string{"Request to Login"},
			"approval_request[transaction][reason]":                 []string{""},
			"approval_request[transaction][requester_details]":      []string{""},
			"approval_request[transaction][status]":                 []string{"approved"},
			"approval_request[transaction][uuid]":                   []string{"cdbabf40-1c65-0133-d113-34363b620e52"},
			"approval_request[transaction][created_at_time]":        []string{"946641599"},
			"approval_request[transaction][customer_uuid]":          []string{"25e026b36343-a29f-3310-bea3-02e05aec"},
			"approval_request[logos]":                               []string{""},
			"approval_request[expiration_timestamp]":                []string{"946641599"},
			"signature":                                             []string{`rzqf\/n08coE0Vi7IjbzAbt0IYMprJGAUx18kSJWE37K0mhvCGwepkm\/pSDXuSs+5kSUFK80L9RT7\/BZ7YwojSt5WhPnpRSImm5qKlvsNnGOPYCKVcFJxXCNJhtaztL\/2BjOMzdC5yNHH5uJIDGBhlb5fLVErsvauvxXWo\/Cj2STfITdSPULFz6XcbM1BDIriW7kP0GkELfUqE1iEuONEdhKYmPGolh3\/U4t8i0NYkQSPhbOGG1DZEsxhnxtelyBNOGK9sFojTsAg7dWesRYnyDkjTHZ1MvggdZwXo4qxphrY2Ve7+o04EHPZW9RPvakwl9yQ6rVsspVF\/xZT14BsgA==`},
			"authy_id":                                              []string{"1234"},
		},
		"1486660268",
	)

	if err != nil {
		t.Errorf("Signature Verification failed with error: %s", err.Error())
	}

	if !result {
		t.Error("Signatures doesn't match")
	}
}
