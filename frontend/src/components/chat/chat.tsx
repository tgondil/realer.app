import React from "react";
import "./chat.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";

const Chat = () => {
  return (
    <Grid className="body" container spacing={2}>
    <Grid item xs={1}> </Grid>
    <Grid item xs={5}> 
    <TextField id="filled-basic" label="Filled" variant="filled" sx={{
        input: {
          color: "white",
          background: "#0B0D0E",
          width: "100%",
          paddingLeft: "0px",
          marginLeft: "0px",
          verticalAlign: "bottom",
        }
      }}/>
    </Grid>
    <Grid item xs={7} style={{backgroundColor: "#0F1B29", }}>
    </Grid>


        
      </Grid>
  );
};

export default Chat;
