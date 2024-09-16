import React, { useEffect, useState } from 'react'
import { Box, Stack, Typography } from '@mui/material';
import ReactPlayer from 'react-player';
import { useParams } from 'react-router-dom';
import { getVideoMetadataByID } from '../../utils/fetchFromAPI';

const VideoDetail = () => {
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
      <Stack direction={{ xs: 'column', md: 'row' }}>
        <Box flex={1}>
          <Box sx={{ width: '100%', position: 'sticky', top: '86px' }}>
            <ReactPlayer 
              url={ 'https://www.youtube.com/watch?v=FHTbsZEJspU&' } 
              className={'react-player'}
              controls
            />
            <Typography color='white' variant='h5' fontWeight={'bold'} p={2}>
              {videoDetail.title}
            </Typography>
            
          </Box>
        </Box>
      </Stack>
    </Box>
  )
}

export default VideoDetail
