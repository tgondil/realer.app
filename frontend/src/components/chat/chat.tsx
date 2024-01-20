import React from 'react';
import "./chat.css";
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import { Slider } from '@mui/material';

const Chat = () => {
  return (
    <Grid className="body" container spacing={2} style={{backgroundColor: "#0B0D0E", height: "100vh"}}>
    <Grid item xs={4}> </Grid>
    <Grid item xs={8} style={{backgroundColor: "#0F1B29", }}>
    </Grid>
    </Grid>


        <TextField id="filled-basic" label="Filled" variant="filled" sx={{
        input: {
          color: "white",
          background: "#0B0D0E",
          width: "100rem",
          paddingLeft: "0px",
          marginLeft: "0px",
          verticalAlign: "bottom",
        }
      }}/>
  );
};

export default Chat;