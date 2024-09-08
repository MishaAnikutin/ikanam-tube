from fastapi import APIRouter, UploadFile, HTTPException
from dishka.integrations.fastapi import FromDishka, inject

from src.entities import Video, UploadException
from src.interactors import UploadVideoI


upload_video_router = APIRouter(prefix="/video")


@upload_video_router.post("/upload")
@inject
async def upload(
        video_metadata: Video,
        video_file: UploadFile,
        picture_file: UploadFile,
        upload_interactor: FromDishka[UploadVideoI]
) -> dict[str, str]:
    try:
        await upload_interactor.upload(
            video_metadata=video_metadata,
            video_file=video_file,
            picture_file=picture_file
        )

    except UploadException as exc:
        raise HTTPException(status_code=400, detail=exc)

    else:
        return {"status": "ok"}
