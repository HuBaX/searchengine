import {IconButton, Box, Paper}  from '@mui/material';
import FilterAltIcon from '@mui/icons-material/FilterAlt';
import SearchIcon from '@mui/icons-material/Search';
import TextField from '@mui/material/TextField';
import Container from '@mui/material/Container';
import React, { useState } from 'react';
import Tags from './autofill';
import DiscreteSliderLabel from './slider';
import RecipeCard from './recipecard'
import InfiniteScroll from "react-infinite-scroll-component";
import Autocomplete from '@mui/material/Autocomplete';

interface RecipePreview {
  id: number 
	minutes: number  
	name: string 
	description: string 
}

interface LoadedRecipes {
  recipes: RecipePreview[]
  alreadyLoaded: number
}

function SearchRequest() {

  const [recipes, setRecipes] = useState<RecipePreview[]>([])
  const [isFilterVisible, setIsFilterVisible] = useState(false)
  const [tags, setTags] = useState<string[]>([])
  const [ingredients, setIngredients] = useState<string[]>([])
  const [alreadyLoaded, setAlreadyLoaded] = useState(0)
  const [name, setName] = useState("")
  const [nameOptions, setNameOptions] = useState<string[]>([])
  const [minutes, setMinutes] = useState(150)
  const [hasMore, setHasMore] = useState(true);

  const  sendSearchRequest = async() => {
    var response
    if(!isFilterVisible) {
      response = await fetch('http://3.218.166.198:8080/name_search', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({from: alreadyLoaded, name: name})
      })
    } else {
      response = await fetch('http://3.218.166.198:8080/filter_search', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({from: alreadyLoaded, minutes: minutes, name: name, tags: tags, ingredients: ingredients})
      })
    }
    const data = await response.json() as LoadedRecipes
    setRecipes([...recipes, ...data.recipes])
    setAlreadyLoaded(alreadyLoaded + data.alreadyLoaded)
    if(data.recipes.length === 0) {
      setHasMore(false)
    } else {
      setHasMore(true)
    }
  };

  const sendNameAutocompleteReq = async(prefix: string) => {
    const response = await fetch('http://3.218.166.198:8080/name_search', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({prefix: prefix})
    })
    const data = await response.json() as string[]
    setNameOptions(data)
  }

  const toggleFilterSection = () => {
    setIsFilterVisible(!isFilterVisible)  
    if(!isFilterVisible) {
      setTags([])
      setIngredients([])
      setMinutes(150)
    }
  }

  return (
      <Container maxWidth="md">
        <Box sx={{margin: 2}}>
          <Paper>
            <Box sx={{display: "flex"}}>
            <Autocomplete
              freeSolo
              id="search-field"
              sx={{flex: "1", margin:1}}
              disableClearable
              options={nameOptions}
              renderInput={(params) => (
              <TextField
                {...params}
                label="Search for a recipe"
                sx={{margin:1}}
                value={name}
                size="medium"
                onChange={(event) => {setName(event.target.value); sendNameAutocompleteReq(event.target.value);}}
                InputProps={{
                  ...params.InputProps,
                  type: 'search',
                }}
              />
            )}
           />
                <IconButton onClick={sendSearchRequest} color="primary"  > <SearchIcon/> </IconButton>
                <IconButton onClick={toggleFilterSection} color="primary" > <FilterAltIcon/></IconButton>
            </Box>  
            {isFilterVisible &&
              <div> 
                <DiscreteSliderLabel setSliderValue={setMinutes}></DiscreteSliderLabel>
                <Box sx={{display: "flex"}}>
                  <Tags filterType="Ingredients" placeholder="e.g Chicken, Fish, ..." setTags={setIngredients} autocompleteRoute="ingredients_autocomplete"></Tags>
                  <Tags filterType="Tags" placeholder="e.g high-protein, dessert, ..." setTags={setTags} autocompleteRoute="tags_autocomplete"></Tags>
                </Box>
              </div> 
            }
          </Paper>
        </Box>
        <InfiniteScroll
          dataLength={15} //This is important field to render the next data
          next={sendSearchRequest}
          hasMore={hasMore}
          loader={<div></div>}
          endMessage={
              <p style={{ textAlign: 'center' }}>
              <b>Yay! You have seen it all</b>
              </p>}>
          <div className="container">
            <div className="row m-2">
              {recipes.map((recipe) => {
                return <RecipeCard recipeName={recipe.name} description={recipe.description} time={recipe.minutes} recipe_id={recipe.id}/>;
              })}
            </div>
          </div>
        </InfiniteScroll>
      </Container>
  );
}


export default SearchRequest


