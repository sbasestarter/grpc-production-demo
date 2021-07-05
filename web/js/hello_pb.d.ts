// package: user.api
// file: hello.proto

import * as jspb from "google-protobuf";

export class HelloRequest extends jspb.Message {
  getRequest(): string;
  setRequest(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HelloRequest): HelloRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloRequest;
  static deserializeBinaryFromReader(message: HelloRequest, reader: jspb.BinaryReader): HelloRequest;
}

export namespace HelloRequest {
  export type AsObject = {
    request: string,
  }
}

export class HelloResponse extends jspb.Message {
  getResponse(): string;
  setResponse(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HelloResponse): HelloResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloResponse;
  static deserializeBinaryFromReader(message: HelloResponse, reader: jspb.BinaryReader): HelloResponse;
}

export namespace HelloResponse {
  export type AsObject = {
    response: string,
  }
}

export class HelloStreamRequest extends jspb.Message {
  getAuth(): string;
  setAuth(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloStreamRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HelloStreamRequest): HelloStreamRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloStreamRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloStreamRequest;
  static deserializeBinaryFromReader(message: HelloStreamRequest, reader: jspb.BinaryReader): HelloStreamRequest;
}

export namespace HelloStreamRequest {
  export type AsObject = {
    auth: string,
  }
}

export class HelloStreamMessage extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloStreamMessage.AsObject;
  static toObject(includeInstance: boolean, msg: HelloStreamMessage): HelloStreamMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloStreamMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloStreamMessage;
  static deserializeBinaryFromReader(message: HelloStreamMessage, reader: jspb.BinaryReader): HelloStreamMessage;
}

export namespace HelloStreamMessage {
  export type AsObject = {
    message: string,
  }
}

