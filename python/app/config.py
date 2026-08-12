import os
from dataclasses import dataclass


@dataclass(frozen=True)
class Config:
    upstream_base_url: str = "https://jsonmock.hackerrank.com/api/countries"
    request_timeout_seconds: float = 5.0


def load_config() -> Config:
    return Config(
        upstream_base_url=os.getenv(
            "NORTHWIND_UPSTREAM_BASE_URL",
            "https://jsonmock.hackerrank.com/api/countries",
        ),
        request_timeout_seconds=float(os.getenv("NORTHWIND_REQUEST_TIMEOUT_SECONDS", "5.0")),
    )
