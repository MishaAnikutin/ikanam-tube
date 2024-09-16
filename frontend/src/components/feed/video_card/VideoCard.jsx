import React from 'react';
import { Card } from '@mui/material';
import VideoPreview from './VideoPreview'
import VideoDescription from './VideoDescription'


import './VideoCard.css';

const VideoCard = ({
    video: { 
      id, 
      channel_id, 
      title, 
      description, 
      tag, 
      likes, 
      dislikes,
      video_path, 
    }
}) => {
    return (
        <Card className="card">
            <VideoPreview id={id} video_path={ video_path } title={title} image_path={''} />
            <VideoDescription title={title} description={description} />
        </Card>
    );
}

export default VideoCard;
