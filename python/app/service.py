class DirectoryService:
    def __init__(self, client):
        self._client = client

    def getPhoneNumbers(self, country: str, phone: str) -> str:
        payload = self._client.fetch_country(country)
        records = payload.get("data") or []
        if not records:
            return "-1"

        calling_codes = records[0].get("callingCodes") or []
        if not calling_codes:
            return "-1"

        calling_code = calling_codes[-1]
        return "+" + calling_code + " " + phone
