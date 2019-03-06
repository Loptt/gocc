// Code generated by gocc; DO NOT EDIT.

package errors

import (
	"fmt"
	"strings"

	"github.com/goccmack/gocc/example/sr/token"
)

type ErrorSymbol interface {
}

type Error struct {
	Err            error
	ErrorToken     *token.Token
	ErrorSymbols   []ErrorSymbol
	ExpectedTokens []string
	StackTop       int
}

func (e *Error) String() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "Error")
	if e.Err != nil {
		fmt.Fprintf(w, " %s\n", e.Err)
	} else {
		fmt.Fprintf(w, "\n")
	}
	fmt.Fprintf(w, "Token: type=%d, lit=%s\n", e.ErrorToken.Type, e.ErrorToken.Lit)
	fmt.Fprintf(w, "Pos: offset=%d, line=%d, column=%d\n", e.ErrorToken.Pos.Offset, e.ErrorToken.Pos.Line, e.ErrorToken.Pos.Column)
	fmt.Fprintf(w, "Expected one of: ")
	for _, sym := range e.ExpectedTokens {
		fmt.Fprintf(w, "%s ", sym)
	}
	fmt.Fprintf(w, "ErrorSymbol:\n")
	for _, sym := range e.ErrorSymbols {
		fmt.Fprintf(w, "%v\n", sym)
	}
	return w.String()
}

func (e *Error) Error() string {
	w := new(strings.Builder)
	fmt.Fprintf(w, "Error in S%d: %s, %s", e.StackTop, token.TokMap.TokenString(e.ErrorToken), e.ErrorToken.Pos.String())
	if e.Err != nil {
		fmt.Fprintf(w, ": %+v", e.Err)
	} else {
		fmt.Fprintf(w, ", expected one of: ")
		for _, expected := range e.ExpectedTokens {
			fmt.Fprintf(w, "%s ", expected)
		}
	}
	return w.String()
}
