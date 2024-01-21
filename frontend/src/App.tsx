import React from "react";
import { Route, Routes } from "react-router";
import logo from "./logo.svg";
import "./App.css";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Slider from "./components/slider/slider";
import MessageBar from "./components/messageBar/messageBar";
import Chat from "./components/chat/chat";
import Login from "./components/login/login";
import Home from "./components/home/home";

import { useState } from "react";
import { Message } from "./types/types";
//npm install @mui/material @emotion/react @emotion/styled

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
  const [selectedFriendId, setSelectedFriendId] = useState<number | null>(null);
  const [selectedChat, setSelectedChat] = useState<Message[]>([]);
  const [token, setToken] = useState<string>("");

  const onLoginSuccess = (newToken: string) => {
    setToken(newToken);
  };

  const onRegisterSuccess = (newToken: string) => {
    setToken(newToken);
  };

  return (
    <>
      <Routes>
        <Route path="/home" element={<Home token={token} />} />
        <Route
          index
          element={
            <Login
              onLoginSuccess={onLoginSuccess}
              onRegisterSuccess={onRegisterSuccess}
            />
          }
        />
      </Routes>
    </>
  );
}

export default App;
