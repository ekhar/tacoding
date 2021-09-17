import Button from "@restart/ui/esm/Button";
export const ButtonCodeRun = () => {
  var send_code = () => {
    let plaintext = localStorage.getItem("plaintext");
    fetch("http://localhost:8000/execute", {
      method: "post",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        "Access-Control-Allow-Origin": "*",
      },
      body: "A=1&B=2",
    })
      .then(function (response) {
        console.log("Authentication Success");
      })
      .catch(function (err) {
        console.log("Authentication fail", err);
      });
    console.log("click");
  };
  return <Button onClick={send_code}>Run Code</Button>;
};
