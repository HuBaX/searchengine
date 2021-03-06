import Typography from '@mui/material/Typography';
import AppBar from '@mui/material/AppBar';

function Header() {

    return (
            <AppBar position="static" sx={{padding: 2, marginBottom: 4}}>
                <Typography variant="h6" color="inherit" textAlign='center'>
                Search engine for recipes. Search in over 40.000 recipes
                </Typography>
            </AppBar>
    );
}


export default Header