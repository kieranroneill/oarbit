import os
from collections.abc import AsyncIterator
from contextlib import asynccontextmanager
from pathlib import Path

from fastapi import FastAPI
from sqlalchemy.ext.asyncio import (
    AsyncSession,
    async_sessionmaker,
    create_async_engine,
)

from api import routers, utilities


def _create_app() -> FastAPI:
    app_title = os.environ["APP_TITLE"]
    _app = FastAPI(lifespan=_lifespan, title=f"{app_title} API")

    _app.include_router(routers.boat.router)
    _app.include_router(routers.chat.router)

    return _app


@asynccontextmanager
async def _lifespan(_app: FastAPI) -> AsyncIterator[None]:
    database_engine = create_async_engine(utilities.database.url(), pool_pre_ping=True)

    # attach the initialized agent, database (session factory) and data path so it is available for each route
    # via dependencies
    _app.state.database = async_sessionmaker(
        bind=database_engine,
        class_=AsyncSession,
        expire_on_commit=False,
        autoflush=False,
    )
    _app.state.agent = utilities.agent.initialize()
    _app.state.data_path = Path(__file__).resolve().parent.parent / "data"

    yield
    # cleanup
    await database_engine.dispose()


app = _create_app()
