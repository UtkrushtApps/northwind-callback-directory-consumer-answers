import httpx


class CountriesClient:
    def __init__(self, config):
        self._config = config

    def fetch_country(self, country: str) -> dict:
        response = httpx.get(
            self._config.upstream_base_url,
            params={"name": country},
            timeout=self._config.request_timeout_seconds,
        )
        response.raise_for_status()
        return response.json()
