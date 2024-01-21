import { Route, Routes } from "react-router";
import "./App.css";
import { socket } from "./socket";
import Login from "./components/login/login";
import Home from "./components/home/home";

import { useState, useEffect } from "react";

import Cookies from "js-cookie";

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
  const [token, setToken] = useState<string>("");
  const [isConnected, setIsConnected] = useState(socket.connected);

  useEffect(() => {
    const tokenFromCookie = Cookies.get("token");
    if (tokenFromCookie) {
      setToken(tokenFromCookie);
    }
  }, []);

  const onLoginSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
    useEffect(() => {
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

      return () => {
        socket.off("connect", onConnect);
        socket.off("disconnect", onDisconnect);
      };
    }, []);
  };
  const onRegisterSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
    useEffect(() => {
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

      return () => {
        socket.off("connect", onConnect);
        socket.off("disconnect", onDisconnect);
      };
    }, []);
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
