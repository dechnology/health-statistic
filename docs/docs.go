// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health_check": {
            "get": {
                "description": "A health checking endpoint to make sure the server is not dead.",
                "tags": [
                    "Health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.HealthStatus"
                        }
                    }
                }
            }
        },
        "/questionnaires": {
            "get": {
                "description": "Get all questionnaires from the database. For each questionnaire, all questions and responses of this questionnaire are also included.\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Get Questionnaires",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Questionnaire"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new questionnaire and optionally its questions.\r\n\r\n## Request Body\r\n\r\n- ` + "`" + `name` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The name of the questionnaire to be created.\r\n\r\n- ` + "`" + `questions` + "`" + ` *` + "`" + `array` + "`" + `* Optional\r\n    The initial questions in this questionnaire. This field may be empty and you can add questions later using post request to ` + "`" + `quesionnaires/:id/new/question` + "`" + `.\r\n    \r\n    - ` + "`" + `body` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n        The body of one of the question in the questionnaire.\r\n    \r\n    - ` + "`" + `type` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n        The type of the question. For now, we accept all strings but in the future this field might be an enum.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Questionnaire",
                "parameters": [
                    {
                        "description": "The questionnaire to be created.",
                        "name": "questionnaire",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SingleQuestionnaireResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Questionnaire"
                        }
                    }
                }
            }
        },
        "/questionnaires/{id}": {
            "get": {
                "description": "Get the questionnaire specified by the ` + "`" + `id` + "`" + ` path param. All questions and responses of the questionnaire are also included.\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Get Questionnaire",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Questionnaire"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the questionnaire specified by the ` + "`" + `id` + "`" + ` path param.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Delete Questionnaire",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/questionnaires/{id}/new/question": {
            "post": {
                "description": "Create a new question for the questionnaire  by the ` + "`" + `id` + "`" + ` path param.\r\n\r\n## Request Body\r\n\r\n- ` + "`" + `body` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The body of one of the question in the questionnaire.\r\n\r\n- ` + "`" + `type` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The type of the question. For now, we accept all strings but in the future this field might be an enum.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The question to be created.",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Question"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ent.Question"
                        }
                    }
                }
            }
        },
        "/questionnaires/{id}/new/response": {
            "post": {
                "description": "Create a new response for the questionnaire specified by ` + "`" + `id` + "`" + ` path param.\r\n\r\n## Request Body\r\n\r\n- ` + "`" + `user_id` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The user auth0 ID. This is the user who submits this response.\r\n\r\n- ` + "`" + `answers` + "`" + ` *` + "`" + `array` + "`" + `* **Required**\r\n    The answers to the questions in the questionnaire submitted by the user above.Note that the length of the answers array must equal to the number of questions in the given questionnaire. Also, all the ` + "`" + `question_id` + "`" + `s must match the questions in the questionnaire.\r\n \r\n    - ` + "`" + `question_id` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The question id of a question in the questionnaire which this answer correspond to.\r\n\r\n    - ` + "`" + `body` + "`" + ` *` + "`" + `string` + "`" + `* **Required**\r\n    The body of the answer.\r\n",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Questionnaire"
                ],
                "summary": "Create Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The questionnaire's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The response to be created.",
                        "name": "response",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireResponseBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.QuestionnaireResponse"
                        }
                    }
                }
            }
        },
        "/questions": {
            "get": {
                "description": "Get all questions from the database. **This will NOT include questionnaires and responses.**\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get Questions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.Question"
                            }
                        }
                    }
                }
            }
        },
        "/questions/{id}": {
            "get": {
                "description": "Get the question specified by the ` + "`" + `id` + "`" + ` path param.\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The question's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Question"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the question specified by the ` + "`" + `id` + "`" + ` path param.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Delete Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The question's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/responses": {
            "get": {
                "description": "Get all responses from the database. **This will NOT include questionnaires and questions.**\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Get Responses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.QuestionnaireResponse"
                            }
                        }
                    }
                }
            }
        },
        "/responses/{id}": {
            "get": {
                "description": "Get the response specified by the ` + "`" + `id` + "`" + ` path param. **This will also includes the corresponding questionnaire.**\r\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Get Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The response's ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SingleQuestionnaireResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the response specified by the ` + "`" + `id` + "`" + ` path param.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Response"
                ],
                "summary": "Delete Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The response's ID.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get all users from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "The user to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get a user by an Auth0 ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            },
            "put": {
                "description": "update a user with specified Auth0 ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user to be updated",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by Auth0 ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user's Auth0 ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Answer": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "ent.Question": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "type": "string"
                }
            }
        },
        "ent.Questionnaire": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "birth_year": {
                    "description": "BirthYear holds the value of the \"birth_year\" field.",
                    "type": "integer"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "demented_among_direct_relatives": {
                    "description": "DementedAmongDirectRelatives holds the value of the \"demented_among_direct_relatives\" field.",
                    "type": "boolean"
                },
                "ear_condition": {
                    "description": "EarCondition holds the value of the \"ear_condition\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EarCondition"
                        }
                    ]
                },
                "education_level": {
                    "description": "EducationLevel holds the value of the \"education_level\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EducationLevel"
                        }
                    ]
                },
                "eyesight_condition": {
                    "description": "EyesightCondition holds the value of the \"eyesight_condition\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.EyesightCondition"
                        }
                    ]
                },
                "first_name": {
                    "description": "FirstName holds the value of the \"first_name\" field.",
                    "type": "string"
                },
                "gender": {
                    "description": "Gender holds the value of the \"gender\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Gender"
                        }
                    ]
                },
                "head_injury_experience": {
                    "description": "HeadInjuryExperience holds the value of the \"head_injury_experience\" field.",
                    "type": "boolean"
                },
                "height": {
                    "description": "Height holds the value of the \"height\" field.",
                    "type": "number"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "last_name": {
                    "description": "LastName holds the value of the \"last_name\" field.",
                    "type": "string"
                },
                "marriage": {
                    "description": "Marriage holds the value of the \"marriage\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Marriage"
                        }
                    ]
                },
                "medical_history": {
                    "description": "MedicalHistory holds the value of the \"medical_history\" field.",
                    "type": "string"
                },
                "medication_status": {
                    "description": "MedicationStatus holds the value of the \"medication_status\" field.",
                    "type": "string"
                },
                "occupation": {
                    "description": "Occupation holds the value of the \"occupation\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Occupation"
                        }
                    ]
                },
                "smoking_habit": {
                    "description": "SmokingHabit holds the value of the \"smoking_habit\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.SmokingHabit"
                        }
                    ]
                },
                "updated_at": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                },
                "weight": {
                    "description": "Weight holds the value of the \"weight\" field.",
                    "type": "number"
                }
            }
        },
        "handlers.AnswerBody": {
            "description": "The json body for creating a new answer in a new response.",
            "type": "object",
            "properties": {
                "body": {
                    "description": "The answer body.",
                    "type": "string"
                },
                "question_id": {
                    "description": "The question this answer relates to, the question also needs to be in\nthe same questionnaire as the response.",
                    "type": "string"
                }
            }
        },
        "handlers.HealthStatus": {
            "description": "Datatype of health status",
            "type": "object",
            "properties": {
                "message": {
                    "description": "Health message",
                    "type": "string"
                }
            }
        },
        "handlers.Question": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "quesionnaires": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Questionnaire"
                    }
                },
                "type": {
                    "description": "Type holds the value of the \"type\" field.",
                    "type": "string"
                }
            }
        },
        "handlers.QuestionBody": {
            "description": "The json body for creating a new question.",
            "type": "object",
            "properties": {
                "body": {
                    "description": "The question body",
                    "type": "string"
                },
                "type": {
                    "description": "The question type, currently we accept string but in the future this\nfield will be enums.",
                    "type": "string"
                }
            }
        },
        "handlers.Questionnaire": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Question"
                    }
                },
                "responses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.QuestionnaireResponse"
                    }
                }
            }
        },
        "handlers.QuestionnaireResponse": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Answer"
                    }
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "handlers.QuestionnaireResponseBody": {
            "description": "The json body for creating a new response.",
            "type": "object",
            "properties": {
                "answers": {
                    "description": "The answers to all questions in a questionnaire.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.AnswerBody"
                    }
                },
                "user_id": {
                    "description": "The user ID of the user who submit the response.",
                    "type": "string"
                }
            }
        },
        "handlers.SingleQuestionnaireResponse": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Answer"
                    }
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "questionnaire": {
                    "$ref": "#/definitions/handlers.Questionnaire"
                }
            }
        },
        "user.EarCondition": {
            "type": "string",
            "enum": [
                "normal",
                "slightly_affecting_conversation",
                "need_hearing_aid"
            ],
            "x-enum-varnames": [
                "EarConditionNormal",
                "EarConditionSlightlyAffectingConversation",
                "EarConditionNeedHearingAid"
            ]
        },
        "user.EducationLevel": {
            "type": "string",
            "enum": [
                "elementry_school_or_below",
                "middle_school",
                "high_school",
                "bachelor",
                "master",
                "doctorate"
            ],
            "x-enum-varnames": [
                "EducationLevelElementrySchoolOrBelow",
                "EducationLevelMiddleSchool",
                "EducationLevelHighSchool",
                "EducationLevelBachelor",
                "EducationLevelMaster",
                "EducationLevelDoctorate"
            ]
        },
        "user.EyesightCondition": {
            "type": "string",
            "enum": [
                "normal",
                "slightly_affecting_reading",
                "need_glasses"
            ],
            "x-enum-varnames": [
                "EyesightConditionNormal",
                "EyesightConditionSlightlyAffectingReading",
                "EyesightConditionNeedGlasses"
            ]
        },
        "user.Gender": {
            "type": "string",
            "enum": [
                "male",
                "female",
                "nonbinary"
            ],
            "x-enum-varnames": [
                "GenderMale",
                "GenderFemale",
                "GenderNonbinary"
            ]
        },
        "user.Marriage": {
            "type": "string",
            "enum": [
                "single",
                "married",
                "divorced",
                "widowed"
            ],
            "x-enum-varnames": [
                "MarriageSingle",
                "MarriageMarried",
                "MarriageDivorced",
                "MarriageWidowed"
            ]
        },
        "user.Occupation": {
            "type": "string",
            "enum": [
                "student",
                "government_employee",
                "service_industry",
                "industry_and_commerce",
                "freelancer",
                "domestic"
            ],
            "x-enum-varnames": [
                "OccupationStudent",
                "OccupationGovernmentEmployee",
                "OccupationServiceIndustry",
                "OccupationIndustryAndCommerce",
                "OccupationFreelancer",
                "OccupationDomestic"
            ]
        },
        "user.SmokingHabit": {
            "type": "string",
            "enum": [
                "none",
                "sometimes",
                "everyday"
            ],
            "x-enum-varnames": [
                "SmokingHabitNone",
                "SmokingHabitSometimes",
                "SmokingHabitEveryday"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "health-statistic.dechnology.com.tw",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Web3 - 健康資料公鏈API開發文件",
	Description:      "This is the API documentation for 「健康資料公鏈」.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}