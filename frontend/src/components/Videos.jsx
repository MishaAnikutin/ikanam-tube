import React from 'react'
import { Stack, Box } from '@mui/material'
import VideoCard from './VideoCard'

const Videos = ({ videos }) => {
  return (
      <Stack direction="row" flexWrap="wrap" justifyContent="start" gap={2}>
        {videos.map((video) => (
              <Box key={video.id}>
                {<VideoCard video={video}/>}
                {/* {<ChannelCard channel_id={video}/>} */}

              </Box>
          ))}
      </Stack>
  );
};

export default Videos