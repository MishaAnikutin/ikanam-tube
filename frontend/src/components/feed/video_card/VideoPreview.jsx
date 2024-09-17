import React from 'react';
import { Link } from 'react-router-dom';
import { CardMedia } from '@mui/material';

import './VideoPreview.css';

const VideoPreview = ({ id, title, video_path, image_path }) => {
    return (
        <Link to={`/video/${id}`} className="link">
            <CardMedia 
                component="img"
                alt={title}
                className="card-media"
                image="https://cdn.culture.ru/images/7b2b6db0-dbd0-59fc-bfec-052d74040332" 
            />
        </Link>
    );
}

export default VideoPreview;
