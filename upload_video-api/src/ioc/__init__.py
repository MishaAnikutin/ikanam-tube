from dishka import make_async_container
from .provider import AdaptersProvider, InteractorProvider


container = make_async_container(AdaptersProvider(), InteractorProvider())

__all__ = ('container', )
