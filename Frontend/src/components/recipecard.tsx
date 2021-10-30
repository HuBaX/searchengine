import * as React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import RecipeDialog from './recipedialog';
import AccessTimeIcon from '@mui/icons-material/AccessTime';
import {Box} from "@mui/material"

interface RecipeCardProps { recipeName: string,  description: string, time: number, recipe_id: number}

function RecipeCard({recipeName, description, time, recipe_id}: RecipeCardProps) {  
    
  return (
      <Card sx={{margin: 2, minWidth: 300, height: 300, display: "flex", flexDirection: "column"}}>
        <CardContent sx={{flex: 1}}>
          <Typography variant="h5">
          {recipeName}
          </Typography>
          <Typography sx={{marginBottom: 1}}>
          {description.length < 350? description : description.substring(0,350) + "..."}
          </Typography>
        </CardContent>
        
        <Box sx={{display: "flex", flexDirection: "row", alignItems: "center"}}>
          <CardActions>
            <RecipeDialog recipe_id={recipe_id}></RecipeDialog>
          </CardActions>
          <Box sx={{flex: 1}}></Box>
          <Box sx={{display: "flex", flexDirection: "row"}}>
            <AccessTimeIcon sx={{marginRight: 0.5}}/>
            <Typography sx={{marginRight: 1}}>
              {time} minutes
            </Typography>
          </Box>
        </Box>
        
      </Card>
  );
}

export default RecipeCard
