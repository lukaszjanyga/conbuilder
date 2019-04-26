package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AttrExists struct {
	id   string
	Path string
}

func (ae AttrExists) String() string {
	return "attribute_exists(" + ae.Path + ")"
}

func (ae AttrExists) Subbed() string {
	return "attribute_exists(" + ae.Path + ")"
}

func (ae AttrExists) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
