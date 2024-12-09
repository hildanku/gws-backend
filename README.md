# gws-backend
## Endpoint

### GET /api/news

#### Description

- This endpoint fetches mental health-related articles from an external API, processes the data, and returns a structured response including headline, description, URL, and content text.

#### Request

- No request body is required for this endpoint.

#### Query Parameters

- None

#### Response

- The response contains a JSON object with the following structure:

- Success Response

Status Code: 200 OK

{
  "code": 200,
  "status": "success",
  "data": [
    {
      "headline": "string",
      "description": "string",
      "url": "string",
      "text": "string"
    }
  ]
}

#### Response Fields

code: The HTTP status code.

status: A string indicating the success of the operation.

data: An array of objects representing mental health articles.

headline: The title of the article.

description: A brief description of the article.

url: The direct URL to the article.

text: Detailed content or description extracted from the article.

