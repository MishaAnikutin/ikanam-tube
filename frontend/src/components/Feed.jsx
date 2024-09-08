import { useState, useEffect } from 'react'
import { Box, Stack, Typography } from '@mui/material';
import { BorderRight } from '@mui/icons-material';

import SideBar from './SideBar';
import Videos from './Videos';

import { getAllMetadata } from '../utils/fetchFromAPI'

const Feed = () => {
    const [selectedCategory, setSelectedCategory] = useState('Главное')
    const [videos, setVideos] = useState([])

    useEffect(() => {
        getAllMetadata()
        .then((data) => {setVideos(data)})
        .catch(error => {
            console.error('Error:', error);
          })
    }, [])

    return (
    <Stack sx={{ flexDirection: { sx: "column", md: "row" } }}>
        
        {/* Боковая панель */}
        <Box sx={{ height: { sx: "auto", md: "92vh" }, borderRight: "1px solid #3d3d3d", px: { sx: 0, md: 2 } }}>
            <SideBar 
                selectedCategory={selectedCategory}
                setSelectedCategory={setSelectedCategory}
                key={selectedCategory}
            />

            <Typography className="copyright" variant="body2" sx={{ mt: 1.5, color: "#fff", }}>
                Copyright © 2024 HIHiGS
            </Typography>
        </ Box>

        {/* Видео */}
        <Box p={2} sx={{overflowY: 'auto', height: '90vh', flex: 2}}>
            <Typography variant='h4' fontWeight='bold' mb={2} sx={{ color: "White"}}>
                {selectedCategory} <span style={{ color: '#FF0000'}}> Видео </span>
            </Typography>

            <Videos videos={videos} />
        </Box>
    </ Stack>
    )
}

export default Feed