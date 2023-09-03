package aws

type DataPoint struct {
	Date        string `dynamodbav:"date" json:"date"`
	Time        string `dynamodbav:"time" json:"time"`
	Temperature int64  `dynamodbav:"temperature" json:"temperature"`
	Humidity    int64  `dynamodbav:"humidity" json:"humidity"`
	Ttl         int64  `dynamodbav:"ttl" json:"ttl"`
}
