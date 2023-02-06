import {useEffect, useState} from "react";

export const Home: React.FC<{  }> = props => {
    const [count, setCount] = useState(0);

    useEffect(() => {
        fetch("http://localhost:8080")
            .then(response => response.json())
            .then(data => {
                setCount(data.count);
            }
        );
        console.log("Test")
    }, []);

    const handleIncrementClick = () => {
        setCount(count + 1);
    }

    const handleDecrementClick = () => {
        if (count > 0) {
            setCount(count - 1);
        }
    }

    return (
        <div>
            <h1>Home</h1>
            <p>Count: {count}</p>
            <button onClick={handleIncrementClick}>Increment</button>
            <button onClick={handleDecrementClick}>Decrement</button>
        </div>
    );
};