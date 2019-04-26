package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Contains struct {
	id      string
	Path    string
	Operand string
	Type    string
}

func (c Contains) String() string {
	return "contains(" + c.Path + "," + c.Operand + ")"
}

func (c Contains) Subbed() string {
	return "contains(" + c.Path + "," + subKey(c.id) + ")"
}

func (c Contains) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(c.id)
	value := valueOfType(c.Operand, c.Type)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) Contains(path, operand string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Contains{clauseID(c), path, operand, firstSafe(typ)})
		return c
	}
}
