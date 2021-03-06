package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AttrNotExists struct {
	id   string
	Path string
}

func (ane AttrNotExists) String() string {
	return "attribute_not_exists(" + ane.Path + ")"
}

func (ane AttrNotExists) Subbed() string {
	return ane.String()
}

func (ane AttrNotExists) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) AttrNotExists(path string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrNotExists{clauseID(c), path})
		return c
	}
}
