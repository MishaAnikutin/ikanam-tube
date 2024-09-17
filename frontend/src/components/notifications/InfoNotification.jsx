import React, { useEffect } from 'react';
import './InfoNotification.css';

const InfoNotification = ({ message, onClose }) => {
  useEffect(() => {
    const timer = setTimeout(onClose, 1500);
    return () => clearTimeout(timer);
  }, [onClose]);

  return (
    <div className="notification">{message}</div>
  );
};

export default InfoNotification;
