import React from 'react'

import { Container, Grid, Card, CardContent, Typography } from '@mui/material'


interface ShowPostsListProps {
    data: {
        title: string,
        content: string,
        genre: string,
        author: string,
    }[]
}


export const ShowPostsList: React.FC<ShowPostsListProps> = ({data}) => {


    return (
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
    )
}