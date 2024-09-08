import os
import aiofiles
from abc import ABC, abstractmethod


class PhotoGatewayInterface(ABC):
    @abstractmethod
    async def add(self, photo_file, title: str) -> str:
        ...

    @abstractmethod
    async def remove(self, title: str):
        ...


class PhotoGateway(PhotoGatewayInterface):
    _photo_path = '../static/pictures/'

    async def add(self, photo_file, title: str) -> str:
        path = self._photo_path + title + '.mp4'

        async with aiofiles.open(path, 'wb') as f:
            await f.write(photo_file.read())

        return f'/static/pictures/{title}.jpeg'

    async def remove(self, title: str):
        path = self._photo_path + title + '.jpeg'

        if os.path.exists(path):
            os.remove(path)
