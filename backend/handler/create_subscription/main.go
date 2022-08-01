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
		s := NewSubscription(e.URL, e.Title)
		
		if err != nil {
			fmt.Printf("Error Unmarshaling XML: %s", url)
			return Response{
				StatusCode: 500,
			}, nil
		}

		fmt.Printf("Created Subscription: %s", s.Title)
		return Response{
			Body: json.NewEncoder.Encode(s),
		}, nil

}

func main() {
	lambda.Start(Handler)
}