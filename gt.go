package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GT struct {
	id    string
	Field string
	Value string
	Type string
}

func (gt GT) String() string {
	return gt.Field + " > " + gt.Value
}

func (gt GT) Subbed() string {
	return gt.Field + " > " + subKey(gt.id)
}

func (gt GT) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(gt.id)
	value := valueOfType(gt.Value, gt.Type)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) GT(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, GT{clauseID(c), field, value, firstSafe(typ)})
		return c
	}
}
