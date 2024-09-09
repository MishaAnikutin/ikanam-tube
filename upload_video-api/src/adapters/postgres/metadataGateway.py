from abc import ABC, abstractmethod

from sqlalchemy.ext.asyncio import AsyncSession

from src.entities import Video
from .models import MetadataORM


class MetadataGatewayInterface(ABC):
    @abstractmethod
    async def new(self, video: Video, video_url: str, picture_url: str) -> int:
        ...

    @abstractmethod
    async def delete(self, video_id: int):
        ...


class MetadataGateway(MetadataGatewayInterface):
    def __init__(self, session: AsyncSession):
        self._session = session

    #TODO: получать индекс только что добавленной строчки
    async def new(self, video: Video, video_url: str, picture_url: str) -> int:
        self._session.add(
            MetadataORM(
                channel_id=video.channel_id,
                video_title=video.title,
                video_description=video.description,
                tag=video.tag,
                video_url=video_url,
                picture_url=picture_url,
            )
        )

        await self._session.commit()

    async def delete(self, video_id: int):
        ...
