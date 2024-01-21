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
import Button from '@mui/material/Button';
import { FormControl, FormLabel } from '@mui/material';
import { Height } from "@mui/icons-material";

const Login = () => {
    return (<>
    <div style={{height: "100vh", display: "flex", alignItems: "center", justifyContent: "center", flexWrap: "wrap", rowGap: "0" }}>
    <h1 className="hero big gradient" style={{flexBasis: "100%"}}>
         realer<br />
         </h1>

<FormControl sx={{
    width: "50%",
}}>
    <FormLabel>Enter Name</FormLabel>
    <TextField fullWidth variant='filled' label="Username" sx={{
          input: {
            color: "black",
            background: "white",
            fontSize: "20px"
          },
          marginBottom: "20px"
        }}></TextField>
        <FormLabel>Enter Name</FormLabel>
    <TextField fullWidth variant='filled' label="Password" sx={{
        borderRadius: "50px",
          input: {
            color: "black",
            background: "white",
            fontSize: "20px"
          },
        marginBottom: "40px"
        }}></TextField>
    <Button variant="contained" sx={{
        height: "50px",
    }}>Login</Button>
</FormControl>


    <h1 style={{flexBasis: "100%"}}>
    <span className="subtext" >
         <ReactTyped
          strings={["realer conversations", "realer connections"]}
          typeSpeed={100}
          loop={true}
          backSpeed={20}
          cursorChar=""
          smartBackspace={true}
          showCursor={true}
    /></span>

    </h1>

    
         
    
    <br />
    
    </div>
    </>
        
    );
  };
  
  export default Login;