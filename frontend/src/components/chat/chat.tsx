import React from "react";
import { useState, useEffect } from "react";
import "./chat.css";
import { Message } from "../../types/types";
import MessageBar from "../messageBar/messageBar";
import Grid from "@mui/material/Grid";
import { Slider } from "@mui/material";
import InsertEmoticonIcon from "@mui/icons-material/InsertEmoticon";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import MicNoneIcon from "@mui/icons-material/MicNone";
import { alpha, styled } from "@mui/material/styles";
import InputBase from "@mui/material/InputBase";
import InputLabel from "@mui/material/InputLabel";
import TextField from "@mui/material/TextField";
import { getChat } from "../../apis/getChat";

const CssTextField = styled(TextField)({
  "& label.Mui-focused": {
    color: "#E0E3E7",
  },
  "& .MuiInput-underline:after": {
    borderBottomColor: "#E0E3E7",
  },
  "& .MuiOutlinedInput-root": {
    "& fieldset": {
      borderColor: "#E0E3E7",
    },
    "&:hover fieldset": {
      borderColor: "#E0E3E7",
    },
    "&.Mui-focused fieldset": {
      borderColor: "#E0E3E7",
    },
  },
});

interface ChatProps {
  receiverId: number;
  token: string;
}

const Chat: React.FC<ChatProps> = ({ receiverId, token }) => {
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const chatMessages = await getChat(token, receiverId); // Use token in API call
        setMessages(chatMessages);
      } catch (error) {
        console.error("Error fetching chat messages:", error);
      }
    };

    if (receiverId !== null) {
      fetchMessages();
    }
  }, [token, receiverId]); // Add token as a dependency

  return (
    <div className="chat">
      <div className="chat-header">
        <img
          className="friend-image"
          src={`https://i.pravatar.cc/150?img=${receiverId}`}
          alt={receiverId.toString()}
        />
        <h1 className="header"> {receiverId} </h1>
      </div>
      <div className="chat-messages">
        {messages.map((message) => (
          <div
            key={message.messageId}
            className={`message ${
              message.fromPersonID !== receiverId ? "sent" : "received"
            }`}
          >
            <p className="font"> {message.content} </p>
          </div>
        ))}
      </div>
      <Grid
        style={{
          paddingRight: "0",
          background: "white",
          borderRadius: "50px",
          width: "100%",
          marginBottom: "0px",
        }}
        className="bar"
        container
        spacing={2}
      >
        <Grid item xs={1} style={{ paddingTop: "0" }}>
          {" "}
          <AddCircleOutlineIcon
            style={{ fontSize: 40, paddingTop: "7px" }}
          />{" "}
        </Grid>
        <Grid item xs={10} style={{ paddingTop: "0" }}>
          <CssTextField label="Type a message" id="custom-css-outlined-input" />
        </Grid>
        <Grid item xs={1} style={{ paddingTop: "0" }}>
          <MicNoneIcon
            style={{
              fontSize: 45,
              paddingTop: "7px",
              marginLeft: "0px",
            }}
          />
        </Grid>
      </Grid>
    </div>
  );
};

export default Chat;
