// package: contentservice
// file: blobs.proto

/* tslint:disable */

import * as jspb from "google-protobuf";

export class BlobRef extends jspb.Message { 
    getOwnerId(): string;
    setOwnerId(value: string): void;

    getName(): string;
    setName(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BlobRef.AsObject;
    static toObject(includeInstance: boolean, msg: BlobRef): BlobRef.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BlobRef, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BlobRef;
    static deserializeBinaryFromReader(message: BlobRef, reader: jspb.BinaryReader): BlobRef;
}

export namespace BlobRef {
    export type AsObject = {
        ownerId: string,
        name: string,
    }
}

export class UploadUrlRequest extends jspb.Message { 

    hasRef(): boolean;
    clearRef(): void;
    getRef(): BlobRef | undefined;
    setRef(value?: BlobRef): void;

    getTimeToLive(): string;
    setTimeToLive(value: string): void;

    getMediaType(): string;
    setMediaType(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadUrlRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UploadUrlRequest): UploadUrlRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadUrlRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadUrlRequest;
    static deserializeBinaryFromReader(message: UploadUrlRequest, reader: jspb.BinaryReader): UploadUrlRequest;
}

export namespace UploadUrlRequest {
    export type AsObject = {
        ref?: BlobRef.AsObject,
        timeToLive: string,
        mediaType: string,
    }
}

export class UploadUrlResponse extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): void;

    getCommitToken(): string;
    setCommitToken(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UploadUrlResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UploadUrlResponse): UploadUrlResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UploadUrlResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UploadUrlResponse;
    static deserializeBinaryFromReader(message: UploadUrlResponse, reader: jspb.BinaryReader): UploadUrlResponse;
}

export namespace UploadUrlResponse {
    export type AsObject = {
        url: string,
        commitToken: string,
    }
}

export class CommitRequest extends jspb.Message { 

    hasRef(): boolean;
    clearRef(): void;
    getRef(): BlobRef | undefined;
    setRef(value?: BlobRef): void;

    getCommitToken(): string;
    setCommitToken(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CommitRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CommitRequest): CommitRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CommitRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CommitRequest;
    static deserializeBinaryFromReader(message: CommitRequest, reader: jspb.BinaryReader): CommitRequest;
}

export namespace CommitRequest {
    export type AsObject = {
        ref?: BlobRef.AsObject,
        commitToken: string,
    }
}

export class CommitResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CommitResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CommitResponse): CommitResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CommitResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CommitResponse;
    static deserializeBinaryFromReader(message: CommitResponse, reader: jspb.BinaryReader): CommitResponse;
}

export namespace CommitResponse {
    export type AsObject = {
    }
}

export class DownloadUrlRequest extends jspb.Message { 

    hasRef(): boolean;
    clearRef(): void;
    getRef(): BlobRef | undefined;
    setRef(value?: BlobRef): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DownloadUrlRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DownloadUrlRequest): DownloadUrlRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DownloadUrlRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DownloadUrlRequest;
    static deserializeBinaryFromReader(message: DownloadUrlRequest, reader: jspb.BinaryReader): DownloadUrlRequest;
}

export namespace DownloadUrlRequest {
    export type AsObject = {
        ref?: BlobRef.AsObject,
    }
}

export class DownloadUrlResponse extends jspb.Message { 
    getUrl(): string;
    setUrl(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DownloadUrlResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DownloadUrlResponse): DownloadUrlResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DownloadUrlResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DownloadUrlResponse;
    static deserializeBinaryFromReader(message: DownloadUrlResponse, reader: jspb.BinaryReader): DownloadUrlResponse;
}

export namespace DownloadUrlResponse {
    export type AsObject = {
        url: string,
    }
}

export class DeleteRequest extends jspb.Message { 

    hasRef(): boolean;
    clearRef(): void;
    getRef(): BlobRef | undefined;
    setRef(value?: BlobRef): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRequest;
    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
    export type AsObject = {
        ref?: BlobRef.AsObject,
    }
}

export class DeleteResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteResponse): DeleteResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteResponse;
    static deserializeBinaryFromReader(message: DeleteResponse, reader: jspb.BinaryReader): DeleteResponse;
}

export namespace DeleteResponse {
    export type AsObject = {
    }
}
