const { LiveDataRequest, LiveDataReply } = require('./greeter-go/greet_go_pb.js');
const { GreeterClient } = require('./greeter-go/greet_go_grpc_web_pb.js');

var echoService = new GreeterClient('http://localhost:80');

function getLiveData() {
    echoService.echo(new LiveDataRequest(), {}, function (err, response) {
        console.log("Received: " + response.getData());
        document.getElementById("data").textContent = response.getData();
    });
}
