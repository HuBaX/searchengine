import React, { useState } from 'react';
import './App.css';
import Header from './components/header';
import SearchRequest from './components/searchrequest';
import RecipeCard from './components/recipecard';
import {Grid} from "@mui/material"



function App() {
  return (
    <div>
      <Header></Header>
      <SearchRequest></SearchRequest>
      <Grid container spacing={1}>
        <Grid item xs={4}> 
          <RecipeCard></RecipeCard>
        </Grid>
        <Grid item xs={4}>
          <RecipeCard></RecipeCard>
        </Grid>
        <Grid item xs={4}>
          <RecipeCard></RecipeCard>
        </Grid>
      </Grid> 
    </div>
  );
}


export default App;
