import React from 'react';
import logo from './logo.svg';
import './App.css';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import Slider from './components/slider/slider';
import Chat from './components/chat/chat';
//npm install @mui/material @emotion/react @emotion/styled

document.body.style.backgroundColor = '#0B0D0E';

function App() {
  return (
    <Grid className="body" container spacing={2} style={{backgroundColor: "#0B0D0E", height: "100vh"}}>
    <Grid item xs={4}>
      <Slider />
    </Grid>
    <Grid item xs={8} style={{backgroundColor: "#0F1B29"}}>
      <Chat />
    </Grid>
</Grid>
  );
}

export default App;
