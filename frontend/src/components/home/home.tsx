import React, { useState, useEffect } from "react";
import Grid from "@mui/material/Grid";
import Slider from "../slider/slider";
import Chat from "../chat/chat";
import MessageBar from "../messageBar/messageBar";
import { getMessages } from "../../apis/getMessages"; // Assuming you have a fetchChats API function

interface HomeProps {
  token: string;
}

const Home: React.FC<HomeProps> = ({ token }) => {
  const [selectedFriendId, setSelectedFriendId] = useState<number | null>(null);
  // const [selectedChat, setSelectedChat] = useState<Message[]>([]);
  const [chats, setChats] = useState([]); // State to store chat data

  useEffect(() => {
    const loadChats = async () => {
      try {
        const chatData = await getMessages(token); // Fetch chat data
        setChats(chatData);
      } catch (error) {
        console.error("Error fetching chats:", error);
      }
    };
    loadChats();
  }, []);

  const handleFriendClick = (id: number) => {
    setSelectedFriendId(id);
    //const chatToDisplay = messagesMap[id]; // Update this to use real chat data
    //setSelectedChat(chatToDisplay);
  };

  return (
    <Grid
      container
      spacing={2}
      style={{ backgroundColor: "rgb(11, 13, 14)", height: "100vh" }}
    >
      <Grid item xs={4}>
        <Slider onFriendClick={handleFriendClick} chats={chats} />
      </Grid>
      <Grid item xs={8} style={{ backgroundColor: "rgb(11, 13, 14)" }}>
        {selectedFriendId ? (
          // <Chat messages={selectedChat} receiverId={selectedFriendId} />
          <> </>
        ) : (
          <MessageBar />
        )}
      </Grid>
    </Grid>
  );
};

export default Home;
