/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as UserV1User from "./user.pb"
export type GetUserInfoRequest = {
  id?: string
}

export type GetUserInfoResponse = {
  info?: UserV1User.User
}

export type UpdateUserInfoRequest = {
  info?: UserV1User.User
}

export type UpdateUserInfoResponse = {
}

export class UserService {
  static GetUserInfo(req: GetUserInfoRequest, initReq?: fm.InitReq): Promise<GetUserInfoResponse> {
    return fm.fetchReq<GetUserInfoRequest, GetUserInfoResponse>(`/gapi/user/v1/info`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateUserInfo(req: UpdateUserInfoRequest, initReq?: fm.InitReq): Promise<UpdateUserInfoResponse> {
    return fm.fetchReq<UpdateUserInfoRequest, UpdateUserInfoResponse>(`/gapi/user/v1/edit`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}