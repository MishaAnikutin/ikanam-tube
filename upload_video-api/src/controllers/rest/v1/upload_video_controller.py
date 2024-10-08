from fastapi import APIRouter, UploadFile, HTTPException
from dishka.integrations.fastapi import FromDishka, inject

from src.entities import Video, UploadException, VideoResponse, TagsEnum
from src.interactors import UploadVideoI


upload_video_router = APIRouter(prefix="/video")


@upload_video_router.post("/upload")
@inject
async def upload(
        title: str,
        description: str,
        tag: TagsEnum,
        channel_id: int,
        video_file: UploadFile,
        picture_file: UploadFile,
        upload_interactor: FromDishka[UploadVideoI]
):

    try:
        await upload_interactor.upload(
            video=Video(
                title=title,
                description=description,
                tag=tag,
                channel_id=channel_id
            ),
            video_file=video_file,
            picture_file=picture_file
        )

    except UploadException as exc:
        raise HTTPException(status_code=400, detail=exc)

    else:
        return {"result": 'Success!'}
