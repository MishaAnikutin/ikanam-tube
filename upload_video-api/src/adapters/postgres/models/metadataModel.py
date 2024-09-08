from datetime import date

from sqlalchemy.orm import Mapped, mapped_column
from sqlalchemy import Integer, VARCHAR, TEXT, DateTime

from .base import BaseTable


class MetadataORM(BaseTable):
    __tablename__ = 'video_metadata'

    video_id: Mapped[int] = mapped_column(Integer, unique=True, nullable=False, primary_key=True)
    channel_id: Mapped[int] = mapped_column(Integer, unique=True, nullable=False)
    video_title: Mapped[str] = mapped_column(VARCHAR(255), nullable=False)
    video_description: Mapped[str] = mapped_column(TEXT, nullable=True)
    video_url: Mapped[str] = mapped_column(VARCHAR(255), nullable=False)
    picture_url: Mapped[str] = mapped_column(VARCHAR(255), nullable=True)  # FIXME
    upload_date: Mapped[date] = mapped_column(DateTime, default=date.today())
    tag: Mapped[str] = mapped_column(VARCHAR(255), nullable=False)
    likes: Mapped[int] = mapped_column(Integer, default=0)
    dislikes: Mapped[int] = mapped_column(Integer, default=0)

    def __str__(self) -> str:
        return f'<Metadata:{self.video_id}>'
