import { Link } from 'react-router-dom';
import { logo } from '../../utils/constants';
import './Logo.css'; 

const Logo = () => (
    <Link to="/" className="logo-container">
        <img src={logo} alt="logo" height={50} />
    </Link>
);

export default Logo;
