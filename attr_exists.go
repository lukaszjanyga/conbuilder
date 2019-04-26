package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AttrExists struct {
	id   string
	Path string
}

func (ae AttrExists) String() string {
	return "attribute_exists(" + ae.Path + ")"
}

func (ae AttrExists) Subbed() string {
	return ae.String()
}

func (ae AttrExists) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) AttrExists(path string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrExists{clauseID(c), path})
		return c
	}
}
