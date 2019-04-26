package conbuilder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

type In struct {
	id       string
	Field    string
	Operands []string
	Type     string
}

func (i In) String() string {
	return i.Field + " IN (" + strings.Join(i.Operands, ", ") + ")"
}

func (i In) Subbed() string {
	subs := []string{}
	for cnt := range i.Operands {
		subs = append(subs, subKey(i.id)+fmt.Sprintf("in%d", cnt))
	}
	return i.Field + " IN (" + strings.Join(subs, ", ") + ")"
}

func (i In) AV() map[string]*dynamodb.AttributeValue {
	avs := map[string]*dynamodb.AttributeValue{}
	for cnt, op := range i.Operands {
		val := valueOfType(op, i.Type)
		avs[subKey(i.id)+fmt.Sprintf("in%d", cnt)] = &val
	}
	return avs
}

func (cf ConditionFunc) In(field string, typ string, operands ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, In{id: clauseID(c), Field: field, Operands: operands, Type: typ})
		return c
	}
}
