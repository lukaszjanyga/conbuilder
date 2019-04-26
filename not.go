package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Not struct {
	id        string
	Condition Condition
}

func (n Not) String() string {
	return "NOT " + n.Condition.String()
}

func (n Not) Subbed() string {
	return "NOT " + n.Condition.Subbed()
}

func (n Not) AV() map[string]*dynamodb.AttributeValue {
	return n.Condition.AV()
}

func (cf ConditionFunc) Not(condFunc ConditionFunc) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		innerID := clauseID(c)
		condition := condFunc(innerID)
		c.Clauses = append(c.Clauses, Not{innerID, condition})
		return c
	}
}
