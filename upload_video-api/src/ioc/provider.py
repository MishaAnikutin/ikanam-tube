from typing import AsyncIterable

from dishka import Provider, Scope, provide
from sqlalchemy.ext.asyncio import AsyncSession, async_sessionmaker

from src.adapters import (
    postgres_session_maker,
    MetadataGateway, MetadataGatewayInterface,
    PhotoGateway, PhotoGatewayInterface,
    VideoGateway, VideoGatewayInterface
)

from src.interactors.uploadVideo import UploadVideoInteractor, UploadVideoI


class AdaptersProvider(Provider):
    scope = Scope.REQUEST

    video_metadata = provide(MetadataGateway, provides=MetadataGatewayInterface)
    preview = provide(PhotoGateway, provides=PhotoGatewayInterface)
    video = provide(VideoGateway, provides=VideoGatewayInterface)

    @provide
    async def get_session_maker(self) -> async_sessionmaker[AsyncSession]:
        return postgres_session_maker()

    @provide(scope=Scope.REQUEST)
    async def session(self, session_maker: async_sessionmaker[AsyncSession]) -> AsyncIterable[AsyncSession]:
        async with session_maker() as session:
            yield session


class InteractorProvider(Provider):
    scope = Scope.REQUEST

    product = provide(UploadVideoInteractor, provides=UploadVideoI)
