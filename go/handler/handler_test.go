package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"northwind/callbackdirectory/client"
	"northwind/callbackdirectory/service"
)

type stubClient struct {
	payload client.CountryPayload
}

func (s stubClient) FetchCountry(country string) (client.CountryPayload, error) {
	return s.payload, nil
}

func buildHandler(payload client.CountryPayload) *Handler {
	stub := stubClient{payload: payload}
	svc := service.NewDirectoryService(stub)
	return NewHandler(svc)
}

func performRequest(h *Handler, target string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)
	return rec
}

func decodeBody(t *testing.T, rec *httptest.ResponseRecorder) phoneNumbersResponse {
	t.Helper()
	var body phoneNumbersResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response was not valid JSON: %v", err)
	}
	return body
}

func TestEndpointKnownCountry(t *testing.T) {
	h := buildHandler(client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghanistan", CallingCodes: []string{"93"}}}})
	rec := performRequest(h, "/phone-numbers?country=Afghanistan&phone=656445445")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	body := decodeBody(t, rec)
	if body.Result != "+93 656445445" {
		t.Fatalf("expected +93 656445445, got %s", body.Result)
	}
}

func TestEndpointMultipleCallingCodes(t *testing.T) {
	h := buildHandler(client.CountryPayload{Data: []client.CountryRecord{{Name: "Puerto Rico", CallingCodes: []string{"1", "1787", "1939"}}}})
	rec := performRequest(h, "/phone-numbers?country=Puerto%20Rico&phone=123456789")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	body := decodeBody(t, rec)
	if body.Result != "+1939 123456789" {
		t.Fatalf("expected +1939 123456789, got %s", body.Result)
	}
}

func TestEndpointUnknownCountry(t *testing.T) {
	h := buildHandler(client.CountryPayload{Data: []client.CountryRecord{}})
	rec := performRequest(h, "/phone-numbers?country=Atlantis&phone=5551234")

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	body := decodeBody(t, rec)
	if body.Result != "-1" {
		t.Fatalf("expected -1, got %s", body.Result)
	}
}

func TestEndpointRejectsMissingCountry(t *testing.T) {
	h := buildHandler(client.CountryPayload{Data: []client.CountryRecord{}})
	rec := performRequest(h, "/phone-numbers?phone=5551234")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
	body := decodeBody(t, rec)
	if body.Error != "country and phone are required" {
		t.Fatalf("unexpected error body: %s", body.Error)
	}
}

func TestEndpointRejectsBlankPhone(t *testing.T) {
	h := buildHandler(client.CountryPayload{Data: []client.CountryRecord{}})
	rec := performRequest(h, "/phone-numbers?country=Afghanistan&phone=%20%20%20")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
	body := decodeBody(t, rec)
	if body.Error != "country and phone are required" {
		t.Fatalf("unexpected error body: %s", body.Error)
	}
}
