package stdreq

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

func GetListStdRequest(requestEvent events.APIGatewayProxyRequest) ListStdRequest {
	request := ListStdRequest{}
	if limit, ok := requestEvent.QueryStringParameters["limit"]; ok {
		limitInt, err := strconv.Atoi(limit)
		if err == nil {
			request.Limit = aws.Int(limitInt)
		}
	}

	if page, ok := requestEvent.QueryStringParameters["page"]; ok {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			request.Page = aws.Int(pageInt)
		}
	}

	tz := ""
	if val, ok := requestEvent.Headers["tz"]; ok {
		tz = val
		request.TZ = tz
	}

	if order, ok := requestEvent.QueryStringParameters["order"]; ok {
		request.Order = aws.String(order)
	}
	return request
}
