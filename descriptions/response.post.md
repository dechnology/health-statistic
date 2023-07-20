Create a new response for the questionnaire specified by `id` path param.

## Request Body

- `user_id` *`string`* **Required**
    The user auth0 ID. This is the user who submits this response.

- `answers` *`array`* **Required**
    The answers to the questions in the questionnaire submitted by the user above.Note that the length of the answers array must equal to the number of questions in the given questionnaire. Also, all the `question_id`s must match the questions in the questionnaire.
 
    - `question_id` *`string`* **Required**
    The question id of a question in the questionnaire which this answer correspond to.

    - `body` *`string`* **Required**
    The body of the answer.
