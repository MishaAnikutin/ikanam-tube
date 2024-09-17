import React, { useState } from 'react';
import { Share } from '@mui/icons-material';
import InfoNotification from '../../../notifications/InfoNotification'; // Импортируем Notification
import './ShareButton.css'; // Подключите файл стилей

const ShareButton = () => {
  const [showNotification, setShowNotification] = useState(false);

  const handleShareClick = () => {
    const urlToShare = window.location.href;
    navigator.clipboard.writeText(urlToShare)
      .then(() => {
        setShowNotification(true); // Показываем уведомление
      })
      .catch(err => {
        console.error('Ошибка при копировании:', err);
      });
  };

  const handleCloseNotification = () => {
    setShowNotification(false); // Закрываем уведомление
  };

  return (
    <div>
      <button className="share-button" onClick={handleShareClick}>
        <Share className="share-icon" style={{ color: 'white' }} />
        <span className="share-text">Поделиться</span>
      </button>
      {showNotification && (
        <InfoNotification message="Ссылка скопирована" onClose={handleCloseNotification} />
      )}
    </div>
  );
};

export default ShareButton;
