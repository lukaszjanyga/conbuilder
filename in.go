package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

type In struct {
	id       string
	Field    string
	Operands []string
}

func (i In) String() string {
	return i.Field + " IN (" + strings.Join(i.Operands, ", ") + ")"
}

func (i In) Subbed() string {
	return i.Field + " IN (" + strings.Join(i.Operands, ", ") + ")"
}

func (i In) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) In(field string, operands ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, In{id: clauseID(c), Field: field, Operands: operands})
		return c
	}
}
