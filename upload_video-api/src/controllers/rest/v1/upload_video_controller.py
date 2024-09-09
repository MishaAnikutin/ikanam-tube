from fastapi import APIRouter, UploadFile, HTTPException
from dishka.integrations.fastapi import FromDishka, inject

from src.entities import Video, UploadException, VideoResponse, TagsEnum
from src.interactors import UploadVideoI


upload_video_router = APIRouter(prefix="/video")


@upload_video_router.post("/upload")
@inject
async def upload(
        # video_metadata: Video,
        video_file: UploadFile,
        picture_file: UploadFile,
        upload_interactor: FromDishka[UploadVideoI]
):

    video = Video(
        title='Математический анализ. Введение',
        description='Математический анализ, 1 курс, 1 семестр, 1 задание',
        tag=TagsEnum.matan,
        channel_id=1
    )

    try:
        await upload_interactor.upload(
            video=video,
            video_file=video_file,
            picture_file=picture_file
        )

    except UploadException as exc:
        raise HTTPException(status_code=400, detail=exc)

    else:
        return {"result": 'Success!'}
