from typing import AsyncIterable

from dishka import Provider, Scope, provide
from sqlalchemy.ext.asyncio import AsyncSession

from src.adapters import (
    session_maker,
    MetadataGateway, MetadataGatewayInterface,
    PhotoGateway, PhotoGatewayInterface,
    VideoGateway, VideoGatewayInterface
)

from src.interactors.uploadVideo import UploadVideoInteractor


class AdaptersProvider(Provider):
    scope = Scope.REQUEST

    video_metadata = provide(MetadataGateway, provides=MetadataGatewayInterface)
    preview = provide(PhotoGateway, provides=PhotoGatewayInterface)
    video = provide(VideoGateway, provides=VideoGatewayInterface)

    @provide
    async def session(self) -> AsyncIterable[AsyncSession]:
        async with session_maker() as session:
            yield session


class InteractorProvider(Provider):
    scope = Scope.REQUEST

    product = provide(UploadVideoInteractor)
