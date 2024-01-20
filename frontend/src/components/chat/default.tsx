import React from "react";
import "./default.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from '@mui/icons-material/InsertEmoticon';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import MicNoneIcon from '@mui/icons-material/MicNone';

const DefaultScreen = () => {
    return (<>
    <div style={{height: "90vh", display: "flex", alignItems: "center", justifyContent: "center", }}>
    <h1 className="hero" style={{paddingTop: "20px"}}> Talking to your friends has never been <br></br></h1>
    
    <h1 className="hero big gradient" style={{width: "22%", paddingLeft: '0'}} > realer. </h1>
    </div>
    </>
        
        // <h2> HERPES </h2> yapping is wild fr
    );
  };
  
  export default DefaultScreen;