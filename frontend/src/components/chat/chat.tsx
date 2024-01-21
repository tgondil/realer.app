import React from "react";
import { useState, useEffect } from "react";
import "./chat.css";
import { Message } from "../../types/types";
import Grid from "@mui/material/Grid";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import MicNoneIcon from "@mui/icons-material/MicNone";
import { styled } from "@mui/material/styles";
import TextField from "@mui/material/TextField";
import { getChat } from "../../apis/getChat";
import { getUserName } from "../../apis/getUserName";
import { sendMessage } from "../../apis/sendMessage";
import { useRef } from "react";

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
  const [newMessage, setNewMessage] = useState("");

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const chatMessages = await getChat(token, receiverId);
        setMessages(chatMessages);
      } catch (error) {
        console.error("Error fetching chat messages:", error);
      }
    };

    if (receiverId !== null) {
      fetchMessages();
    }
  }, [token, receiverId]);

  const [userName, setUserName] = useState<string>("");

  useEffect(() => {
    const fetchUserName = async () => {
      try {
        const name = await getUserName(token, receiverId);
        setUserName(name);
      } catch (error) {
        console.error("Error fetching user name:", error);
      }
    };

    if (receiverId !== null) {
      fetchUserName();
    }
  }, [token, receiverId]);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages]);

  const handleKeyPress = async (event: React.KeyboardEvent) => {
    if (event.key === "Enter" && newMessage.trim() !== "") {
      event.preventDefault();
      // Call the API to send the message
      try {
        await sendMessage(token, receiverId, newMessage); // Replace with your actual API call
        setNewMessage(""); // Reset the input field after sending
        // Optionally, fetch the latest messages or update the UI
      } catch (error) {
        console.error("Error sending message:", error);
      }
    }
  };

  const messagesEndRef =  React.useRef<HTMLInputElement>(null);

  return (
    <div className="chat">
      <div className="chat-header">
        <img
          className="friend-image"
          src={`https://i.pravatar.cc/150?img=${receiverId}`}
          alt={receiverId.toString()}
        />
        <h1 className="header"> {userName} </h1>
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
        <div ref={messagesEndRef} />
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
          <CssTextField
            label="Type a message"
            id="custom-css-outlined-input"
            value={newMessage}
            onChange={(e) => setNewMessage(e.target.value)}
            onKeyPress={handleKeyPress} // Add the key press handler here
          />
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
