import { Stack } from '@mui/material';
import SearchBar from './SearchBar';
import Logo from './Logo';
import './Header.css'; 

const Header = () => (
    <Stack className="header" direction="row" alignItems="center" spacing={2}>
        <Logo />
        <SearchBar className='search' />
    </Stack>
);

export default Header;
