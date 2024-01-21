import { Route, Routes } from "react-router";
import "./App.css";
import Login from "./components/login/login";
import Home from "./components/home/home";

import { useState, useEffect } from "react";

import Cookies from "js-cookie";

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const tokenFromCookie = Cookies.get("token");
    if (tokenFromCookie) {
      setToken(tokenFromCookie);
    }
  }, []);

  const onRegisterSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
  };

  const onLoginSuccess = (newToken: string) => {
    setToken(newToken);
    Cookies.set("token", newToken, { expires: 1 });
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
