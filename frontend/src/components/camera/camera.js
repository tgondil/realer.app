import * as faceapi from 'face-api.js';
import React, { useRef } from 'react';
import { useState, useEffect } from "react";

var flag = false;

function Camera() {

  const [modelsLoaded, setModelsLoaded] = React.useState(false);
  const [captureVideo, setCaptureVideo] = React.useState(false);

  const videoRef = React.useRef();
  const videoHeight = 480;
  const videoWidth = 640;
  const canvasRef = React.useRef();
  const messagesEndRef = useRef(null);

  React.useEffect(() => {
    const loadModels = async () => {
      const MODEL_URL = process.env.PUBLIC_URL + '/models';

      Promise.all([
        faceapi.nets.tinyFaceDetector.loadFromUri(MODEL_URL),
        faceapi.nets.faceLandmark68Net.loadFromUri(MODEL_URL),
        faceapi.nets.faceRecognitionNet.loadFromUri(MODEL_URL),
        faceapi.nets.faceExpressionNet.loadFromUri(MODEL_URL),
      ]).then(setModelsLoaded(true));
    }
    loadModels();
  }, []);

  const startVideo = () => {
    setCaptureVideo(true);
    navigator.mediaDevices
      .getUserMedia({ video: { width: 300 } })
      .then(stream => {
        let video = videoRef.current;
        video.srcObject = stream;
        video.play();
      })
      .catch(err => {
        console.error("error:", err);
      });
      messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
    
      setTimeout(function() { closeWebcam(); }, 10000);
  }

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  //create a counter for each emotion and then display the emotion with the highest count
  var happyCounter = 0;
  var veryHappyCounter = 0;
  var sadCounter = 0;
  var angryCounter = 0;
  var fearfulCounter = 0;
  var surprisedCounter = 0;
  var disgustedCounter = 0;
  var neutralCounter = 0;

  //create a hashmap of emotions and their counters
  var emotions = {
    happy: happyCounter,
    veryHappy: veryHappyCounter,
    sad: sadCounter,
    angry: angryCounter,
    fearful: fearfulCounter,
    surprised: surprisedCounter,
    disgusted: disgustedCounter,
    neutral: neutralCounter
  }

  let max;
  


  const handleVideoOnPlay = () => {
    if (flag !== true) {
    setInterval(async () => {
      if (canvasRef && canvasRef.current) {
        canvasRef.current.innerHTML = faceapi.createCanvasFromMedia(videoRef.current);
        const displaySize = {
          width: videoWidth,
          height: videoHeight
        }

        //how to store a counter for each emotion and then display the emotion with the highest count?

        faceapi.matchDimensions(canvasRef.current, displaySize);

        const detections = await faceapi.detectAllFaces(videoRef.current, new faceapi.TinyFaceDetectorOptions()).withFaceLandmarks().withFaceExpressions();

        const resizedDetections = faceapi.resizeResults(detections, displaySize);

        canvasRef && canvasRef.current && canvasRef.current.getContext('2d').clearRect(0, 0, videoWidth, videoHeight);
        canvasRef && canvasRef.current && faceapi.draw.drawDetections(canvasRef.current, resizedDetections);
        canvasRef && canvasRef.current && faceapi.draw.drawFaceLandmarks(canvasRef.current, resizedDetections);
        canvasRef && canvasRef.current && faceapi.draw.drawFaceExpressions(canvasRef.current, resizedDetections);

        if (resizedDetections[0]) {
          let expressions = resizedDetections[0].expressions;
          
          if (expressions.disgusted > 0.3) {
            emotions.disgusted++;
            /*//set all other counters to 0
            for (var key in emotions) {
              if (key !== "disgusted") {
                emotions[key] = 0;
              }
            }
            if (emotions.disgusted > 3) {
              console.log("disgusted");
              emotions.disgusted = 0;
            }*/
          }
          else if (expressions.surprised > 0.4) {
            emotions.surprised++;
            /*for (var key in emotions) {
              if (key !== "surprised") {
                emotions[key] = 0;
              }
            }
            if (emotions.surprised > 3) {
              console.log("surprised");
              emotions.surprised = 0;
            }*/
          }
          else if (expressions.fearful > 0.3) {
            emotions.fearful++;
            /*for (var key in emotions) {
              if (key !== "fearful") {
                emotions[key] = 0;
              }
            }
            if (emotions.fearful > 5) {
              console.log("fearful");
              emotions.fearful = 0;
            }*/
          }
          else if (expressions.angry > 0.6) {
            emotions.angry++;
            /*for (var key in emotions) {
              if (key !== "angry") {
                emotions[key] = 0;
              }
            }
            if (emotions.angry > 4) {
              console.log("angry");
              emotions.angry = 0;
            }*/
          }
          else if (expressions.sad > 0.6) {
            emotions.sad++;
            /*for (var key in emotions) {
              if (key !== "sad") {
                emotions[key] = 0;
              }
            }
            if (emotions.sad > 5) {
              console.log("sad");
              emotions.sad = 0;
            }*/
          }
          else if (expressions.happy > 0.6 && expressions.happy < 0.99) {
            emotions.happy++;
            /*if (emotions.happy > 5) {
              console.log("happy");
              emotions.happy = 0;
              for (var key in emotions) {
                if (emotions.veryHappy < 5 && emotions.happy > 2) {
                  console.log("happy");
                }
                emotions[key] = 0;
              
            }
            }*/
            
            
          }
          else if (expressions.happy >= 0.99) {
            emotions.veryHappy++;
            /*for (var key in emotions) {
              if (key !== "veryHappy") {
                emotions[key] = 0;
              }
            }
            if (emotions.veryHappy > 6) {
              console.log("very happy");
              emotions.veryHappy = 0;
            }*/
          }
          else if (expressions.neutral > 0.9) {
            emotions.neutral++;
            /*for (var key in emotions) {
              if (key !== "neutral") {
                emotions[key] = 0;
              }
            }
            if (emotions.neutral > 10) {
              console.log("neutral");
              emotions.neutral = 0;
            }*/
          

          //if alternating between happy and very happy, then happy:

        }
      }
    }

    max = Object.keys(emotions).reduce(function(a, b){ return emotions[a] > emotions[b] ? a : b });
    console.log('Current max value:', max);

        if (emotions.veryHappy === 15) {
          console.log("very happy");
          emotions.veryHappy++;
          flag = true;
          closeWebcam();
        }

        if (emotions.happy === 15) {
            console.log("very happy");
            emotions.veryHappy++;
            flag = true;
            closeWebcam();
        }

        if (emotions.sad === 15) {
          console.log("sad");
          emotions.sad++;
          flag = true;
          closeWebcam();
        }

        if (emotions.angry === 15) {
          console.log("angry");
          emotions.angry++;
          flag = true;
          closeWebcam();
        }

        if (emotions.fearful === 15) {
          console.log("fearful");
          emotions.fearful++;
          flag = true;
          closeWebcam();
        }

        if (emotions.surprised === 15) {
          console.log("surprised");
          emotions.surprised++;
          flag = true;
          closeWebcam();
        }

        if (emotions.disgusted === 15) {
          console.log("disgusted");
          emotions.disgusted++;
          flag = true;
          closeWebcam();
        }
  
    }, 100)
  }
  }

  const closeWebcam = () => {
    if (videoRef && videoRef.current) {
      videoRef.current.pause();
      videoRef.current.srcObject.getTracks()[0].stop();
    }

    
    setCaptureVideo(false);
  }

  return (
    <>
    <div>
      <div style={{ textAlign: 'center', padding: '10px' }}>
        {
          captureVideo && modelsLoaded ?
            <button onClick={closeWebcam} style={{ cursor: 'pointer', backgroundColor: 'green', color: 'white', padding: '15px', fontSize: '25px', border: 'none', borderRadius: '10px' }}>
              Close Webcam
            </button>
            :
            <button onClick={startVideo} style={{ cursor: 'pointer', backgroundColor: 'green', color: 'white', padding: '15px', fontSize: '25px', border: 'none', borderRadius: '10px' }}>
              Open Webcam
            </button>
        }
      </div>
      {
        captureVideo ?
          modelsLoaded ?
            <div>
              <div style={{ display: 'flex', justifyContent: 'center', padding: '10px' }}>
                <video ref={videoRef} height={videoHeight} width={videoWidth} onPlay={handleVideoOnPlay} style={{ borderRadius: '10px' }} />
                <canvas ref={canvasRef} style={{ position: 'absolute' }} />
              </div>
            </div>
            :
            <div>loading...</div>
          :
          <>
          </>
      }
    
    </div>
    <div ref={messagesEndRef} />
    </>
  );
}

export default Camera;