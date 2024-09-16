import React from 'react';
import { Typography, CardContent } from '@mui/material';

import './VideoDescription.css';

const VideoDescription = ({ title, description }) => {
    return (
        <CardContent className="card-content">
                <Typography variant='subtitle1' className="title">
                    {title.slice(0, 32).toUpperCase() + '...'}
                </Typography>
                <Typography variant='subtitle2' className="description">
                    {description.slice(0, 64) + '...'}
                </Typography>
            </CardContent>
    );
}

export default VideoDescription;
