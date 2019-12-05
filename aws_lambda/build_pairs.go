package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devingen/atama-api/atama"
	"github.com/devingen/atama-api/dto"
	"log"
	"time"
)

func HandleBuildPairsRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body dto.BuildPairsBody
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid payload",
			Headers: map[string]string{
				//"Content-Type": "application/json",
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	start := time.Now()
	log.Printf("Received build pairs request")
	log.Printf("%d %d", len(body.List1), len(body.List2))

	scoreMatrix := atama.GenerateScoreMatrix(body.Rules, body.List1, body.List2, body.List1Fields, body.List2Fields)
	log.Printf("GenerateScoreMatrix: %s", time.Since(start))

	result := atama.CalculateList(scoreMatrix, nil, 0, false)
	log.Printf("CalculateList: %s", time.Since(start))

	response, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	log.Printf("Response is built")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(HandleBuildPairsRequest)
}
