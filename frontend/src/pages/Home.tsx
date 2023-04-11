import React from "react";
import {useEffect, useState} from "react";
import ButtonAppBar from "../components/ButtonAppBar";
import { Grid, Card, CardContent, Typography, Container } from "@mui/material"
import {CreatePostForm} from "../components/CreatePostForm";

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
            <Container maxWidth="md" style={{ marginTop: 20 }}>
                <Grid container spacing={2} direction="column">
                    {data.map((post, index) => (
                        <Grid item key={index}>
                            <Card>
                                <CardContent>
                                    <Typography variant="h5" component="h2">
                                        {post.title}
                                    </Typography>
                                    <Typography color="textSecondary" gutterBottom>
                                        By: {post.author}
                                    </Typography>
                                    <Typography variant="body2" component="p">
                                        Genre: {post.genre}
                                    </Typography>
                                    <Typography variant="body2" component="p">
                                        {post.content}
                                    </Typography>
                                </CardContent>
                            </Card>
                        </Grid>
                    ))}
                </Grid>
            </Container>
        </>
    );
};