import { LiveDataRequest } from './greeter-go/greet_go_pb';
import { GreeterClient } from './greeter-go/greet_go_grpc_web_pb';

// Change "127.0.0.1" to "host.docker.internal" to run this script from vscode dev container
// Note: chrome does not accept self signed certificates for localhost, ERR_CERT_AUTHORITY_INVALID error will occur
const greeterClient = new GreeterClient('https://127.0.0.1:443');

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
