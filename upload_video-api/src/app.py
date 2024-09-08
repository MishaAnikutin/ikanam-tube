from fastapi import FastAPI
from .controllers.rest.v1 import router


def create_app() -> FastAPI:
    app = FastAPI()
    app.include_router(router)
    return app
