// WebSocketComponent.js
import React, { useEffect, useState } from 'react';

const WebSocketComponent = () => {
  const [websocket, setWebsocket] = useState(null);
  const [message, setMessage] = useState({ username: '', message: '' });
  const [inputMessage, setInputMessage] = useState('');

  useEffect(() => {
    // Connect to the WebSocket server
    const ws = new WebSocket('ws://localhost:8080/echo');

    // Set up event listeners
    ws.onopen = () => {
      console.log('WebSocket connected');
    };

    ws.onmessage = (event) => {
      const receivedMessage = JSON.parse(event.data);
      console.log("Message from server", receivedMessage)
      setMessage(receivedMessage);
    };

    ws.onclose = () => {
      console.log('WebSocket closed');
    };

    // Save the WebSocket connection in state
    setWebsocket(ws);

    // Clean up on component unmount
    return () => {
        console.log('Closing WebSocket connection...');
        ws.close();
    };
  }, []); // Empty dependency array ensures this effect runs once on component mount

  const sendMessage = () => {
    // Send a JSON message to the server
    const messageObject = { text: inputMessage };
    websocket.send(JSON.stringify(messageObject));
    setInputMessage('');
  };

  return (
    <div>
      <h1>WebSocket Chat</h1>
      <div>
        <p>
          Received message from {message.username}: {message.message}
        </p>
      </div>
      <div>
        <input
          type="text"
          value={inputMessage}
          onChange={(e) => setInputMessage(e.target.value)}
        />
        <button onClick={sendMessage}>Send Message</button>
        <p>Received message from server: {message.text}</p>
      </div>
    </div>
  );
};

export default WebSocketComponent;
