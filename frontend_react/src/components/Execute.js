import { useState } from "react";

var socket = new WebSocket("ws://localhost:8000/ws");
export const Execute = () => {
  const [executed, setExecuted] = useState(
    "Once run, code will be outputted here..."
  );
  socket.addEventListener("message", (data) => {
    let action = data.action;
    if (action === "execute") {
      setExecuted(data.output);
    }
  });
  return <p>{executed}</p>; //<Button variant="contained">Default</Button>;
};
