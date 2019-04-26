package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type BeginsWith struct {
	id     string
	Path   string
	Substr string
}

func (bw BeginsWith) String() string {
	return "begins_with(" + bw.Path + "," + bw.Substr + ")"
}

func (bw BeginsWith) Subbed() string {
	return "begins_with(" + bw.Path + "," + bw.Substr + ")"
}

func (bw BeginsWith) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
