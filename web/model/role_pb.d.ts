import * as jspb from 'google-protobuf'

import * as google_protobuf_descriptor_pb from 'google-protobuf/google/protobuf/descriptor_pb';


export class Role extends jspb.Message {
  getProjectId(): string;
  setProjectId(value: string): Role;

  getProjectRole(): Role.ProjectRole;
  setProjectRole(value: Role.ProjectRole): Role;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Role.AsObject;
  static toObject(includeInstance: boolean, msg: Role): Role.AsObject;
  static serializeBinaryToWriter(message: Role, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Role;
  static deserializeBinaryFromReader(message: Role, reader: jspb.BinaryReader): Role;
}

export namespace Role {
  export type AsObject = {
    projectId: string,
    projectRole: Role.ProjectRole,
  }

  export enum ProjectRole { 
    VIEWER = 0,
    EDITOR = 1,
    ADMIN = 2,
  }
}

