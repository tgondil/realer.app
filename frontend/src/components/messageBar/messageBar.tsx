import React from "react";
import "./messageBar.css";
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


const MessageBar = () => {
  return (
    <>
      <DefaultScreen></DefaultScreen>
      <Grid
        style={{
          paddingRight: "0",
          margin: "10px",
          background: "white",
          borderRadius: "50px 50px 0px 0px",
          width: "100%",
        }}
        className="bar"
        container
        spacing={2}
      >
        {/*<Grid item xs={1} style={{ paddingTop: "0" }}>
          {" "}
          <AddCircleOutlineIcon style={{ fontSize: 40, paddingTop: "7px" }} />{" "}
        </Grid>
        <Grid item xs={10} style={{ paddingTop: "0" }}>
          <CssTextField label="Type a message" id="custom-css-outlined-input" />
        </Grid>
        <Grid item xs={1} style={{ paddingTop: "0" }}>
          <MicNoneIcon
            style={{
              fontSize: 45,
              paddingTop: "7px",
              paddingRight: "30px",
              marginLeft: "0px",
            }}
          />
        </Grid>*/}
      </Grid>
    </>
  );
};


export default MessageBar;