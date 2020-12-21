package types

import (
	"errors"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

// _validate for ID support
func _validate(value string) error {
	if len(value) < 3 {
		return errors.New("The minimum length required is 3")
	}
	return nil
}

// ID support
var ID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ID",
	Description: "The `id` scalar type represents a ID Object.",
	Serialize: func(value interface{}) interface{} {
		fmt.Println("SERIALIZE", value)
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		var err error
		switch value.(type) {
		case string:
			err = _validate(value.(string))
		default:
			err = errors.New("Must be of type string.")
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("PARSE VALUE", value)
		return value
	},
	ParseLiteral: func(valueAst ast.Value) interface{} {
		if valueAst.GetKind() == kinds.StringValue {
			err := _validate(valueAst.GetValue().(string))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("PARSE LITERAL", valueAst)
			return valueAst
		} else {
			log.Fatal("Must be of type string.")
			return nil
		}
	},
})
