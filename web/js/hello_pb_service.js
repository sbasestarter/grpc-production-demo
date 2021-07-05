/* eslint-disable */
// package: user.api
// file: hello.proto

var hello_pb = require("./hello_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Hellos = (function () {
  function Hellos() {}
  Hellos.serviceName = "user.api.Hellos";
  return Hellos;
}());

Hellos.SayHello = {
  methodName: "SayHello",
  service: Hellos,
  requestStream: false,
  responseStream: false,
  requestType: hello_pb.HelloRequest,
  responseType: hello_pb.HelloResponse
};

Hellos.HelloStream = {
  methodName: "HelloStream",
  service: Hellos,
  requestStream: false,
  responseStream: true,
  requestType: hello_pb.HelloStreamRequest,
  responseType: hello_pb.HelloStreamMessage
};

exports.Hellos = Hellos;

function HellosClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

HellosClient.prototype.sayHello = function sayHello(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Hellos.SayHello, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

HellosClient.prototype.helloStream = function helloStream(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Hellos.HelloStream, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.HellosClient = HellosClient;

