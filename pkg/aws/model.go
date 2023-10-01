package aws

type DataPoint struct {
	Date        string  `dynamodbav:"date" json:"date"`
	Time        string  `dynamodbav:"time" json:"time"`
	Temperature float64 `dynamodbav:"temperature" json:"temperature"`
	Humidity    float64 `dynamodbav:"humidity" json:"humidity"`
	Ttl         int64   `dynamodbav:"ttl" json:"ttl"`
}
