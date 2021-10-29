import * as React from 'react';
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';
import {Container} from '@mui/material';
import { useState } from "react";

interface TagProps { filterType: string,  placeholder: string, setTags: React.Dispatch<React.SetStateAction<string[]>>, autocompleteRoute:string}



function Tags({filterType,placeholder,setTags,autocompleteRoute}: TagProps) {

  const [options, setOptions] = useState<string[]>([]);

  const requestOptions = async (prefix: string) => {
    const response = await fetch("http://3.218.166.198:8080/" + autocompleteRoute, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({prefix: prefix})
    })
    const data = await response.json() as string[];
    setOptions(data);
  }

  return (
    <Container>
      <Autocomplete
        multiple
        id="tags-outlined"
        options={options}
        getOptionLabel={ (item) => {return item}}
        filterSelectedOptions
        onInputChange={(event, value, reason) => {
          requestOptions(value)
        }}
        onChange={(event, value) => setTags(value)}
        renderInput={(params) => (
          <TextField
            {...params}
            label={filterType}
            placeholder={placeholder}
            sx={{marginBottom: 2}}
          />
        )}
      />
    </Container>
  );
}


export default  Tags
