Create a new questionnaire and optionally its questions.

## Request Body

- `name` *`string`* **Required**
    The name of the questionnaire to be created.

- `questions` *`array`* Optional
    The initial questions in this questionnaire. This field may be empty and you can add questions later using post request to `quesionnaires/:id/new/question`.
    
    - `body` *`string`* **Required**
        The body of one of the question in the questionnaire.
    
    - `type` *`string`* **Required**
        The type of the question. For now, we accept all strings but in the future this field might be an enum.