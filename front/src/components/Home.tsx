import React from "react";
import {useEffect, useState} from "react";

export const Home: React.FC<{  }> = props => {
    const [data, setData] =  useState<any[]>([])

    useEffect(() => {
        fetch("http://localhost:8080/users")
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setData(data.users);
            }
        );
        console.log("Test")
    }, []);

    return (
        <div>
            <h1>Home</h1>
            {data && data.map((user: any) => {
                return <p>{user.username}</p>
            })}
        </div>
    );
};