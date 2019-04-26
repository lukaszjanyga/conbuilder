package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

type Condition struct {
	id      string
	Clauses []Clause
}

func (c Condition) String() string {
	return c.build(false)
}

func (c Condition) Subbed() string {
	return c.build(true)
}

func (c Condition) build(sub bool) string {
	clauseStrings := []string{}
	for _, clause := range c.Clauses {
		var clauseStr string
		if sub {
			clauseStr = clause.Subbed()
		} else {
			clauseStr = clause.String()
		}
		clauseStrings = append(clauseStrings, clauseStr)
	}
	return "(" + strings.Join(clauseStrings, " ") + ")"
}

func (c Condition) AV() map[string]*dynamodb.AttributeValue {
	avs := map[string]*dynamodb.AttributeValue{}
	for _, clause := range c.Clauses {
		clauseAVs := clause.AV()
		for key, val := range clauseAVs {
			avs[key] = val
		}
	}
	return avs
}

type ConditionFunc func(id ...string) Condition

func New() ConditionFunc {
	return func(id ...string) Condition {
		idVal := "0"
		if len(id) > 0 {
			idVal = id[0]
		}
		return Condition{
			id: idVal,
		}
	}
}

func (cf ConditionFunc) String() string {
	return cf().String()
}

func (cf ConditionFunc) Subbed() string {
	return cf().Subbed()
}

func (cf ConditionFunc) AV() map[string]*dynamodb.AttributeValue {
	return cf().AV()
}

func (cf ConditionFunc) Inner(condFunc ConditionFunc) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		condition := condFunc(clauseID(c))
		c.Clauses = append(c.Clauses, condition)
		return c
	}
}
