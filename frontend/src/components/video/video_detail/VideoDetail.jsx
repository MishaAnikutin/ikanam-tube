import React from 'react'
import ReactPlayer from 'react-player';
import { Box, Typography } from '@mui/material';
import VideoActions from './actions/VideoActions';
import Description from './Description';

const VideoDetail = ({ videoDetail }) => {
  return (
    <Box flex={1} p={2}>
      <Box sx={{ width: '100%', position: 'sticky', top: '86px' }}>
        <ReactPlayer 
          url={ 'https://www.youtube.com/watch?v=VC-cTUMonsE&t=9s' } 
          className={'react-player'}
          controls
        />
        <Typography color='white' variant='h5' fontWeight={'bold'} p={1}>
          {videoDetail.title}
        </Typography>
s
        <VideoActions videoId={1} initialLikes={123} initialDislikes={1} /> 
        <Description description={videoDetail.description}/>
      </Box>
    </Box>
  )
}

export default VideoDetail