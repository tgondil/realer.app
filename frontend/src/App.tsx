import React from "react";
import logo from "./logo.svg";
import "./App.css";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Slider from "./components/slider/slider";
import MessageBar from "./components/messageBar/messageBar";
import Chat from "./components/chat/chat";

import {messagesMap} from "./dummy_data/users";
import { useState } from "react";
import { Message } from "./types/types";
//npm install @mui/material @emotion/react @emotion/styled

document.body.style.backgroundColor = "#0B0D0E";


function App() {
  const [selectedFriendId, setSelectedFriendId] = useState<number | null>(null);
  const [selectedChat, setSelectedChat] = useState<Message[]>([]);

  const handleFriendClick = (id: number) => {
    setSelectedFriendId(id);
    const chatToDisplay = messagesMap[id];
    setSelectedChat(chatToDisplay);
  };

  return (
    <Grid
      container
      spacing={2}
      style={{ backgroundColor: "#0F1B29", height: "100vh" }}
    >
      <Grid item xs={4}>
        <Slider onFriendClick={handleFriendClick}/>
      </Grid>
      <Grid item xs={8} style={{ backgroundColor: "#0F1B29" }}>
        {selectedFriendId ? <>
        <Chat messages={selectedChat} receiverId={selectedFriendId }/> 
        </>: <MessageBar />}
      </Grid>
    </Grid>
  );
}

export default App;
