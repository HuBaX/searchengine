import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import RecipeDialog from './recipedialog';
import {Grid} from "@mui/material"

interface RecipeCardProps { recipeName: string,  description: string, time: number}

function RecipeCard({recipeName, description, time}: RecipeCardProps) {

    const showFullRecipe = () => {
        
    }
  
    
  return (
      <Grid item xs={4}> 
        <Card sx={{margin: 2}}>
        <CardContent>
            <Typography variant="h5">
            {recipeName}
            </Typography>
            <Typography>
            {description}
            </Typography>
            <Typography>
            {time}
            </Typography>
        </CardContent>
        <CardActions>
          <RecipeDialog></RecipeDialog>
        </CardActions>
        </Card>
      </Grid>
  );
}

export default RecipeCard
