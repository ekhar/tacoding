// Import React dependencies.
import React, { useMemo, useRef, useState, useEffect } from "react";
// Import the Slate editor factory.
import { createEditor } from "slate";
import { Node } from "slate";
// Import the Slate components and React plugin.
import { Slate, Editable, withReact } from "slate-react";
import { initialvalue } from "./slateinitialvalue";
var socket = new WebSocket("ws://localhost:8000/ws");

export const SyncingEditor = () => {
  const editor = useMemo(() => withReact(createEditor()), []);
  // Add the initial value when setting up our state.
  const [value, setValue] = useState(initialvalue);
  const id = useRef(String(Date.now()));
  const remote = useRef(false);

  useEffect(() => {
    socket.addEventListener(
      "message",
      (data) => {
        var x = JSON.parse(JSON.parse(data.data).body);
        let action = x.action;
        if (action === "typing") {
          let editor_id = x.editor_id;
          let ops = x.ops;

          if (id.current !== editor_id) {
            remote.current = true;
            JSON.parse(ops).forEach((op) => {
              editor.apply(op);
            });
            remote.current = false;
          }
        }
      },
      { once: true }
    );
  });

  return (
    <Slate
      editor={editor}
      value={value}
      onChange={(newValue) => {
        //changes value of editor
        setValue(newValue);
        //saves file locally
        const serialize = (nodes) => {
          return nodes.map((n) => Node.string(n)).join("\n");
        };
        const content = JSON.stringify(newValue);
        localStorage.setItem("content", content);
        localStorage.setItem("plaintext", serialize(newValue));
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
        if (ops.length && !remote.current) {
          socket.send(
            JSON.stringify({
              action: "typing",
              editor_id: id.current,
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