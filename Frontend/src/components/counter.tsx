import { Button } from '@mui/material';
import React, { useState } from 'react';


function Counter() {
    const [count, setCount] = useState(0);
    const  fun = () => {
        setCount(count + 1)
    };
      
    return (
        <div>
            <button onClick={fun}> </button>
            <p>{count}</p>
            <Button>Test</Button>
        </div>
    )
}

export default Counter