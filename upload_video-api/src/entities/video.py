from pydantic import BaseModel

from .tagsEnum import TagsEnum


class Video(BaseModel):
    title: str
    description: str
    tag: TagsEnum
    channel_id: int
