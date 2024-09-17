import React from 'react'
import { ThumbDown } from '@mui/icons-material';
import './LikeDislike.css';


const Dislike = ({ dislikes, disliked, handleDislike }) => {
  return (
        <button className="like-dislike-button" onClick={handleDislike}>
            <ThumbDown className="dislike-icon" style={{color: disliked ? 'white' : '#4d4d4d'}} />
            <span className="dislike-count">{dislikes}</span>
        </button>  
    )
}

export default Dislike