from abc import ABC, abstractmethod

from src.entities import Video, UploadException
from src.adapters import MetadataGatewayInterface, VideoGatewayInterface, PhotoGatewayInterface


class UploadVideoI(ABC):
    @abstractmethod
    async def upload(self, video_metadata: Video, video_file, picture_file):
        ...
    

class UploadVideoInteractor(UploadVideoI):
    __slots__ = (
        '_metadata_gateway',
        '_video_gateway',
        '_photo_gateway',
    )

    def __init__(
            self,
            metadata_gateway: MetadataGatewayInterface,
            video_gateway: VideoGatewayInterface,
            photo_gateway: PhotoGatewayInterface
    ):
        self._metadata_gateway = metadata_gateway
        self._video_gateway = video_gateway
        self._photo_gateway = photo_gateway


    async def upload(self, video: Video, video_file, photo_file):

        try:
            # save video file in storage
            video_url = await self._video_gateway.add(video_file=video_file, title=video.title)

            # convert mp4 to m3u8
            await self._video_gateway.convert(title=video.title)

            # save video preview
            photo_url = await self._photo_gateway.add(photo_file=photo_file, title=video.title)

            # insert metadata in database
            await self._metadata_gateway.new(video_metadata=video, video_url=video_url, photo_url=photo_url)

        except Exception as exc:
            # FIXME:
            await self._video_gateway.remove(title=video.title)
            await self._photo_gateway.remove(title=video.title)
            raise UploadException(exc)
