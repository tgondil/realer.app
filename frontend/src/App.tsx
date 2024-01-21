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
      const msg: Message = {
        messageId: message[0].messageID, // Ensure this is a number
        fromPersonID: parseInt(message[0].fromPerson), // Parse to number if necessary
        content: message[0].message, // Ensure this is a string
        timestamp: message[0].timestamp, // Ensure this is a string
      };
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
