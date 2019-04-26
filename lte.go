package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LTE struct {
	id    string
	Field string
	Value string
	Type string
}

func (lte LTE) String() string {
	return lte.Field + " <= " + lte.Value
}

func (lte LTE) Subbed() string {
	return lte.Field + " <= " + subKey(lte.id)
}

func (lte LTE) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(lte.id)
	value := valueOfType(lte.Value, lte.Type)
	return map[string]*dynamodb.AttributeValue{
		key: &value,
	}
}

func (cf ConditionFunc) LTE(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, LTE{clauseID(c), field, value, firstSafe(typ)})
		return c
	}
}
