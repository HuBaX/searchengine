import React, { useState } from 'react';
import './App.css';
import Header from './components/header';
import SearchRequest from './components/searchrequest';
import RecipeCard from './components/recipecard';
import {Grid} from "@mui/material"
import Scroller from './components/scroller'


function App() {
  return (
    <div>
      <Header></Header>
      <SearchRequest></SearchRequest>
      <Grid container spacing={1}>
        <Scroller></Scroller>
      </Grid> 
    </div>
  );
}


export default App;
