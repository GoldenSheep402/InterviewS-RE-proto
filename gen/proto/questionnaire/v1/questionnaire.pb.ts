/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/
export type Questionnaire = {
  id?: string
  gid?: string
  description?: string
  name?: string
}

export type QuestionsList = {
  question?: Question
  options?: Option[]
}

export type Question = {
  id?: string
  questionnaireId?: string
  questionType?: string
  questionText?: string
}

export type Option = {
  id?: string
  questionId?: string
  optionText?: string
}

export type Response = {
  id?: string
  questionnaireId?: string
  userId?: string
}

export type Answer = {
  id?: string
  responseId?: string
  optionId?: string
  answerText?: string
}