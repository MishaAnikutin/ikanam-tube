import React from 'react'
import { useState, useEffect } from 'react';
import { Stack } from '@mui/material';
import { getAllMetadata } from '../../utils/fetchFromAPI';

import VideoCard from '../feed/video_card/VideoCard';


const VideoRecomendations = () => {
    const [videos, setVideos] = useState([]);

    useEffect(() => {
        getAllMetadata()
            .then((data) => { setVideos(data); })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    return (
        <Stack direction="column" flexWrap="wrap" justifyContent="start" gap={2}>
            {videos.map((video) => <VideoCard key={video.id} video={video}/>)}
        </Stack>
    )
}

export default VideoRecomendations