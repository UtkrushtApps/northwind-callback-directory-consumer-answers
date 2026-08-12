# Solution Steps

1. Read the upstream base URL and request timeout from configuration instead of hard-coding the countries API URL in the HTTP client.

2. Build upstream requests with query-parameter APIs (`params`, `url.Values`, `UriComponentsBuilder`, etc.) so country names such as `Puerto Rico` or `United States` are URL-encoded correctly.

3. Update `getPhoneNumbers` in the service layer to keep the required domain rules: return `-1` when the upstream `data` array is empty, use the last/highest-index calling code when multiple codes exist, and format known-country results exactly as `+<CallingCode> <PhoneNumber>`.

4. Add handler/controller validation before calling the service: reject missing, null, or whitespace-only `country` or `phone` values with HTTP 400 and a clear JSON error.

5. Change successful endpoint responses from raw text to a stable JSON response contract, returning `{ "result": "..." }` for all successful resolutions, including `{ "result": "-1" }` for unknown countries.

6. Return JSON for framework-level errors as well, so clients always receive machine-parseable responses.

7. Fix framework dependency injection/test wiring where needed; for example, FastAPI must use `Depends(get_service)` so tests can override the service and avoid network calls.

8. Extend service tests to cover the three sample behaviors offline: a single calling code, a multi-word/multi-code country using the highest-index code, and an unknown country returning `-1`.

9. Extend endpoint tests using framework test clients (`TestClient`, `supertest`, `httptest`, `MockMvc`) to assert status codes and JSON response bodies for the same three successful cases, plus validation failures for missing/blank inputs.

10. Use stubs/mocks in all tests so no test reaches the real upstream network.

