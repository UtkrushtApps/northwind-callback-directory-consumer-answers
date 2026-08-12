package service

import (
	"testing"

	"northwind/callbackdirectory/client"
)

type stubClient struct {
	payload client.CountryPayload
}

func (s stubClient) FetchCountry(country string) (client.CountryPayload, error) {
	return s.payload, nil
}

func TestSingleCallingCode(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghanistan", CallingCodes: []string{"93"}}}}}
	svc := NewDirectoryService(stub)
	result, err := svc.GetPhoneNumbers("Afghanistan", "656445445")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "+93 656445445" {
		t.Fatalf("expected +93 656445445, got %s", result)
	}
}

func TestMultipleCallingCodesUsesHighestIndexCode(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Puerto Rico", CallingCodes: []string{"1", "1787", "1939"}}}}}
	svc := NewDirectoryService(stub)
	result, err := svc.GetPhoneNumbers("Puerto Rico", "123456789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "+1939 123456789" {
		t.Fatalf("expected +1939 123456789, got %s", result)
	}
}

func TestUnknownCountryReturnsMinusOne(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{}}}
	svc := NewDirectoryService(stub)
	result, err := svc.GetPhoneNumbers("Atlantis", "5551234")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "-1" {
		t.Fatalf("expected -1, got %s", result)
	}
}
