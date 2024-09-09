import os
import aiofiles
from abc import ABC, abstractmethod


class PhotoGatewayInterface(ABC):
    @abstractmethod
    async def add(self, picture_file, title: str) -> str:
        ...

    @abstractmethod
    async def remove(self, title: str):
        ...


class PhotoGateway(PhotoGatewayInterface):
    _photo_path = '../static/pictures/'
    _format = '.jpg'

    async def add(self, picture_file, title: str) -> str:
        path = self._photo_path + title + self._format

        async with aiofiles.open(path, 'wb') as f:
            await f.write(await picture_file.read())

        return f'/static/pictures/{title}{self._format}'

    async def remove(self, title: str):
        path = self._photo_path + title + self._format

        if os.path.exists(path):
            os.remove(path)
