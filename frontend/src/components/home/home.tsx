import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import Slider from "../slider/slider";
import Chat from "../chat/chat";
import MessageBar from "../messageBar/messageBar";
import { getChats } from "../../apis/getChats";

interface HomeProps {
  token: string;
}

const Home: React.FC<HomeProps> = ({ token }) => {
  const [selectedChatId, setSelectedChatId] = useState<number | null>(null);
  const [chats, setChats] = useState([]);

  useEffect(() => {
    const loadChats = async () => {
      try {
        const chatData = await getChats(token);
        setChats(chatData);
      } catch (error) {
        console.error("Error fetching chats:", error);
      }
    };
    loadChats();
  }, []);

  const handleChatClick = (id: number) => {
    setSelectedChatId(id);
  };

  return (
    <Grid
      container
      spacing={2}
      style={{ backgroundColor: "rgb(11, 13, 14)", height: "100vh" }}
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
