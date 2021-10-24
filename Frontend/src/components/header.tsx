import Typography from '@mui/material/Typography';
import AppBar from '@mui/material/AppBar';

function Header() {
 
    const mystyle = {
        color: "black",
        backgroundColor: "white",
        padding: "10px",
      };

    return (
            <AppBar position="static" sx={{padding: 2, marginBottom: 4}}>
                <Typography variant="h6" color="inherit">
                Search engine for recipes. Search in over 40.000 recipes
                </Typography>
            </AppBar>
    );
}


export default Header