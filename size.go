package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Size struct {
	id   string
	Path string
}

func (s Size) String() string {
	return "size(" + s.Path + ")"
}

func (s Size) Subbed() string {
	return "size(" + s.Path + ")"
}

func (s Size) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
