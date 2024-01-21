import { Route, Routes } from "react-router";
import "./App.css";
import { socket } from "./socket";
import Login from "./components/login/login";
import Home from "./components/home/home";

import { useState, useEffect } from "react";

import Cookies from "js-cookie";
import { Message } from "./types/types";

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
  const [token, setToken] = useState<string>("");
  const [isConnected, setIsConnected] = useState(socket.connected);
  const [mapOfUnreadMessagesCount, setMapOfUnreadMessagesCount] = useState<
    Record<number, number>
  >({});

  const [currentChatMessages, setCurrentChatMessages] = useState<Message[]>([]);
  const [selectedChatId, setSelectedChatId] = useState<number | null>(null);

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

    socket.on("new_message", (message: any) => {
      console.log(message);
      const msg: Message = {
        messageId: message[0].messageID,
        fromPersonID: message[0].fromPerson,
        content: message[0].message,
        timestamp: message[0].timestamp,
      };
      console.log(msg)

      console.log("I'm here", selectedChatId);
      if (msg.fromPersonID === selectedChatId) {
        setCurrentChatMessages((prevMessages) => [...prevMessages, msg]);
        console.log("new message");
      } else {
        console.log("new message but not selected");
        // ... (update mapOfUnreadMessagesCount as in previous example)
      }
    });
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
    <>
      <Routes>
        <Route
          path="/home"
          element={
            <Home
              token={token}
              currentChatMessages={currentChatMessages}
              selectedChatId={selectedChatId}
              setSelectedChatId={setSelectedChatId}
            />
          }
        />
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
