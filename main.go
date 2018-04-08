package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

//IP is just a tring but called IP for XML conversion
type IP string

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TO DO ADD CORS
	response := getIP(request.Headers["X-Forwarded-For"])
	contentType := "text/plain; charset=UTF-8"

	if format, hasFormat := request.QueryStringParameters["format"]; hasFormat {
		if format == "json" {
			contentType = "application/json; charset=UTF-8"
			out, _ := json.Marshal(map[string]string{"ip": response})
			response = string(out)
		} else if format == "xml" {
			// return c.XML(http.StatusOK, IP(c.RealIP()))
		} else if callback, hasCallback := request.QueryStringParameters["callback"]; format == "jsonp" && hasCallback {
			response = callback // TO DO: make real
		}

	}

	return events.APIGatewayProxyResponse{Body: response, StatusCode: 200, Headers: map[string]string{"Content-Type": contentType}}, nil
}

func getIP(in string) string {
	return strings.TrimSpace(strings.Split(in, ",")[0])
}
