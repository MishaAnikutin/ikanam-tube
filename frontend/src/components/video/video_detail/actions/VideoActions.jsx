import { Stack } from '@mui/material';
import React from 'react';
import LikeDislike from './LikeDislike';
import ShareButton from './ShareButton';

const VideoActions = ({ videoId, initialLikes, initialDislikes }) => {
  return (
    <Stack
      direction={'row'} 
      justifyContent={'flex-end'} 
      spacing={1} 
      sx={{ paddingBottom: 2,}}
    >
      <LikeDislike 
        videoId={videoId} 
        initialLikes={initialLikes} 
        initialDislikes={initialDislikes}
      />
      <ShareButton />
    </Stack>
  );
};

export default VideoActions;