from fastapi import APIRouter
from .upload_video_controller import upload_video_router


router = APIRouter()
router.include_router(upload_video_router)
