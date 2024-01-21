import React from "react";
import { useState } from "react";

import TopBar from "./topBar";

import "./slider.css";
import { Chat } from "../../types/types";

interface SliderProps {
  onFriendClick: (id: number) => void;
  chats: Chat[]; // Add type for chats
}

const Slider: React.FC<SliderProps> = ({ onFriendClick, chats }) => {
  const [searchQuery, setSearchQuery] = useState("");

  const filterChats = (query: string, chats: Chat[]): Chat[] => {
    if (!query) {
      return chats;
    }
    return chats.filter((chat) =>
      chat.PersonName.toLowerCase().includes(query.toLowerCase()),
    );
  };

  const filteredChats: Chat[] = filterChats(searchQuery, chats);

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
        {filteredChats.map((chat) => (
          <div
            key={chat.ChatID}
            className="friend-item"
            onClick={() => onFriendClick(chat.PersonID)}
          >
            <p> {chat.PersonName} </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Slider;
