# DynamoDB condition builder

A condition building API for AWS DynamoDB.

# Usage

## Import

```go
import (
  cb "github.com/lukaszjanyga/conbuilder"
)
```

## Build conditions

```go
var condition = cb.New().
	Not(cb.New().
		In("state", "S", "created", "active").
		Or().
		Between("count", "1", "10")).
	And().
	Inner(cb.New().
		AttrExists("id").
		Or().
		LT("price", "10", "N").
		And().
		GT("amount", "2000").
		Or().
		Eq("special_offer", "true", "b")).
	Or().
	NEq("flag", "xyz1", "s").
	Or().
	BeginsWith("name", "Jo")

var condition2 = cb.New().
	Not(cb.New().
		In("#other_field", "S", ":my_value1", ":my_value2").
		Or().
		Between("#between_field", ":low", ":high"))

fmt.Println(condition.And().Raw(condition2.String()).String())
fmt.Print("\n")
fmt.Println(condition.Subbed())
fmt.Print("\n")
fmt.Println(condition.AV())
```

## Output

```
(NOT (state IN (created, active) OR count BETWEEN 1 AND 10) AND (attribute_exists(id) OR price < 10 AND amount > 2000 OR special_offer = true) OR flag <> xyz1 OR begins_with(name,Jo) AND (NOT (#other_field IN (:my_value1, :my_value2) OR #between_field BETWEEN :low AND :high)))

(NOT (state IN (:0_0_0in0, :0_0_0in1) OR count BETWEEN :0_0_2_l AND :0_0_2_r) AND (attribute_exists(id) OR price < :0_2_2 AND amount > :0_2_4 OR special_offer = :0_2_6) OR flag <> :0_4 OR begins_with(name,:0_6))

map[:0_0_2_r:{
  S: "10"
} :0_4:{
  S: "xyz1"
} :0_6:{
  S: "Jo"
} :0_2_2:{
  N: "10"
} :0_2_4:{
  S: "2000"
} :0_2_6:{
  BOOL: true
} :0_0_0in0:{
  S: "created"
} :0_0_0in1:{
  S: "active"
} :0_0_2_l:{
  S: "1"
}]
```