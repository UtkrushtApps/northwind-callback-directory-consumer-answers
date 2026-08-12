from fastapi import APIRouter, Depends, HTTPException, Query

from app.client import CountriesClient
from app.config import load_config
from app.service import DirectoryService

router = APIRouter()

_config = load_config()
_service = DirectoryService(CountriesClient(_config))


def get_service() -> DirectoryService:
    return _service


def _is_blank(value: str | None) -> bool:
    return value is None or value.strip() == ""


@router.get("/phone-numbers")
def phone_numbers(
    country: str | None = Query(None),
    phone: str | None = Query(None),
    service: DirectoryService = Depends(get_service),
):
    if _is_blank(country) or _is_blank(phone):
        raise HTTPException(status_code=400, detail="country and phone are required")

    resolved = service.getPhoneNumbers(country, phone)
    return {"result": resolved}
