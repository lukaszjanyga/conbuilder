package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LTE struct {
	id    string
	Field string
	Value string
}

func (lte LTE) String() string {
	return lte.Field + " <= " + lte.Value
}

func (lte LTE) Subbed() string {
	return lte.Field + " <= " + lte.Value
}

func (lte LTE) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) LTE(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, LTE{clauseID(c), field, value})
		return c
	}
}
