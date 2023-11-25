/**
 * @fileoverview gRPC-Web generated client stub for go_service
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v4.25.1
// source: greeter-go/greet_go.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


const proto = {};
proto.go_service = require('./greet_go_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.go_service.GreeterClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.go_service.GreeterPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.go_service.HelloRequest,
 *   !proto.go_service.HelloReply>}
 */
const methodDescriptor_Greeter_SayHello = new grpc.web.MethodDescriptor(
  '/go_service.Greeter/SayHello',
  grpc.web.MethodType.UNARY,
  proto.go_service.HelloRequest,
  proto.go_service.HelloReply,
  /**
   * @param {!proto.go_service.HelloRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.go_service.HelloReply.deserializeBinary
);


/**
 * @param {!proto.go_service.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.go_service.HelloReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.go_service.HelloReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.go_service.GreeterClient.prototype.sayHello =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/go_service.Greeter/SayHello',
      request,
      metadata || {},
      methodDescriptor_Greeter_SayHello,
      callback);
};


/**
 * @param {!proto.go_service.HelloRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.go_service.HelloReply>}
 *     Promise that resolves to the response
 */
proto.go_service.GreeterPromiseClient.prototype.sayHello =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/go_service.Greeter/SayHello',
      request,
      metadata || {},
      methodDescriptor_Greeter_SayHello);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.go_service.LiveDataRequest,
 *   !proto.go_service.LiveDataReply>}
 */
const methodDescriptor_Greeter_GetLiveData = new grpc.web.MethodDescriptor(
  '/go_service.Greeter/GetLiveData',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.go_service.LiveDataRequest,
  proto.go_service.LiveDataReply,
  /**
   * @param {!proto.go_service.LiveDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.go_service.LiveDataReply.deserializeBinary
);


/**
 * @param {!proto.go_service.LiveDataRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.go_service.LiveDataReply>}
 *     The XHR Node Readable Stream
 */
proto.go_service.GreeterClient.prototype.getLiveData =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/go_service.Greeter/GetLiveData',
      request,
      metadata || {},
      methodDescriptor_Greeter_GetLiveData);
};


/**
 * @param {!proto.go_service.LiveDataRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.go_service.LiveDataReply>}
 *     The XHR Node Readable Stream
 */
proto.go_service.GreeterPromiseClient.prototype.getLiveData =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/go_service.Greeter/GetLiveData',
      request,
      metadata || {},
      methodDescriptor_Greeter_GetLiveData);
};


module.exports = proto.go_service;

