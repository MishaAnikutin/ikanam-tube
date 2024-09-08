import uvicorn
from dishka.integrations.fastapi import setup_dishka

from src.app import create_app
from src.ioc import container
from src.settings import AppConfig


def main():
    app = create_app()
    setup_dishka(container, app)
    return app


if __name__ == '__main__':
    uvicorn.run(main(), host=AppConfig.host, port=AppConfig.port)
