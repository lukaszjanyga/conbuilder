package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Between struct {
	id         string
	Field      string
	BoundLeft  string
	BoundRight string
	Typ        []string
}

func (b Between) String() string {
	return b.Field + " BETWEEN " + b.BoundLeft + " AND " + b.BoundRight
}

func (b Between) Subbed() string {
	return b.Field + " BETWEEN " + subKey(b.id+"_l") + " AND " + subKey(b.id+"_r")
}

func (b Between) AV() map[string]*dynamodb.AttributeValue {
	keyLeft := subKey(b.id + "_l")
	keyRight := subKey(b.id + "_r")
	valueLeft := valueOfType(b.BoundLeft, b.Typ...)
	valueRight := valueOfType(b.BoundRight, b.Typ...)
	return map[string]*dynamodb.AttributeValue{
		keyLeft:  &valueLeft,
		keyRight: &valueRight,
	}
}

func (cf ConditionFunc) Between(field, boundLeft, boundRight string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Between{clauseID(c), field, boundLeft, boundRight, typ})
		return c
	}
}
