// CreatePostForm.tsx
import React, { useState } from "react";
import {
  Button,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Grid,
  Typography,
  Paper,
  Box,
} from "@mui/material";

interface CreatePostFormProps {
  handleSubmit: (title: string, content: string, genre: string) => void;
}

export const CreatePostForm: React.FC<CreatePostFormProps> = ({ handleSubmit }) => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [genre, setGenre] = useState("");

  const genres = ["Technology", "Sports", "Entertainment", "Science", "Health"];

  const handleFormSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    handleSubmit(title, content, genre);
    setTitle("");
    setContent("");
    setGenre("");
  };

  return (
    <Box component={Paper} p={3} maxWidth="sm" mx="auto" style={{marginTop:20}}>
      <form onSubmit={handleFormSubmit}>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <TextField
              required
              fullWidth
              label="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              autoComplete="off" // Add this line to prevent the browser from treating it as a password input
              variant="filled"
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              required
              fullWidth
              multiline
              rows={4}
              label="Content"
              value={content}
              onChange={(e) => setContent(e.target.value)}
            />
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth>
              <InputLabel required>Genre</InputLabel>
              <Select
                value={genre}
                onChange={(e) => setGenre(e.target.value as string)}
              >
                {genres.map((item, index) => (
                  <MenuItem key={index} value={item}>
                    {item}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button type="submit" variant="contained" color="primary">
              Submit
            </Button>
          </Grid>
        </Grid>
      </form>
    </Box>
  );
};
