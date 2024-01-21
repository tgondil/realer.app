import { Route, Routes } from "react-router";
import "./App.css";
import Login from "./components/login/login";
import Home from "./components/home/home";

import { useState } from "react";

document.body.style.backgroundColor = "rgb(11, 13, 14)";

function App() {
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
