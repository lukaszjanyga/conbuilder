package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AttrNotExists struct {
	id   string
	Path string
}

func (ane AttrNotExists) String() string {
	return "attribute_not_exists(" + ane.Path + ")"
}

func (ane AttrNotExists) Subbed() string {
	return "attribute_not_exists(" + ane.Path + ")"
}

func (ane AttrNotExists) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
