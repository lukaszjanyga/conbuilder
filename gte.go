package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GTE struct {
	id    string
	Field string
	Value string
}

func (gte GTE) String() string {
	return gte.Field + " >= " + gte.Value
}

func (gte GTE) Subbed() string {
	return gte.Field + " >= " + gte.Value
}

func (gte GTE) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) GTE(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, GTE{clauseID(c), field, value})
		return c
	}
}
