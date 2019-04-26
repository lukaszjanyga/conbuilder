package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Not struct {
	id        string
	Condition Condition
}

func (n Not) String() string {
	return "NOT " + n.Condition.String()
}

func (n Not) Subbed() string {
	return "NOT " + n.Condition.Subbed()
}

func (n Not) AV() map[string]*dynamodb.AttributeValue {
	return n.Condition.AV()
}
