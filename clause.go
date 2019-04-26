package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Clause interface {
	String() string
	Subbed() string
	AV() map[string]*dynamodb.AttributeValue
}
