import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Typography } from '@mui/material';


interface Recipe {n_steps: number, n_ingredients: number, minutes: number, description: string, steps: string[], tags: string[], nutrition: number[], name: string, ingredients: string[]}
interface Recipe_id {recipe_id: number}

function RecipeDialog({recipe_id}: Recipe_id) {
  const [open, setOpen] = React.useState(false);
  const [recipeData, setRecipeData] = React.useState<Recipe>({n_steps: 0, n_ingredients: 0, minutes: 0, description: "Description ", steps:["Steps..."], tags: ["Tags..."], nutrition: [], name: "Name", ingredients: ["Ingredients..."] })

  const handleClickOpen = async() => {
    setOpen(true);

    const response = await fetch('http://3.218.166.198:8080/recipe_search', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({Id: recipe_id})
    })
    const data = await response.json() as Recipe
    setRecipeData(data)
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <Button variant="outlined" onClick={handleClickOpen}>
        Show recipe
      </Button>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
        fullWidth={true}
      >
        <DialogTitle id="recipe-card-recipe-name">
        <Typography variant="h4" fontWeight={600}>{recipeData.name}</Typography>
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="recipe-card-recipe-description">
            <Typography variant="h6" fontWeight={600} color='black' sx={{marginBottom:1, marginTop:1}}>Description:</Typography>
            <Typography variant="body1" color='black' >{recipeData.description} </Typography>
            <Typography variant="h6" fontWeight={600} color='black' sx={{marginBottom:1, marginTop:1}}>Estimated Time: </Typography>
            <Typography variant="body1" color='black' >{recipeData.minutes} minutes</Typography>
            <Typography variant="h6" fontWeight={600} color='black' sx={{marginBottom:1, marginTop:1}}>Ingredients: </Typography>
            <Typography variant="body1" color='black' >{recipeData.ingredients != null ? recipeData.ingredients.map((ingredient => {return <p>{ingredient}</p>})): <p>Ingredinets not available</p>}</Typography>
            <Typography variant="h6" fontWeight={600} color='black' sx={{marginBottom:1, marginTop:1}}>Steps: </Typography>
            <Typography variant="body1" color='black' sx={{paddingLeft:2}}><ol>{recipeData.steps != null ? recipeData.steps.map((step => {return <li>{step}</li>})): <p>No steps available</p>}</ol></Typography>
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} sx={{}}>Close</Button>
        </DialogActions>      
      </Dialog>
    </div>
  );
}

export default RecipeDialog