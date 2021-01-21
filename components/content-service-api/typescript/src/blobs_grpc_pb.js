// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.
//
'use strict';
var grpc = require('grpc');
var blobs_pb = require('./blobs_pb.js');

function serialize_contentservice_DeleteRequest(arg) {
  if (!(arg instanceof blobs_pb.DeleteRequest)) {
    throw new Error('Expected argument of type contentservice.DeleteRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_DeleteRequest(buffer_arg) {
  return blobs_pb.DeleteRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contentservice_DeleteResponse(arg) {
  if (!(arg instanceof blobs_pb.DeleteResponse)) {
    throw new Error('Expected argument of type contentservice.DeleteResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_DeleteResponse(buffer_arg) {
  return blobs_pb.DeleteResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contentservice_DownloadUrlRequest(arg) {
  if (!(arg instanceof blobs_pb.DownloadUrlRequest)) {
    throw new Error('Expected argument of type contentservice.DownloadUrlRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_DownloadUrlRequest(buffer_arg) {
  return blobs_pb.DownloadUrlRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contentservice_DownloadUrlResponse(arg) {
  if (!(arg instanceof blobs_pb.DownloadUrlResponse)) {
    throw new Error('Expected argument of type contentservice.DownloadUrlResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_DownloadUrlResponse(buffer_arg) {
  return blobs_pb.DownloadUrlResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contentservice_UploadUrlRequest(arg) {
  if (!(arg instanceof blobs_pb.UploadUrlRequest)) {
    throw new Error('Expected argument of type contentservice.UploadUrlRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_UploadUrlRequest(buffer_arg) {
  return blobs_pb.UploadUrlRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_contentservice_UploadUrlResponse(arg) {
  if (!(arg instanceof blobs_pb.UploadUrlResponse)) {
    throw new Error('Expected argument of type contentservice.UploadUrlResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_contentservice_UploadUrlResponse(buffer_arg) {
  return blobs_pb.UploadUrlResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var BlobServiceService = exports.BlobServiceService = {
  // UploadUrl provides a URL to which clients can upload the content via HTTP
// PUT. Possible errors:
// - INVALID_ARGUMENT if owner_id or name do not match `[a-zA-Z0-9._\-\/]+`
// - RESOURCE_EXHAUSTED if a user's upload area is exhausted
uploadUrl: {
    path: '/contentservice.BlobService/UploadUrl',
    requestStream: false,
    responseStream: false,
    requestType: blobs_pb.UploadUrlRequest,
    responseType: blobs_pb.UploadUrlResponse,
    requestSerialize: serialize_contentservice_UploadUrlRequest,
    requestDeserialize: deserialize_contentservice_UploadUrlRequest,
    responseSerialize: serialize_contentservice_UploadUrlResponse,
    responseDeserialize: deserialize_contentservice_UploadUrlResponse,
  },
  // DownloadUrl provides a URL from which clients download the content via HTTP.
// Expect the download to be resumable, i.e. support for the `Range` header.
// Possible errors:
// - INVALID_ARGUMENT if owner_id or name do not match `[a-zA-Z0-9._\-\/]+`
// - NOT_FOUND if the blob does not exist
downloadUrl: {
    path: '/contentservice.BlobService/DownloadUrl',
    requestStream: false,
    responseStream: false,
    requestType: blobs_pb.DownloadUrlRequest,
    responseType: blobs_pb.DownloadUrlResponse,
    requestSerialize: serialize_contentservice_DownloadUrlRequest,
    requestDeserialize: deserialize_contentservice_DownloadUrlRequest,
    responseSerialize: serialize_contentservice_DownloadUrlResponse,
    responseDeserialize: deserialize_contentservice_DownloadUrlResponse,
  },
  // Delete removes blob content from the store immediately.
// This renders previously provided up-/download URLs invalid.
// Ongoing up-/downloads may or may not succeed.
// Possible errors:
// - INVALID_ARGUMENT if owner_id or name do not match `[a-zA-Z0-9._\-\/]+`
// - NOT_FOUND if the blob does not exist
delete: {
    path: '/contentservice.BlobService/Delete',
    requestStream: false,
    responseStream: false,
    requestType: blobs_pb.DeleteRequest,
    responseType: blobs_pb.DeleteResponse,
    requestSerialize: serialize_contentservice_DeleteRequest,
    requestDeserialize: deserialize_contentservice_DeleteRequest,
    responseSerialize: serialize_contentservice_DeleteResponse,
    responseDeserialize: deserialize_contentservice_DeleteResponse,
  },
};

exports.BlobServiceClient = grpc.makeGenericClientConstructor(BlobServiceService);
