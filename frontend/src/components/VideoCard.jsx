import React from 'react'
import { Link } from 'react-router-dom'
import { Typography, Card, CardContent, CardMedia,  } from '@mui/material'
import { CheckCircle } from '@mui/icons-material'

const VideoCard = ({
    video: { 
      id, 
      channel_id, 
      title, 
      description, 
      tag, 
      likes, 
      dislikes, 
      video_path 
    }
   }) => {
    return (
    <Card sx={{ width: {md: '350px', xs: '100%'}, borderRadius: '10px'}}>
      <Link to={`video/${video_path}`}>
        <CardMedia 
          alt={title}
          sx={{ width: 358, height: 180 }}
        />
        <CardContent sx={{backgroundColor: '#1e1e1e', height: '50px'}}>
          <Link to={`video/${video_path}`}>
            <Typography variant='subtitle1' fontWeight='medium' color='white'>
              {title.slice(0, 32).toUpperCase() + '...'}
            </Typography>
            <Typography variant='subtitle1' fontSize={12} color='gray'>
              {description.slice(0, 64) + '...'}
            </Typography>
          </Link>
        </CardContent>
      </Link>
    </Card>
  )
}

export default VideoCard