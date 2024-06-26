/**
 * @fileoverview gRPC-Web generated client stub for poc.protos.cloud
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.1
// source: src/protos/main.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_any_pb = require('google-protobuf/google/protobuf/any_pb.js')
const proto = {};
proto.poc = {};
proto.poc.protos = {};
proto.poc.protos.cloud = require('./main_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.poc.protos.cloud.CloudClient =
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
proto.poc.protos.cloud.CloudPromiseClient =
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
 *   !proto.poc.protos.cloud.CloudObject,
 *   !proto.poc.protos.cloud.OperationResult>}
 */
const methodDescriptor_Cloud_Save = new grpc.web.MethodDescriptor(
  '/poc.protos.cloud.Cloud/Save',
  grpc.web.MethodType.UNARY,
  proto.poc.protos.cloud.CloudObject,
  proto.poc.protos.cloud.OperationResult,
  /**
   * @param {!proto.poc.protos.cloud.CloudObject} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.poc.protos.cloud.OperationResult.deserializeBinary
);


/**
 * @param {!proto.poc.protos.cloud.CloudObject} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.poc.protos.cloud.OperationResult)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.poc.protos.cloud.OperationResult>|undefined}
 *     The XHR Node Readable Stream
 */
proto.poc.protos.cloud.CloudClient.prototype.save =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/poc.protos.cloud.Cloud/Save',
      request,
      metadata || {},
      methodDescriptor_Cloud_Save,
      callback);
};


/**
 * @param {!proto.poc.protos.cloud.CloudObject} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.poc.protos.cloud.OperationResult>}
 *     Promise that resolves to the response
 */
proto.poc.protos.cloud.CloudPromiseClient.prototype.save =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/poc.protos.cloud.Cloud/Save',
      request,
      metadata || {},
      methodDescriptor_Cloud_Save);
};


module.exports = proto.poc.protos.cloud;

