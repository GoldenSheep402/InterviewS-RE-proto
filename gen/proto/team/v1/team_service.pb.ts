/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as UserV1User from "../../user/v1/user.pb"
import * as TeamV1Team from "./team.pb"
export type CreateTeamRequest = {
  team?: TeamV1Team.Team
}

export type CreateTeamResponse = {
}

export type UpdateTeamRequest = {
  team?: TeamV1Team.Team
}

export type UpdateTeamResponse = {
}

export type GetTeamInfoRequest = {
  teamId?: string
}

export type GetTeamInfoResponse = {
  team?: TeamV1Team.Team
  group?: TeamV1Team.Group[]
}

export type GetTeamMembersRequest = {
  teamId?: string
}

export type GetTeamMembersResponse = {
  user?: UserV1User.User[]
}

export type CreateGroupRequest = {
  teamId?: string
  group?: TeamV1Team.Group
}

export type CreateGroupResponse = {
}

export type GetGroupInfoRequest = {
  groupId?: string
}

export type GetGroupInfoResponse = {
  group?: TeamV1Team.Group
}

export type GetGroupMembersRequest = {
  groupId?: string
}

export type GetGroupMembersResponse = {
  user?: UserV1User.User[]
}

export class TeamService {
  static CreateTeam(req: CreateTeamRequest, initReq?: fm.InitReq): Promise<CreateTeamResponse> {
    return fm.fetchReq<CreateTeamRequest, CreateTeamResponse>(`/gapi/team/v1/create`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateTeam(req: UpdateTeamRequest, initReq?: fm.InitReq): Promise<UpdateTeamResponse> {
    return fm.fetchReq<UpdateTeamRequest, UpdateTeamResponse>(`/gapi/team/v1/update`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetTeamInfo(req: GetTeamInfoRequest, initReq?: fm.InitReq): Promise<GetTeamInfoResponse> {
    return fm.fetchReq<GetTeamInfoRequest, GetTeamInfoResponse>(`/gapi/team/v1/info`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetTeamMembers(req: GetTeamMembersRequest, initReq?: fm.InitReq): Promise<GetTeamMembersResponse> {
    return fm.fetchReq<GetTeamMembersRequest, GetTeamMembersResponse>(`/gapi/team/v1/members`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateGroup(req: CreateGroupRequest, initReq?: fm.InitReq): Promise<CreateGroupResponse> {
    return fm.fetchReq<CreateGroupRequest, CreateGroupResponse>(`/gapi/team/v1/group/create`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetGroupInfo(req: GetGroupInfoRequest, initReq?: fm.InitReq): Promise<GetGroupInfoResponse> {
    return fm.fetchReq<GetGroupInfoRequest, GetGroupInfoResponse>(`/gapi/team/v1/group/info`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetGroupMembers(req: GetGroupMembersRequest, initReq?: fm.InitReq): Promise<GetGroupMembersResponse> {
    return fm.fetchReq<GetGroupMembersRequest, GetGroupMembersResponse>(`/gapi/team/v1/group/members`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}