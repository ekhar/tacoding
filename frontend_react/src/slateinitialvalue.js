export const initialvalue = JSON.parse(localStorage.getItem("content")) || [
  {
    type: "paragraph",
    children: [{ text: "A line of text in a paragraph." }],
  },
];
