/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type RegisterRequest = {
  name?: string
  email?: string
  password?: string
}

export type RegisterResponse = {
}

export type LoginRequest = {
  email?: string
  password?: string
}

export type LoginResponse = {
  accessToken?: string
  refreshToken?: string
}

export type OauthStartRequest = {
}

export type OauthStartResponse = {
  url?: string
}

export type OauthFinishRequest = {
  code?: string
  state?: string
}

export type OauthFinishResponse = {
  accessToken?: string
  refreshToken?: string
}

export type RefreshTokenRequest = {
  refreshToken?: string
}

export type RefreshTokenResponse = {
  accessToken?: string
  refreshToken?: string
}

export type UpdateEmailLoginPasswordRequest = {
  email?: string
  newPassword?: string
}

export type UpdateEmailLoginPasswordResponse = {
}

export class AuthService {
  static Register(req: RegisterRequest, initReq?: fm.InitReq): Promise<RegisterResponse> {
    return fm.fetchReq<RegisterRequest, RegisterResponse>(`/gapi/auth/v1/register`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Login(req: LoginRequest, initReq?: fm.InitReq): Promise<LoginResponse> {
    return fm.fetchReq<LoginRequest, LoginResponse>(`/gapi/auth/v1/login`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static OauthStart(req: OauthStartRequest, initReq?: fm.InitReq): Promise<OauthStartResponse> {
    return fm.fetchReq<OauthStartRequest, OauthStartResponse>(`/gapi/auth/v1/oauth?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static OauthFinish(req: OauthFinishRequest, initReq?: fm.InitReq): Promise<OauthFinishResponse> {
    return fm.fetchReq<OauthFinishRequest, OauthFinishResponse>(`/gapi/auth/v1/oauth/callback?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static RefreshToken(req: RefreshTokenRequest, initReq?: fm.InitReq): Promise<RefreshTokenResponse> {
    return fm.fetchReq<RefreshTokenRequest, RefreshTokenResponse>(`/gapi/auth/v1/refreshToken`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateEmailLoginPassword(req: UpdateEmailLoginPasswordRequest, initReq?: fm.InitReq): Promise<UpdateEmailLoginPasswordResponse> {
    return fm.fetchReq<UpdateEmailLoginPasswordRequest, UpdateEmailLoginPasswordResponse>(`/gapi/auth/v1/updatePassword/email`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}