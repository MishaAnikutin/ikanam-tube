import React, { useEffect, useState } from 'react'
import { Box, Stack } from '@mui/material';
import { useParams } from 'react-router-dom';
import { getVideoMetadataByID } from '../../utils/fetchFromAPI';

import VideoDetail from './video_detail/VideoDetail';
import VideoRecomendations from './VideoRecomendations';

const VideoPage = () => {
  const [ videoDetail, setVideoDetail ] = useState({});
  const { id } = useParams();

  useEffect(() => {
    getVideoMetadataByID(id)
    .then((data) => { setVideoDetail(data); })
    .catch(error => {
        console.error('Error:', error);
    });
  }, [id])

  return (
    <Box minHeight={'95vh'}>
      <Stack  direction={{ xs: 'column', md: 'row' }}>
        <VideoDetail videoDetail={videoDetail}/>
        <VideoRecomendations />
      </Stack>
    </Box>
  )
}

export default VideoPage
