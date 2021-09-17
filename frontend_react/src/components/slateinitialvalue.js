export const initialvalue = JSON.parse(localStorage.getItem("content")) || [
  {
    type: "paragraph",
    children: [{ text: "print('hello world')" }],
  },
];
