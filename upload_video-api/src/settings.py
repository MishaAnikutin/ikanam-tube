import os
import logging
from typing import Optional

from dotenv import load_dotenv
from dataclasses import dataclass
from sqlalchemy.engine import URL


dotenv_path = os.path.join(os.path.dirname(__file__), '.env')

if os.path.exists(dotenv_path):
    load_dotenv(dotenv_path)


@dataclass
class NGINXConfig:
    APP_PREFIX: str = os.getenv("APP_NGINX_PREFIX")


@dataclass
class PostgresConfig:
    """Database connection variables."""

    name: Optional[str] = os.getenv('POSTGRES_DATABASE')
    user: Optional[str] = os.getenv('POSTGRES_USER')
    password: Optional[str] = os.getenv('POSTGRES_PASSWORD')
    port: int = int(os.getenv('POSTGRES_PORT', 5432))
    host: str = 'localhost'

    driver: str = 'asyncpg'
    database_system: str = 'postgresql'

    url_str = URL.create(
            drivername=f'{database_system}+{driver}',
            username=user,
            database=name,
            password=password,
            port=port,
            host=host,
        ).render_as_string(hide_password=False)


@dataclass
class AppConfig:
    """Bot configuration."""

    title = "Video Upload Service"
    description = "Сервис для загрузки видео на IkanamTube"
    version = "0.1 beta"
    root_path = NGINXConfig.APP_PREFIX
    host = '0.0.0.0'
    port = 8005


@dataclass
class Configuration:
    """All in one configuration's class."""

    debug = bool(os.getenv('DEBUG'))
    logging_level = int(os.getenv('LOGGING_LEVEL', logging.INFO))

    postgres = PostgresConfig()
    app = AppConfig()
