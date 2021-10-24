import {IconButton, Box, Paper}  from '@mui/material';
import FilterAltIcon from '@mui/icons-material/FilterAlt';
import SearchIcon from '@mui/icons-material/Search';
import TextField from '@mui/material/TextField';
import Container from '@mui/material/Container';
import React, { useState } from 'react';
import Tags from './autofill';
import DiscreteSliderLabel from './slider';


function SearchRequest() {

    const [state, setState] = useState<any[]>([])
    const [isFilterVisible, setIsFilterVisible] = useState(false)

    const  sendRequest = () => {
        fetch('http://3.218.166.198:8080')
            .then(response => response.json())
            .then(data => setState(data));
    };

    // TODO: Reset state of slider
    const toggleFilterSection = () => {
      setIsFilterVisible(!isFilterVisible)  
    }

  return (
      <Container maxWidth="md">
        <Box sx={{margin: 2}}>
        <Paper>
        <Box sx={{display: "flex"}}>
            <TextField label="Search for a recipe" type="search" size="medium" sx={{flex: "1", margin:2}}/>
            <IconButton onClick={sendRequest} color="primary"  > <SearchIcon/> </IconButton>
            <IconButton onClick={toggleFilterSection} color="primary" > <FilterAltIcon/></IconButton>
        </Box>  
        {isFilterVisible &&
          <div> 
            <DiscreteSliderLabel></DiscreteSliderLabel>
            <Box sx={{display: "flex"}}>
              <Tags filterType="Ingredients" placeholder="e.g Chicken, Fish, ..."></Tags>
              <Tags filterType="Tags" placeholder="e.g high-protein, dessert, ..."></Tags>
            </Box>
          </div> 
        }
        </Paper>
        </Box>
      </Container>
  );
}


export default SearchRequest


