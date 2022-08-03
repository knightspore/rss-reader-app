package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

type Response events.APIGatewayProxyResponse

type Event struct {
	URL string `json:"url"`
	Title string `json:"title"`
}

func Handler(request events.APIGatewayProxyRequest) (Response, error) {

		e := Event{URL: request.Body.URL, Title: request.Body.Title}
		s, err := vo.NewSubscription(e.URL, e.Title)
		
		if err != nil {
			fmt.Printf("Error Unmarshaling XML: %s", request.Body.URL)
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