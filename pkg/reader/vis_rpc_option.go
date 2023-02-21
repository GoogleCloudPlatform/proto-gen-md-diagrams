package reader

import (
  "strings"
)

func NewRpcOptionVisitor(debug bool) *RpcOptionVisitor {
  return &RpcOptionVisitor{}
}

type RpcOptionVisitor struct {
}

func (o *RpcOptionVisitor) CanVisit(line *Line) bool {
  return strings.HasPrefix(line.Syntax, "option")
}

func (o *RpcOptionVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
  optionName := in.Syntax[strings.Index(in.Syntax, "(")+1 : strings.Index(in.Syntax, ")")]
  var optionBody = ""
  for scanner.Scan() {
    in := scanner.ReadLine()
    if in.Token == ClosedBrace {
      break
    } else {
      optionBody += in.Syntax
    }
  }
  return ProtobufFactory.NewOption(optionName, optionBody, "")
}
