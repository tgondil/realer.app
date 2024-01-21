import { Route, Routes } from "react-router";
import React from "react";
import "./App.css";
import { socket } from "./socket";
import Login from "./components/login/login";
import Home from "./components/home/home";
import { ChatProvider } from "./ChatContext";
import { useState, useEffect, useRef } from "react";

import Cookies from "js-cookie";
import { Message } from "./types/types";
import listenForNewMessages from "./components/home/home";

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
  const [token, setToken] = useState<string>("");
  const [isConnected, setIsConnected] = useState(socket.connected);
  const [mapOfUnreadMessagesCount, setMapOfUnreadMessagesCount] = useState<
    Record<number, number>
  >({});

  const [currentChatMessages, setCurrentChatMessages] = useState<Message[]>([]);

  const [selectedChatId, setSelectedChatId] = useState<number | null>(null);
  const selectedChatIdRef = useRef<number | null>(null);
  const currentChatMessagesRef = useRef<Message[]>([]);

  useEffect(() => {
    currentChatMessagesRef.current = currentChatMessages; // Update the ref whenever currentChatMessages changes
  }, [currentChatMessages]);

  useEffect(() => {
    selectedChatIdRef.current = selectedChatId; // Update the ref whenever selectedChatId changes
  }, [selectedChatId]);

  useEffect(() => {
    const tokenFromCookie = Cookies.get("token");
    if (tokenFromCookie) {
      setToken(tokenFromCookie);
    }
  }, []);

  const onLoginSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
    function onConnect() {
      setIsConnected(true);
    }

    function onDisconnect() {
      setIsConnected(false);
    }

    socket.connect();

    socket.emit("join_subscription", newToken);

    socket.on("connect", onConnect);
    socket.on("disconnect", onDisconnect);

    listenForNewMessages();

    return () => {
      socket.off("connect", onConnect);
      socket.off("disconnect", onDisconnect);
    };
  };
  const onRegisterSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
    function onConnect() {
      console.log("connected socket");
      setIsConnected(true);
    }

    function onDisconnect() {
      console.log("disconnected socket");
      setIsConnected(false);
    }

    socket.connect();

    socket.emit("join_subscription", newToken);

    socket.on("connect", onConnect);
    socket.on("disconnect", onDisconnect);

    socket.on("new_message", (message: any) => {
      console.log(message);
    });

    return () => {
      socket.off("connect", onConnect);
      socket.off("disconnect", onDisconnect);
    };
  };

  return (
    <ChatProvider>
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
    </ChatProvider>
  );
}

export default App;
