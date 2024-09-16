import React, { useState } from 'react'; // Импортируем useState
import { Paper, IconButton } from '@mui/material';
import { Search } from '@mui/icons-material';
import './SearchBar.css';

const SearchBar = () => {
    const [text, setText] = useState('');

    const handleChange = (event) => {
        setText(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault(); // Предотвращаем перезагрузку страницы
        console.log(text);      // Здесь можно добавить логику, например, поиск
    };    


    return (
        <Paper component="form" onSubmit={handleSubmit} className="search-bar-container">
            <input
                className='search-bar'
                placeholder='Поиск...'
                value={text} 
                onChange={handleChange} 
            />
            <IconButton type='submit' className='icon-button'>
                <Search />
            </IconButton>
        </Paper>
    );
}

export default SearchBar;
