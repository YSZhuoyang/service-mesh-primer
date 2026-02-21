import { LiveDataRequest } from './greeter-go/greet_go_pb';
import { GreeterClient } from './greeter-go/greet_go_grpc_web_pb';

const greeterClient = new GreeterClient('http://' + window.location.hostname + ':8080');

console.log("URL: " + window.location.hostname)

const getLiveData = () => {
    const req = new LiveDataRequest();
    const stream = greeterClient.getLiveData(req, {});
    stream.on('data', function (response) {
        console.log("Received: " + response.getData());
        document.getElementById("data").textContent = response.getData();
    });
    stream.on('status', function (status) {
        console.log("Received status code: " + status.code);
        console.log("Received status details: " + status.details);
        console.log("Received status metadata: " + status.metadata);
    });
    stream.on('error', function (err) {
        console.log('Error code: ' + err.code + ' "' + err.message + '"');
    });
    stream.on('end', function () {
        console.log("stream end signal received");
    });
}

document.getElementById("send").onclick = getLiveData;
