import * as React from 'react';
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';
import {Container} from '@mui/material';
import { useState } from "react";

interface TagProps { filterType: string,  placeholder: string}



function Tags({filterType,placeholder}: TagProps) {

  const [items, setItems] = useState<string[]>([]);


  return (
    <Container>
      <Autocomplete
        multiple
        id="tags-outlined"
        options={items}
        getOptionLabel={ (item) => {return item}}
        filterSelectedOptions
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
