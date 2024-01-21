import React from "react";
import "./camera.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from '@mui/icons-material/InsertEmoticon';
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import MicNoneIcon from '@mui/icons-material/MicNone';
import { alpha, styled } from '@mui/material/styles';
import InputBase from '@mui/material/InputBase';
import InputLabel from '@mui/material/InputLabel';
import Webcam from "react-webcam";

const Camera = () => {
    const WebcamCapture = () => {
        const webcamRef = React.useRef<Webcam>(null); // Initialize webcamRef with the correct type
        const [imgSrc, setImgSrc] = React.useState<string | null>(null); // Specify the type of imgSrc
        const [emotion, setEmotion] = React.useState<string | null>(null);

        const capture = React.useCallback(() => {
            if (webcamRef.current) {
                const imageSrc = webcamRef.current.getScreenshot();
                setImgSrc(imageSrc);
            }
        }, [webcamRef, setImgSrc]);


        //use the imageSrc to print to console the highest confidence emotion:
        console.log(imgSrc);



        return (
            <>
                <Webcam
                    audio={false}
                    ref={webcamRef}
                    screenshotFormat="image/jpeg"
                />
                <button onClick={capture}>Capture photo</button>
                {imgSrc && (
                    <img
                        src={imgSrc}
                    />

                )}

                

                
            </>
        );
    };

    return (
        <div style={{height: '100%'}}>
            <WebcamCapture />
        </div>
    );
};

export default Camera;