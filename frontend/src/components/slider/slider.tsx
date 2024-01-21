import React from "react";
import { useState } from "react";

import IconButton from "@mui/material/IconButton";
import SearchIcon from "@mui/icons-material/Search";
import TextField from "@mui/material/TextField";
import TopBar from "./topBar";

import "./slider.css";
import { friends } from "../../dummy_data/users"; // Assuming friendsMap is imported from your data file
import { Chat } from "../../types/types";

interface Friend {
  id: number;
  name: string;
}

interface SliderProps {
  onFriendClick: (id: number) => void;
  chats: { chatID: number; personID: number; personName: string }[]; // Add type for chats
}

const Slider: React.FC<SliderProps> = ({ onFriendClick, chats }) => {
  const [searchQuery, setSearchQuery] = useState("");

  const filterFriends = (
    query: string,
    friendsMap: Record<number, Friend>,
  ): Friend[] => {
    const friendsArray = Object.values(friendsMap); // Convert map to array
    if (!query) {
      return friendsArray;
    }
    return friendsArray.filter((friend) =>
      friend.name.toLowerCase().includes(query.toLowerCase()),
    );
  };

  const filteredFriends: Friend[] = filterFriends(searchQuery, friends);

  return (
    <div className="container">
      <div className="scrollable-section">
        <TopBar />
        <div className="search-bar">
          <input
            type="text"
            placeholder="Search friends..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
        </div>
        {chats.map((chat) => (
          <div
            key={chat.chatID}
            className="friend-item"
            onClick={() => onFriendClick(chat.personID)}
          >
            <p>{chat.personName}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Slider;
