import {io} from 'socket.io-client';

export const socket = io('http://host.docker.internal:3000', {
    autoConnect: false
});
