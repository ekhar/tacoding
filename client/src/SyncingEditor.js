// Import React dependencies.
import React, { useMemo, useRef, useState, useEffect } from "react";
// Import the Slate editor factory.
import { createEditor } from "slate";

// Import the Slate components and React plugin.
import { Slate, Editable, withReact } from "slate-react";
import { initialvalue } from "./slateinitialvalue";

let url = window.location.href
url = url.split("/")
url = url[url.length-1]
url = "ws://localhost:8000/ws/" + url
console.log("THIS IS THE URL", url)
var socket = new WebSocket(url);

export const SyncingEditor = () => {
  const editor = useMemo(() => withReact(createEditor()), []);
  // Add the initial value when setting up our state.
  const [value, setValue] = useState(initialvalue);
  const editor_id = useRef(String(Date.now()));
  const remote = useRef(false);

  useEffect(() => {
        console.log("I am running")
    socket.addEventListener(
      "message",
      (msg) => {
        let tmp = JSON.parse(msg.data)
        console.log("MSG DATA",msg.data)
        let id = tmp.id;
        let ops = tmp.ops;
        if (id !== editor_id.current) {
          remote.current = true;
          JSON.parse(ops).forEach((op) => {
            editor.apply(op);
          });
          remote.current = false;
        }
      },
    );
    }, []);

  return (
    <Slate
      editor={editor}
      value={value}
      onChange={(newValue) => {
        //changes value of editor
        setValue(newValue);
        //saves file locally
            //const content = JSON.stringify(newValue);
            //localStorage.setItem("content", content);
        //go through the ops to make sure they are changes to the text
        const ops = editor.operations
          .filter((o) => {
            if (o) {
              return (
                o.type !== "set_selection" &&
                o.type !== "set_value" &&
                (!o.data || (o.data instanceof Map && !o.data.has("source")))
              );
            }
            return false;
          })
          .map((o) => ({ ...o, data: { source: "one" } }));
        //if operations have been made send a message to server
        if (ops.length && !remote.current) {
        //console.log(JSON.stringify(ops))
          socket.send(
            JSON.stringify({
              id: editor_id.current,
              ops: JSON.stringify(ops),
            })
          );
        }
      }}
    >
      <Editable
        style={{
          backgroundColor: "#fafafa",
          maxWidth: 800,
          minHeight: 150,
        }}
      />
    </Slate>
  );
};
