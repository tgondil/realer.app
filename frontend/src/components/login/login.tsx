import React from "react";
import "./login.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from '@mui/icons-material/InsertEmoticon';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import MicNoneIcon from '@mui/icons-material/MicNone';
import Typewriter from 'typewriter-effect';
import ReactTyped from "react-typed";

const Login = () => {
    return (<>
    <div style={{height: "100vh", display: "flex", alignItems: "center", justifyContent: "center", flexWrap: "wrap", rowGap: "0" }}>
    <h1 className="hero big gradient" >
         realer<br />
         <span className="subtext">
         <ReactTyped
          strings={["Make conversations realer", "Texting never felt realer", "Connecting to people just got realer"]}
          typeSpeed={100}
          loop
          backSpeed={20}
          cursorChar=""
          showCursor={true}
        /></span>
         </h1>
    <h1>
    
    
        
    </h1>
         
    
    <br />
    
    </div>
    </>
        
    );
  };
  
  export default Login;