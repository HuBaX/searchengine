import { Button } from '@mui/material';
import React, { useState } from 'react';


function Search() {
    
    const [state, setState] = useState<any[]>([]);

    const  sendRequest = () => {
        fetch('http://3.218.166.198:8080')
            .then(response => response.json())
            .then(data => setState(data));
    };
 
    return (
        <div>
            <Button onClick={sendRequest}>Test</Button>
            <p>Data: {JSON.stringify(state)}</p>
        </div>
    )
}



export default Search 