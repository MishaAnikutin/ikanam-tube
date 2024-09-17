import React from 'react'
import { ThumbUp } from '@mui/icons-material';
import './LikeDislike.css';


const Like = ({ likes, liked, handleLike}) => {

  return (
        <button className="like-dislike-button" onClick={handleLike}>
            <ThumbUp className="like-icon" style={{color: liked ? 'white' : '#4d4d4d'}} />
            <span className="like-count">{likes}</span>
        </button>  
    )
}

export default Like