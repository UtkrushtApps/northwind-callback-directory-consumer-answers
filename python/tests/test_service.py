from app.service import DirectoryService


class StubClient:
    def __init__(self, payload):
        self._payload = payload

    def fetch_country(self, country):
        return self._payload


def test_single_calling_code():
    client = StubClient({"data": [{"name": "Afghanistan", "callingCodes": ["93"]}]})
    service = DirectoryService(client)
    assert service.getPhoneNumbers("Afghanistan", "656445445") == "+93 656445445"


def test_multiple_calling_codes_uses_highest_index_code():
    client = StubClient({"data": [{"name": "Puerto Rico", "callingCodes": ["1", "1787", "1939"]}]})
    service = DirectoryService(client)
    assert service.getPhoneNumbers("Puerto Rico", "123456789") == "+1939 123456789"


def test_unknown_country_returns_minus_one():
    client = StubClient({"data": []})
    service = DirectoryService(client)
    assert service.getPhoneNumbers("Atlantis", "5551234") == "-1"
