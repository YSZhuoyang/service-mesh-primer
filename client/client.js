
var msgEle = document.getElementById("msg");

console.log(msgEle.textContent);

var req = {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: {
        Msg: "hello"
    }
};
var resMsg = fetch("http://localhost:8000/service/1/Greet.Greeter/SayHello", req).then(res => {
    if (!res.ok) {
        console.error("Error received ...");
        return;
    }

    return res.json();
}).then(resMsg => {
    msgEle.textContent = resMsg;
});
