import React from 'react';
import './Category.css'; 

const Category = ({ selectedCategory, setSelectedCategory, category }) => (
    <button 
        key={category.name}
        className={`category-btn ${category.name === selectedCategory ? 'active' : ''}`} 
        onClick={() => setSelectedCategory(category.name)}
    >
        <span style={{color: "white"}}>
            {category.icon}
        </span>
        <span style={{marginLeft: '25px'}}>{category.name}</span>
    </button>

);

export default Category;
