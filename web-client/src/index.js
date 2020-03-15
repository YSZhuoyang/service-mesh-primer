const { LiveDataRequest } = require('./greeter-go/greet_go_pb.js');
const { GreeterClient } = require('./greeter-go/greet_go_grpc_web_pb.js');

const greeterClient = new GreeterClient('http://localhost:80');

const getLiveData = () => {
    const req = new LiveDataRequest();
    const stream = greeterClient.getLiveData(req, {});
    stream.on('data', function (response) {
        console.log("Received: " + response.getData());
        document.getElementById("data").textContent = response.getData();
    });
    stream.on('status', function (status) {
        console.log("Received status code: " + status);
        if (status.metadata) {
            console.log("Received metadata");
            console.log(status.metadata);
        }
    });
    stream.on('error', function (err) {
        console.log('Error code: ' + err.code + ' "' + err.message + '"');
    });
    stream.on('end', function () {
        console.log("stream end signal received");
    });
}

document.getElementById("send").onclick = getLiveData;
