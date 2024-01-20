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
import { alpha, styled } from '@mui/material/styles';
import InputBase from '@mui/material/InputBase';
import InputLabel from '@mui/material/InputLabel';

const CssTextField = styled(TextField)({
    '& label.Mui-focused': {
      color: '#E0E3E7',
    },
    '& .MuiInput-underline:after': {
      borderBottomColor: '#E0E3E7',
    },
    '& .MuiOutlinedInput-root': {
      '& fieldset': {
        borderColor: '#E0E3E7',
      },
      '&:hover fieldset': {
        borderColor: '#E0E3E7',
      },
      '&.Mui-focused fieldset': {
        borderColor: '#E0E3E7',
      },
    },
  });


const Chat = () => {
  return (
    <><DefaultScreen></DefaultScreen>
    <Grid className="body" style={{paddingRight: "0", background: "white", borderRadius: "50px", width: "100%"}} container spacing={2}>
    <Grid item xs={1} style={{paddingTop: "0"}}> <InsertEmoticonIcon style={{ fontSize: 40, paddingTop: "7px" }}/> <AddCircleOutlineIcon style={{ fontSize: 40, paddingTop: "7px" }}/> </Grid>
    <Grid item xs={10} style={{paddingTop: "0"}}> 
    <CssTextField label="Type a message" id="custom-css-outlined-input" />
    </Grid>
    <Grid item xs={1} style={{paddingTop: "0"}}>
        <MicNoneIcon style={{ fontSize: 45, paddingTop: "7px", paddingRight: "30px", marginLeft: "0px"}}/>
    </Grid>
{/* sick, mui has a search bar component alr ill check it ou */}
        
      </Grid></>
    
  );
};

export default Chat;
