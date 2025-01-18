package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nexryai/summergo"
	"net/url"
)

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// クエリパラメータを取得
	queryParams, err := url.ParseQuery(request.RawQueryString)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf(`{"error": "Invalid query string: %s"}`, err),
		}, nil
	}

	// "url" パラメータを取得
	urlParam := queryParams.Get("url")

	// "url" パラメータが無い場合のエラーハンドリング
	if urlParam == "" {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       `{"error": "Missing 'url' parameter"}`,
		}, nil
	}

	summary, err := summergo.Summarize(urlParam)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       "Failed to summarize",
		}, nil
	}

	// JSONに変換
	responseJSON, err := json.Marshal(summary)
	if err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	// 成功レスポンスを返却
	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(responseJSON),
	}, nil
}

func main() {
	lambda.Start(handler)
}
