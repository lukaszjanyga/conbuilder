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

func (cf ConditionFunc) And() ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, And{id: clauseID(c)})
		return c
	}
}

func (cf ConditionFunc) Or() ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Or{})
		return c
	}
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

func (cf ConditionFunc) In(field string, operands ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, In{id: clauseID(c), Field: field, Operands: operands})
		return c
	}
}

func (cf ConditionFunc) Between(field, boundLeft, boundRight string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Between{clauseID(c), field, boundLeft, boundRight, typ})
		return c
	}
}

func (cf ConditionFunc) Eq(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Eq{clauseID(c), field, value})
		return c
	}
}

func (cf ConditionFunc) NEq(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, NEq{clauseID(c), field, value, typ})
		return c
	}
}

func (cf ConditionFunc) LT(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, LT{clauseID(c), field, value, typ})
		return c
	}
}

func (cf ConditionFunc) GT(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, GT{clauseID(c), field, value})
		return c
	}
}

func (cf ConditionFunc) LTE(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, LTE{clauseID(c), field, value})
		return c
	}
}

func (cf ConditionFunc) GTE(field, value string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, GTE{clauseID(c), field, value})
		return c
	}
}

func (cf ConditionFunc) AttrExists(path string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrExists{clauseID(c), path})
		return c
	}
}

func (cf ConditionFunc) AttrNotExists(path string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrNotExists{clauseID(c), path})
		return c
	}
}

func (cf ConditionFunc) AttrType(path, attrType string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrType{clauseID(c), path, attrType})
		return c
	}
}

func (cf ConditionFunc) BeginsWith(path, substr string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, BeginsWith{clauseID(c), path, substr})
		return c
	}
}

func (cf ConditionFunc) Contains(path, operand string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Contains{clauseID(c), path, operand})
		return c
	}
}

func (cf ConditionFunc) Size(path string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Size{clauseID(c), path})
		return c
	}
}
