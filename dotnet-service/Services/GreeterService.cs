using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;

namespace dotnet_service
{
    public class GreeterService : Greeter2.Greeter2Base
    {
        private readonly ILogger<GreeterService> _logger;
        public GreeterService(ILogger<GreeterService> logger)
        {
            _logger = logger;
        }

        public override Task<HelloReply> SayHello2(HelloRequest request, ServerCallContext context)
        {
            return Task.FromResult(new HelloReply
            {
                Msg = "Hello " + request.Msg
            });
        }
    }
}
