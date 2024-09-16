import { useState, useEffect } from 'react';
import { Stack } from '@mui/material';
import SideBar from './side_bar/SideBar';
import VideosFeed from './VideosFeed';
import { getAllMetadata } from '../../utils/fetchFromAPI';

import './Feed.css';

const Feed = () => {
    const [selectedCategory, setSelectedCategory] = useState('Главное');
    const [videos, setVideos] = useState([]);

    useEffect(() => {
        getAllMetadata()
            .then((data) => { setVideos(data); })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    return (
        <Stack className="feed-container" sx={{ flexDirection: { sx: "column", md: "row" } }}>
            <SideBar
                selectedCategory={selectedCategory}
                setSelectedCategory={setSelectedCategory}
                key={selectedCategory}
            />
            <VideosFeed
                videos={videos}
                selectedCategory={selectedCategory}
            />
        </Stack>
    );
};

export default Feed;
