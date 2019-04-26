package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LT struct {
	id    string
	Field string
	Value string
	Type string
}

func (lt LT) String() string {
	return lt.Field + " < " + lt.Value
}

func (lt LT) Subbed() string {
	return lt.Field + " < " + subKey(lt.id)
}

func (lt LT) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(lt.id)
	value := valueOfType(lt.Value, lt.Type)
	return map[string]*dynamodb.AttributeValue{
		key: &value,
	}
}

func (cf ConditionFunc) LT(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, LT{clauseID(c), field, value, firstSafe(typ)})
		return c
	}
}
