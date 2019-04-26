package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Raw struct {
	id   string
	stmt string
}

func (r Raw) String() string {
	return r.stmt
}

func (r Raw) Subbed() string {
	return r.stmt
}

func (r Raw) AV() map[string]*dynamodb.AttributeValue {
	return nil
}

func (cf ConditionFunc) Raw(stmt string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Raw{clauseID(c), stmt})
		return c
	}
}
