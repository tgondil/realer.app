import React from "react";
import { useState } from "react";
import "./login.css";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from "@mui/icons-material/InsertEmoticon";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import MicNoneIcon from "@mui/icons-material/MicNone";
import Typewriter from "typewriter-effect";
import ReactTyped from "react-typed";
import Button from "@mui/material/Button";
import { FormControl, FormLabel } from "@mui/material";
import { Height } from "@mui/icons-material";
import { login } from "../../apis/login";
import { signup } from "../../apis/signup";
import { useNavigate } from "react-router-dom";

interface LoginProps {
  onLoginSuccess: (token: string) => void;
  onRegisterSuccess: (token: string) => void;
}

const Login: React.FC<LoginProps> = ({ onLoginSuccess, onRegisterSuccess }) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [isLoginMode, setIsLoginMode] = useState(true);
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const receivedToken = await login(username, password);
      onLoginSuccess(receivedToken);
      navigate("/home");
    } catch (error) {
      console.error("Login error:", error);
    }
  };

  const toggleMode = () => {
    setIsLoginMode(!isLoginMode);
  };

  const handleRegister = async () => {
    try {
      const receivedToken = await signup(username, password);
      onRegisterSuccess(receivedToken);
      navigate("/home");
    } catch (error) {
      console.error("Register error:", error);
    }
  };

  return (
    <>
      <div
        style={{
          height: "100vh",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          flexWrap: "wrap",
          rowGap: "0",
        }}
      >
        <h1 className="hero big gradient" style={{ flexBasis: "100%" }}>
          realer
          <br />
        </h1>

        <FormControl
          sx={{
            width: "50%",
          }}
        >
          <FormLabel>Enter Name</FormLabel>
          <TextField
            fullWidth
            variant="filled"
            label="Username"
            onChange={(e) => setUsername(e.target.value)}
            sx={{
              input: {
                color: "black",
                background: "white",
                fontSize: "20px",
              },
              marginBottom: "20px",
            }}
          ></TextField>
          <FormLabel>Enter Name</FormLabel>
          <TextField
            fullWidth
            variant="filled"
            type="password"
            label="Password"
            onChange={(e) => setPassword(e.target.value)}
            sx={{
              borderRadius: "50px",
              input: {
                color: "black",
                background: "white",
                fontSize: "20px",
              },
              marginBottom: "40px",
            }}
          ></TextField>
          {isLoginMode ? (
            <Button
              variant="outlined"
              onClick={handleLogin}
              sx={{ height: "50px", border: 3 }}
            >
              Login
            </Button>
          ) : (
            <Button
              variant="outlined"
              onClick={handleRegister}
              sx={{ height: "50px", border: 3 }}
            >
              Register
            </Button>
          )}
        </FormControl>

        <h1 style={{ flexBasis: "100%" }}>
          <span className="subtext">
            <ReactTyped
              strings={["realer conversations", "realer connections"]}
              typeSpeed={100}
              loop={true}
              backSpeed={20}
              cursorChar=""
              smartBackspace={true}
              showCursor={true}
            />
          </span>
        </h1>

        <h1 className="subtext2">
          {isLoginMode ? (
            <>
              Not a member?{" "}
              <a href="#!" onClick={toggleMode}>
                <u>Sign up</u>
              </a>
            </>
          ) : (
            <>
              Already a member?{" "}
              <a href="#!" onClick={toggleMode}>
                <u>Log in</u>
              </a>
            </>
          )}
        </h1>

        <br />
      </div>
    </>
  );
};

export default Login;
