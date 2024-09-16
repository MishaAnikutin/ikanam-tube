import React from 'react'
import { Box, Stack, Typography } from '@mui/material';
import VideoCard from './video_card/VideoCard'

const VideosFeed = ({ videos, selectedCategory }) => {
  return (
    <Box className="videos-container">
      <Typography variant='h4' fontWeight='bold' mb={2} className="videos-title">
        {selectedCategory}
      </Typography>
      <Stack direction="row" flexWrap="wrap" justifyContent="start" gap={2}>
        {videos.map((video) => <VideoCard key={video.id} video={video}/>)}
      </Stack>
    </Box>
  );
};

export default VideosFeed