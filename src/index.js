import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import axios from "axios";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

console.log("posting test to /testweb")
axios.post("/testweb")
  .then((res) => console.log("received:", res.data))
  .catch((err) => console.error(err));

console.log("opening ws to /testws")
let ws = new WebSocket("ws://" + window.location.host + "/testws");
ws.addEventListener('open', function (event) {
  console.log("ws open")
  ws.send("echo!")
});
ws.addEventListener('message', function (event) {
  console.log('Message from server ', event.data);
});
