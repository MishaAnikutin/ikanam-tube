import React from 'react'

import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { Box } from '@mui/material'

import { Navbar, Feed, SearchFeed, VideoDetail, ChannelDetail } from './components';

const App = () => {
  return (
    <BrowserRouter>
      <Box sx={{backgroundColor: '#0F0F0F'}}>
        <Navbar />
        <Routes>
          <Route path='/' exact element={ <Feed /> } />
          <Route path='/video/:id' exact element={ <VideoDetail /> } />
          <Route path='/channel/:id' exact element={ <ChannelDetail /> } />
          <Route path='/search/:searchTerm' exact element={ <SearchFeed /> } />
        </Routes>
      </Box>
    </BrowserRouter>
  )
}

export default App