import React, { useEffect } from "react";

import logo from "./logo.svg";
import "./App.css";
import { connect, sendMsg } from "./api";

function App() {
  const handleSend = () => {
    console.log("hello");
    sendMsg("Hello");
  };

  useEffect(() => {
    // Connect to socket.
    connect();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />

        <button type="button" onClick={handleSend}>
          Hit
        </button>
      </header>
    </div>
  );
}

export default App;
