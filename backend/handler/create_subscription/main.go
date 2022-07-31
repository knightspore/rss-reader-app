package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	subscription "github.com/knightspore/rss-reader-app/backend/module/Subscription"
)

type Response events.APIGatewayProxyResponse

type Event struct {
	URL string `json:"url"`
}

func Handler(request events.APIGatewayProxyRequest) (Response, error) {

		e := Event{URL: request.Body}
		s, err := subscription.Create(e.URL)

		if err != nil {
			return Response{
				StatusCode: 500,
			}, nil
		}

		fmt.Printf("Created Subscription: %s", s.Title)

		return Response{
			StatusCode: 200,
		}, nil
}

func main() {
	lambda.Start(Handler)
}