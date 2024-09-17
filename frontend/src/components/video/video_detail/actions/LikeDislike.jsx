import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Like from './Like'
import Dislike from './Dislike';

import './LikeDislike.css';

const LikeDislike = ({ videoId, initialLikes, initialDislikes }) => {
  const [likes, setLikes] = useState(initialLikes);
  const [dislikes, setDislikes] = useState(initialDislikes);
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);

  const handleLike = async () => {
    // TODO: ходить в микросервис 

    if (liked) {
      setLikes(likes - 1);
      setLiked(false);
    } else {
      setLikes(likes + 1);
      setLiked(true);
      if (disliked) {
        setDislikes(dislikes - 1);
        setDisliked(false);
      }
    }
  };

  const handleDislike = async () => {  
    if (disliked) {
      setDislikes(dislikes - 1);
      setDisliked(false);
    } else {
      setDislikes(dislikes + 1);
      setDisliked(true);
      if (liked) {
        setLikes(likes - 1);
        setLiked(false);
      }
    }
  };

  return (
    <div className="like-dislike-container">
      <Like likes={likes} liked={liked} handleLike={handleLike} />
      <div className="divider" />
      <Dislike dislikes={dislikes} disliked={disliked} handleDislike={handleDislike} />
    </div>
  );
};

LikeDislike.propTypes = {
  videoId: PropTypes.string.isRequired,
  initialLikes: PropTypes.number.isRequired,
  initialDislikes: PropTypes.number.isRequired,
};

export default LikeDislike;
