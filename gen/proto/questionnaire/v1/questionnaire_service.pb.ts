/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as UserV1User from "../../user/v1/user.pb"
import * as QuestionnaireV1Questionnaire from "./questionnaire.pb"
export type CreateQuestionnaireRequest = {
  questionnaire?: QuestionnaireV1Questionnaire.Questionnaire
  questions?: QuestionnaireV1Questionnaire.QuestionsList[]
}

export type CreateQuestionnaireResponse = {
}

export type UpdateQuestionnaireRequest = {
  questionnaire?: QuestionnaireV1Questionnaire.Questionnaire
  questions?: QuestionnaireV1Questionnaire.QuestionsList[]
}

export type UpdateQuestionnaireResponse = {
}

export type GetQuestionnaireRequest = {
  questionnaireId?: string
}

export type GetQuestionnaireResponse = {
  questionnaire?: QuestionnaireV1Questionnaire.Questionnaire
  questions?: QuestionnaireV1Questionnaire.QuestionsList[]
}

export type SubmitAnswerRequest = {
  response?: QuestionnaireV1Questionnaire.Response
  answers?: QuestionnaireV1Questionnaire.Answer[]
}

export type SubmitAnswerResponse = {
}

export type GetAnswerRequest = {
  questionnaireId?: string
}

export type GetAnswerResponse = {
  user?: UserV1User.User
  answers?: QuestionnaireV1Questionnaire.Answer[]
}

export class QuestionnaireService {
  static CreateQuestionnaire(req: CreateQuestionnaireRequest, initReq?: fm.InitReq): Promise<CreateQuestionnaireResponse> {
    return fm.fetchReq<CreateQuestionnaireRequest, CreateQuestionnaireResponse>(`/gapi/questionnaire/v1/create`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateQuestionnaire(req: UpdateQuestionnaireRequest, initReq?: fm.InitReq): Promise<UpdateQuestionnaireResponse> {
    return fm.fetchReq<UpdateQuestionnaireRequest, UpdateQuestionnaireResponse>(`/gapi/v1/questionnaire/update`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetQuestionnaire(req: GetQuestionnaireRequest, initReq?: fm.InitReq): Promise<GetQuestionnaireResponse> {
    return fm.fetchReq<GetQuestionnaireRequest, GetQuestionnaireResponse>(`/gapi/v1/questionnaire/delete`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static SubmitAnswer(req: SubmitAnswerRequest, initReq?: fm.InitReq): Promise<SubmitAnswerResponse> {
    return fm.fetchReq<SubmitAnswerRequest, SubmitAnswerResponse>(`/gapi/v1/questionnaire/submit`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAnswer(req: GetAnswerRequest, initReq?: fm.InitReq): Promise<GetAnswerResponse> {
    return fm.fetchReq<GetAnswerRequest, GetAnswerResponse>(`/gapi/v1/questionnaire/Info`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}