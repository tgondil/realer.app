import React, { useState } from "react";
import Grid from "@mui/material/Grid";
import Slider from "../slider/slider";
import Chat from "../chat/chat";
import MessageBar from "../messageBar/messageBar";
import { messagesMap } from "../../dummy_data/users";
import { Message } from "../../types/types";

const Home = () => {
  const [selectedFriendId, setSelectedFriendId] = useState<number | null>(null);
  const [selectedChat, setSelectedChat] = useState<Message[]>([]);

  const handleFriendClick = (id: number) => {
    setSelectedFriendId(id);
    const chatToDisplay = messagesMap[id];
    setSelectedChat(chatToDisplay);
  };

  return (
    <Grid
      container
      spacing={2}
      style={{ backgroundColor: "#0F1B29", height: "100vh" }}
    >
      <Grid item xs={4}>
        <Slider onFriendClick={handleFriendClick}/>
      </Grid>
      <Grid item xs={8} style={{ backgroundColor: "#0F1B29" }}>
        {selectedFriendId ? (
          <Chat messages={selectedChat} receiverId={selectedFriendId} />
        ) : (
          <MessageBar />
        )}
      </Grid>
    </Grid>
  );
};

export default Home;