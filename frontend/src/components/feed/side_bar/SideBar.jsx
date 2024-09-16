import React from 'react';
import { Box, Stack, Typography } from '@mui/material';
import { categories } from '../../../utils/constants';
import Category from './Category';
import './SideBar.css'; 


const SideBar = ({ selectedCategory, setSelectedCategory }) => (
    <Box className="sidebar">
        <Stack direction={{xs: 'row', md: 'column'}} className="category-stack">
            {categories.map((category) => (
                <Category className="category"
                    key={category.name}
                    category={category} 
                    selectedCategory={selectedCategory} 
                    setSelectedCategory={setSelectedCategory}
                />
            ))}
        </Stack>
        <Typography className="copyright" variant="body2">
            Copyright © 2024 HIHiGS
        </Typography>
    </Box>
);

export default SideBar;
