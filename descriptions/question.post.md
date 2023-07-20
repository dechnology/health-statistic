Create a new question for the questionnaire  by the `id` path param.

## Request Body

- `body` *`string`* **Required**
    The body of one of the question in the questionnaire.

- `type` *`string`* **Required**
    The type of the question. For now, we accept all strings but in the future this field might be an enum.