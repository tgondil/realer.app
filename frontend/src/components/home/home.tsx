import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import Slider from "../slider/slider";
import Chat from "../chat/chat";
import MessageBar from "../messageBar/messageBar";
import { getChats } from "../../apis/getChats";
import { socket } from "../../socket";
import { Message } from "../../types/types";

import CircularProgress from "@mui/material/CircularProgress";

interface HomeProps {
  token: string;
}

const Home: React.FC<HomeProps> = ({ token }) => {
  const [selectedChatId, setSelectedChatId] = useState<number | null>(null);
  const [chats, setChats] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  const listenForNewMessages = () => {
    socket.on("new_message", (message: any) => {
      console.log(message);
      const msg: Message = {
        messageId: message[0].messageID,
        fromPersonID: parseInt(message[0].fromPerson),
        content: message[0].message,
        timestamp: message[0].timestamp,
      };

      console.log("I'm here");
      if (msg.fromPersonID === selectedChatId) {
        console.log("new message");
      } else {
        console.log("new message but not selected");
      }
    });

    listenForNewMessages();
  };

  useEffect(() => {
    const loadChats = async () => {
      setIsLoading(true);
      if (token) {
        try {
          const chatData = await getChats(token);
          setChats(chatData);
        } catch (error) {
          console.error("Error fetching chats:", error);
        }
      }
      setIsLoading(false);
    };

    loadChats();
  }, [token]);

  const handleChatClick = (id: number) => {
    setSelectedChatId(id);
  };

  if (isLoading) {
    return (
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
        }}
      >
        <CircularProgress />
      </div>
    );
  }

  return (
    <Grid
      container
      spacing={2}
      style={{
        backgroundColor: "rgb(11, 13, 14)",
        height: "100vh",
        overflow: "clip",
      }}
    >
      <Grid item xs={4}>
        <Slider onFriendClick={handleChatClick} chats={chats} />
      </Grid>
      <Grid item xs={8} style={{ backgroundColor: "rgb(11, 13, 14)" }}>
        {selectedChatId ? (
          <Chat token={token} receiverId={selectedChatId} />
        ) : (
          <MessageBar />
        )}
      </Grid>
    </Grid>
  );
};

export default Home;
