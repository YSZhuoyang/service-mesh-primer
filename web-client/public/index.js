const { LiveDataRequest, LiveDataReply } = require('./live_pb.js');
const { GreetServiceClient } = require('./live_grpc_web_pb.js');

var echoService = new GreetServiceClient('http://localhost:80');

function getLiveData() {
    echoService.echo(new LiveDataRequest(), {}, function (err, response) {
        // ...
    });
}
