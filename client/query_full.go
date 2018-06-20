package client

import (
	"fmt"
	"github.com/dyninc/qstring"
)


type QueryType struct {
	ID   []string
	View string
}

func MakeQuery(newIDs []string) string {
	query := &QueryType{
		ID: newIDs,
		View: "full",
	}
	q, errQ := qstring.MarshalString(query)
	if errQ != nil {
		fmt.Println(errQ)
	}

	return q
}
