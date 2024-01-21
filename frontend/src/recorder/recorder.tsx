import * as React from 'react';
import { AudioRecorder } from 'react-audio-voice-recorder';
import { useForm } from "react-hook-form";

export default function Recorder() {
  const addAudioElement = (blob: Blob) => {
    const url = URL.createObjectURL(blob);
    const audio = document.createElement('audio');
    audio.src = url;
    audio.controls = true;
    onSubmit(audio);
  };

  const { register, handleSubmit } = useForm();

  const onSubmit = async (data:any) => {
    const formData = new FormData();
    formData.append("file", data.file[0]);

    const res = await fetch("http://localhost:5000/upload-file", {
      method: "POST",
      body: formData,
    }).then((res) => res.json());
  };

  return (
    <div>
      <AudioRecorder
        onRecordingComplete={addAudioElement}
        audioTrackConstraints={{
          noiseSuppression: true,
          echoCancellation: true,
          // autoGainControl,
          // channelCount,
          // deviceId,
          // groupId,
          // sampleRate,
          // sampleSize,
        }}
        onNotAllowedOrFound={(err: any) => console.table(err)}
        downloadOnSavePress={true}
        downloadFileExtension="webm"
        mediaRecorderOptions={{
          audioBitsPerSecond: 128000,
        }}
        // showVisualizer={true}
      />
      <br />
    </div>
  );
}