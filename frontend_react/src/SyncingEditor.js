// Import React dependencies.
import React, { useMemo, useRef, useState, useEffect } from "react";
// Import the Slate editor factory.
import { createEditor } from "slate";
import Mitt from "mitt";

// Import the Slate components and React plugin.
import { Slate, Editable, withReact } from "slate-react";
import { initialvalue } from "./slateinitialvalue";

const emitter = new Mitt();
export const SyncingEditor = () => {
  const editor = useMemo(() => withReact(createEditor()), []);
  // Add the initial value when setting up our state.
  const [value, setValue] = useState(initialvalue);
  const id = useRef(String(Date.now()));
  const remote = useRef(false);
  const editor_ref = useRef(editor);

  useEffect(() => {
    emitter.on("*", (type, ops) => {
      console.log("emmiter running");
      if (id.current !== type) {
        remote.current = true;
        console.log(ops);
        ops.forEach((op) => {
          editor.apply(op);
        });
        remote.current = false;
      }
    });
  });

  return (
    <Slate
      editor={editor}
      value={value}
      onChange={(newValue) => {
        setValue(newValue);
        const ops = editor_ref.current.operations
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
          console.log("OP are this long" + ops.length);
          emitter.emit(id.current, ops);
          emitter.all.clear();
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
