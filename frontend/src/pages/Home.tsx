import React from "react";
import {useEffect, useState} from "react";
import ButtonAppBar from "../components/ButtonAppBar";
import { Grid, Card, CardContent, Typography, Container } from "@mui/material"
import {CreatePostForm} from "../components/CreatePostForm";
import {ShowPostsList} from "../components/ShowPostsList";

export const Home: React.FC<{  }> = props => {
    const [data, setData] =  useState<any[]>([])

    useEffect(() => {
        fetch("http://localhost:8080/home")
            .then(response => response.json())
            .then(data => {
                console.log(data)
                setData(data);
            }
        );
    }, []);

    const handleSunmit = (title: string, content: string, genre: string) => {

        // post title, content and genre to backend
        fetch("http://localhost:8080/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                title,
                content,
                genre,
            }),
        })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            setData([...data]);
        });
    }

    return (
        <>
            <ButtonAppBar />
            <CreatePostForm handleSubmit={handleSunmit} />
            <ShowPostsList data={data} />
        </>
    );
};