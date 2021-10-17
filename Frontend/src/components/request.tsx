import { Button } from '@mui/material';
import React, { useState } from 'react';


function Search() {
    
    const [state, setState] = useState([]);

    const  sendRequest = () => {
        fetch('3.218.166.198:8080')
            .then(response => response.json())
            .then(data => setState(data));
    };
 
    return (
        <div>
            <Button>Test</Button>
            <p>Data: {state}</p>
        </div>
    )
}



export default Search 