import React from "react";
import "./chat.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from '@mui/icons-material/InsertEmoticon';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import MicNoneIcon from '@mui/icons-material/MicNone';
import Default from "./default";
import DefaultScreen from "./default";

const Chat = () => {
  return (
    <><DefaultScreen></DefaultScreen>
    <Grid className="body" style={{background: "#090979"}} container spacing={2}>
    <Grid item xs={1}> <InsertEmoticonIcon style={{ fontSize: 40, paddingTop: "7px" }}/> <AddCircleOutlineIcon style={{ fontSize: 40, paddingTop: "7px" }}/> </Grid>
    <Grid item xs={10}> 
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
    <Grid item xs={1}>
        <MicNoneIcon style={{ fontSize: 45, paddingTop: "7px", paddingRight: "30px", marginLeft: "0px" }}/>
    </Grid>
{/* sick, mui has a search bar component alr ill check it ou */}
        
      </Grid></>
    
  );
};

export default Chat;
