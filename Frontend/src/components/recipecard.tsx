import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import RecipeDialog from './recipedialog';



function RecipeCard() {

    const showFullRecipe = () => {
        
    }
  
    
  return (
        <Card sx={{margin: 2}}>
        <CardContent>
            <Typography variant="h5">
            Recipe name
            </Typography>
            <Typography>
            Description
            </Typography>
            <Typography>
            Time
            </Typography>
        </CardContent>
        <CardActions>
          <RecipeDialog></RecipeDialog>
        </CardActions>
        </Card>
  );
}

export default RecipeCard

// sx={{ minWidth: 275 , width: 500,  justifyContent: "center", alignItems: "center", margin:1}}