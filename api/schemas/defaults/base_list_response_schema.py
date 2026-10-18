from typing import Generic, TypeVar

from pydantic import Field

from api.schemas.defaults import BaseSchema

T = TypeVar("T")


class BaseListResponseSchema(BaseSchema, Generic[T]):
    items: list[T] = Field(default_factory=list)
