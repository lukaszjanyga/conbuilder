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
			In("state", "created", "active").
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
		NEq("flag", "xyz1", "s")

	var condition2 = cb.New().
		Not(cb.New().
			In("state", "created", "active").
			Or().
			Between("count", "1", "10"))

	fmt.Println(condition.And().Raw(condition2.String()).String())
	fmt.Print("\n")
	fmt.Println(condition.Subbed())
	fmt.Print("\n")
	fmt.Println(condition.AV())
```

## Output

```
(NOT (state IN (created, active) OR count BETWEEN 1 AND 10) AND (attribute_exists(id) OR price < 10 AND amount > 2000 OR special_offer = true) OR flag <> xyz1 AND (NOT (state IN (created, active) OR count BETWEEN 1 AND 10)))

(NOT (state IN (created, active) OR count BETWEEN :0_0_2_l AND :0_0_2_r) AND (attribute_exists(id) OR price < :0_2_2 AND amount > 2000 OR special_offer = :0_2_6) OR flag <> :0_4)

map[:0_0_2_l:{
  S: "1"
} :0_0_2_r:{
  S: "10"
} :0_2_2:{
  N: "10"
} :0_2_6:{
  BOOL: true
} :0_4:{
  S: "xyz1"
}]
```