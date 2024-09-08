import os
import ffmpeg

import aiofiles
from abc import ABC, abstractmethod


class VideoGatewayInterface(ABC):
    @abstractmethod
    async def add(self, video_file, title: str) -> str:
        ...

    @abstractmethod
    async def convert(self, title: str):
        ...

    @abstractmethod
    async def remove(self, title: str):
        ...


class VideoGateway(VideoGatewayInterface):
    _video_path = '../static/videos/'

    async def add(self, video_file, title: str) -> str:
        async with aiofiles.open(self._video_path + title + '.mp4', 'wb') as f:
            while content := await video_file.read(1024):
                await f.write(content)

        return f'/static/videos/{title}.m3u8'

    async def remove(self, title: str):
        path = self._video_path + title + '.mp4'

        if os.path.exists(path):
            os.remove(path)

    def convert(self, title: str):
        try:
            return ffmpeg \
                .input(self._video_path + title + '.mp4') \
                .output(
                    self._video_path + title + '.m3u8',
                    format='hls',
                    hls_time=10,
                    hls_segment_type='mpegts') \
                .run()

        except ffmpeg.Error as e:
            return False
