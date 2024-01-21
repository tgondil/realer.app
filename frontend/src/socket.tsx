import { io } from 'socket.io-client';

// "undefined" means the URL will be computed from the `window.location` object
const URL = 'http://44.221.67.84:4000';

export const socket = io(URL, {
    autoConnect: false
});