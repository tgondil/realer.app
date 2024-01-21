// data/data.tsx
import React, { createContext, useContext, useState, ReactNode } from "react";
import { Message } from "../types/types";

interface ChatDataContextType {
  messages: Message[];
  setMessages: (newMessages: Message[]) => void;
  addMessage: (message: Message) => void;
}

const ChatDataContext = createContext<ChatDataContextType>(null!);

interface ChatDataProviderProps {
  children: ReactNode; // This allows any valid React child (elements, strings, etc.)
}

export const ChatDataProvider: React.FC<ChatDataProviderProps> = ({
  children,
}) => {
  const [messages, setMessages] = useState<Message[]>([]);

  const updateMessages = (newMessages: Message[]) => {
    setMessages(newMessages);
  };

  const addMessage = (newMessage: Message) => {
    setMessages((prevMessages) => [...prevMessages, newMessage]);
  };

  return (
    <ChatDataContext.Provider
      value={{ messages, setMessages: updateMessages, addMessage }}
    >
      {children}
    </ChatDataContext.Provider>
  );
};

export const useChatData = () => useContext(ChatDataContext);
