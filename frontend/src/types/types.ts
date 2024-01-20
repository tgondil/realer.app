export interface Message {
  messageId: number;
  timestamp: string;
  type: 'audio' | 'text';
  isSenderYou: boolean;
  content: string;
  senderId: number;
  receiverId: number;
}

export interface Friend {
    id: number;
    name: string;
}
 
