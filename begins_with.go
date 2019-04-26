package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type BeginsWith struct {
	id     string
	Path   string
	Substr string
	Type   []string
}

func (bw BeginsWith) String() string {
	return "begins_with(" + bw.Path + "," + bw.Substr + ")"
}

func (bw BeginsWith) Subbed() string {
	return "begins_with(" + bw.Path + "," + subKey(bw.id) + ")"
}

func (bw BeginsWith) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(bw.id)
	value := valueOfType(bw.Substr, bw.Type...)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) BeginsWith(path, substr string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, BeginsWith{clauseID(c), path, substr, typ})
		return c
	}
}
