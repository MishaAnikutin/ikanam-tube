import React from 'react'

import { Stack } from '@mui/material';
import { categories } from '../utils/constants';

const SideBar = ({ selectedCategory, setSelectedCategory }) => (
    <Stack direction="row" sx={{ overflowY: "auto", height: { sx: "auto", md: '95%' }, borderRadius: '5px', flexDirection: {md: "column"}}}>
                
        {categories.map((category) => (
            <button 
                className='category-btn' 
                onClick={() => setSelectedCategory(category.name)}
                style={{
                    margin: '5px',
                    borderRadius: '10px',
                    background: category.name === selectedCategory && '#272727',
                    color: 'white'
                }}
            >
                <span style={{color: "white"}}>
                    {category.icon}
                </span>
                <span style={{marginLeft: '25px'}}>{category.name}</span>
            </button>
        ))}
    </Stack>
)

export default SideBar