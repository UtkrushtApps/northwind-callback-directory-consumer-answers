from fastapi.testclient import TestClient

from app.main import create_app
from app.routes import get_service
from app.service import DirectoryService


class StubClient:
    def __init__(self, payload):
        self._payload = payload

    def fetch_country(self, country):
        return self._payload


def build_client(payload):
    app = create_app()
    service = DirectoryService(StubClient(payload))
    app.dependency_overrides[get_service] = lambda: service
    return TestClient(app)


def test_endpoint_single_calling_code():
    client = build_client({"data": [{"name": "Afghanistan", "callingCodes": ["93"]}]})
    response = client.get("/phone-numbers", params={"country": "Afghanistan", "phone": "656445445"})
    assert response.status_code == 200
    assert response.json() == {"result": "+93 656445445"}


def test_endpoint_multiple_calling_codes_uses_highest_index_code():
    client = build_client({"data": [{"name": "Puerto Rico", "callingCodes": ["1", "1787", "1939"]}]})
    response = client.get("/phone-numbers", params={"country": "Puerto Rico", "phone": "123456789"})
    assert response.status_code == 200
    assert response.json() == {"result": "+1939 123456789"}


def test_endpoint_unknown_country_returns_minus_one_successfully():
    client = build_client({"data": []})
    response = client.get("/phone-numbers", params={"country": "Atlantis", "phone": "5551234"})
    assert response.status_code == 200
    assert response.json() == {"result": "-1"}


def test_endpoint_rejects_missing_country():
    client = build_client({"data": []})
    response = client.get("/phone-numbers", params={"phone": "5551234"})
    assert response.status_code == 400
    assert response.json() == {"detail": "country and phone are required"}


def test_endpoint_rejects_blank_phone():
    client = build_client({"data": []})
    response = client.get("/phone-numbers", params={"country": "Afghanistan", "phone": "   "})
    assert response.status_code == 400
    assert response.json() == {"detail": "country and phone are required"}
