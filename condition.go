package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

// Condition is the root building block of the condition API
type Condition struct {
	id      string
	Clauses []Clause
}

// String returns a string representation without substitutions
func (c Condition) String() string {
	return c.build(false)
}

// Subbed returns a string representation with values replaced by value keys
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

// AV returns a map[string]*dynamodb.AttributeValue containing substituted values
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

// ConditionFunc is a builder function for conditions. It allows many builder methods.
type ConditionFunc func(id ...string) Condition

// New starts a new condition chain
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

// String calls the String() method on the resulting Condition
func (cf ConditionFunc) String() string {
	return cf().String()
}

// Subbed calls the Subbed() method on the resulting Condition
func (cf ConditionFunc) Subbed() string {
	return cf().Subbed()
}

// AV calls the AV() method on the resulting Condition
func (cf ConditionFunc) AV() map[string]*dynamodb.AttributeValue {
	return cf().AV()
}

// Inner nests a condition while maintaining substitution key uniqueness
func (cf ConditionFunc) Inner(condFunc ConditionFunc) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		condition := condFunc(clauseID(c))
		c.Clauses = append(c.Clauses, condition)
		return c
	}
}
