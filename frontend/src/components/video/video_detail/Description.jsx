import React from 'react'
import { Box, Typography } from '@mui/material';

const Description = ({ description }) => {
  return (
    <Box bgcolor={'#1a1a1a'} borderRadius={'10px'} top={2}>
        <Typography color='white' fontSize={18} p={2}>
            {description}
        </Typography>
    </Box>
  )
}

export default Description