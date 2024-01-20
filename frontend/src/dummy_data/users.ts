import { Friend, Message } from '../types/types';
 
export const friends: Record<number, Friend> = {
  1: { id: 1, name: 'Timothy Edwards' },
  2: { id: 2, name: 'Mikayla Brown' },
  3: { id: 3, name: 'Tiffany Calhoun' },
  4: { id: 4, name: 'Patrick Jones' },
  5: { id: 5, name: 'Sara Good' },
  6: { id: 6, name: 'Rachel Becker' },
  7: { id: 7, name: 'Mrs. Brianna Adams' },
  8: { id: 8, name: 'Michael Williamson' },
  9: { id: 9, name: 'William Gray' },
  10: { id: 10, name: 'Mrs. Ashley Lucas MD' },
  11: { id: 11, name: 'Christine Morales' },
  12: { id: 12, name: 'Melissa Smith' },
  13: { id: 13, name: 'Tyler Horton' },
  14: { id: 14, name: 'Noah Mccormick' },
  15: { id: 15, name: 'Alicia Ferrell' },
};



export const messagesMap: Record<number, Message[]> = {
  1: [
    { messageId: 1, timestamp: '2024-01-20 08:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 1 },
    { messageId: 2, timestamp: '2024-01-20 08:05', type: 'text', isSenderYou: false, content: 'Hey, good morning to you too!', senderId: 1, receiverId: 0 },
    { messageId: 3, timestamp: '2024-01-20 08:10', type: 'text', isSenderYou: true, content: 'Did you finish the project?', senderId: 0, receiverId: 1 },
    { messageId: 4, timestamp: '2024-01-20 08:15', type: 'audio', isSenderYou: false, content: 'Audio Message', senderId: 1, receiverId: 0 },
    { messageId: 5, timestamp: '2024-01-20 08:20', type: 'text', isSenderYou: true, content: 'Let’s meet for coffee later.', senderId: 0, receiverId: 1 },
    { messageId: 6, timestamp: '2024-01-20 08:25', type: 'text', isSenderYou: false, content: 'Sure, sounds great!', senderId: 1, receiverId: 0 },
    { messageId: 7, timestamp: '2024-01-20 08:30', type: 'audio', isSenderYou: true, content: 'Audio Message', senderId: 0, receiverId: 1 },
    { messageId: 8, timestamp: '2024-01-20 08:35', type: 'text', isSenderYou: false, content: 'I’ll see you at our usual spot.', senderId: 1, receiverId: 0 },
    { messageId: 9, timestamp: '2024-01-20 08:40', type: 'text', isSenderYou: true, content: 'Perfect, see you there!', senderId: 0, receiverId: 1 },
    { messageId: 10, timestamp: '2024-01-20 08:45', type: 'text', isSenderYou: false, content: 'Can’t wait!', senderId: 1, receiverId: 0 },
    { messageId: 11, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'I’m leaving now, will be there in 10.', senderId: 0, receiverId: 1 },
    { messageId: 12, timestamp: '2024-01-20 09:05', type: 'audio', isSenderYou: false, content: 'Audio Message', senderId: 1, receiverId: 0 },
    { messageId: 13, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Just parked. Where are you?', senderId: 0, receiverId: 1 },
    { messageId: 14, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'I’m inside, got us a table.', senderId: 1, receiverId: 0 },
    { messageId: 15, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Walking in now.', senderId: 0, receiverId: 1 },
    { messageId: 16, timestamp: '2024-01-20 09:25', type: 'audio', isSenderYou: false, content: 'Audio Message', senderId: 1, receiverId: 0 },
    { messageId: 17, timestamp: '2024-01-20 09:30', type: 'text', isSenderYou: true, content: 'Great catching up with you!', senderId: 0, receiverId: 1 },
    { messageId: 18, timestamp: '2024-01-20 09:35', type: 'text', isSenderYou: false, content: 'Absolutely, let’s do this again soon.', senderId: 1, receiverId: 0 },
    { messageId: 19, timestamp: '2024-01-20 09:40', type: 'audio', isSenderYou: true, content: 'Audio Message', senderId: 0, receiverId: 1 },
    { messageId: 20, timestamp: '2024-01-20 09:45', type: 'text', isSenderYou: false, content: 'Will plan something for next week.', senderId: 1, receiverId: 0 },
  ],
  2: [
    { messageId: 6, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 2 },
    { messageId: 7, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 2, receiverId: 0 },
    { messageId: 8, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 2 },
    { messageId: 9, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 2, receiverId: 0 },
    { messageId: 10, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 2 },
  ],
  3: [
    { messageId: 11, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 3 },
    { messageId: 12, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 3, receiverId: 0 },
    { messageId: 13, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 3 },
    { messageId: 14, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 3, receiverId: 0 },
    { messageId: 15, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 3 },
  ],
  4: [
    { messageId: 16, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 4 },
    { messageId: 17, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 4, receiverId: 0 },
    { messageId: 18, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 4 },
    { messageId: 19, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 4, receiverId: 0 },
    { messageId: 20, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 4 },
  ],
  5: [
    { messageId: 21, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 5 },
    { messageId: 22, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 5, receiverId: 0 },
    { messageId: 23, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 5 },
    { messageId: 24, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 5, receiverId: 0 },
    { messageId: 25, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 5 },
  ],
  6: [
    { messageId: 26, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 6 },
    { messageId: 27, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 6, receiverId: 0 },
    { messageId: 28, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 6 },
    { messageId: 29, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 6, receiverId: 0 },
    { messageId: 30, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 6 },
  ],

// Continuing from the previous data...
  7: [
    { messageId: 31, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 7 },
    { messageId: 32, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 7, receiverId: 0 },
    { messageId: 33, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 7 },
    { messageId: 34, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 7, receiverId: 0 },
    { messageId: 35, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 7 },
  ],
  8: [
    { messageId: 36, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 8 },
    { messageId: 37, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 8, receiverId: 0 },
    { messageId: 38, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 8 },
    { messageId: 39, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 8, receiverId: 0 },
    { messageId: 40, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 8 },
  ],
  9: [
    { messageId: 41, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 9 },
    { messageId: 42, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 9, receiverId: 0 },
    { messageId: 43, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 9 },
    { messageId: 44, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 9, receiverId: 0 },
    { messageId: 45, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 9 },
  ],
  10: [
    { messageId: 46, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 10 },
    { messageId: 47, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 10, receiverId: 0 },
    { messageId: 48, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 10 },
    { messageId: 49, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 10, receiverId: 0 },
    { messageId: 50, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 10 },
  ],
  11: [
    { messageId: 51, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 11 },
    { messageId: 52, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 11, receiverId: 0 },
    { messageId: 53, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 11 },
    { messageId: 54, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 11, receiverId: 0 },
    { messageId: 55, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 11 },
  ],
  12: [
    { messageId: 56, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 12 },
    { messageId: 57, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 12, receiverId: 0 },
    { messageId: 58, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 12 },
    { messageId: 59, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 12, receiverId: 0 },
    { messageId: 60, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 12 },
  ],
  13: [
    { messageId: 61, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 13 },
    { messageId: 62, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 13, receiverId: 0 },
    { messageId: 63, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 13 },
    { messageId: 64, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 13, receiverId: 0 },
    { messageId: 65, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 13 },
  ],
  14: [
    { messageId: 66, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 14 },
    { messageId: 67, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 14, receiverId: 0 },
    { messageId: 68, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 14 },
    { messageId: 69, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 14, receiverId: 0 },
    { messageId: 70, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 14 },
  ],
  15: [
    { messageId: 71, timestamp: '2024-01-20 09:00', type: 'text', isSenderYou: true, content: 'Good morning!', senderId: 0, receiverId: 15 },
    { messageId: 72, timestamp: '2024-01-20 09:05', type: 'text', isSenderYou: false, content: 'Morning! How are you?', senderId: 15, receiverId: 0 },
    { messageId: 73, timestamp: '2024-01-20 09:10', type: 'text', isSenderYou: true, content: 'Doing well, thanks. You?', senderId: 0, receiverId: 15 },
    { messageId: 74, timestamp: '2024-01-20 09:15', type: 'text', isSenderYou: false, content: 'Pretty good, just started my day.', senderId: 15, receiverId: 0 },
    { messageId: 75, timestamp: '2024-01-20 09:20', type: 'text', isSenderYou: true, content: 'Have a great day ahead!', senderId: 0, receiverId: 15 },
  ],
};
