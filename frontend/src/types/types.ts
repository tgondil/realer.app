export interface Message {
  messageId: number;
  timestamp: string;
  content: string;
  fromPersonID: number;
}

export interface Friend {
  id: number;
  name: string;
}

export interface Chat {
  ChatID: number;
  PersonID: number;
  PersonName: string;
}
