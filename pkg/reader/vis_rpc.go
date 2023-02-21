package reader

import (
  "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/api"
  "github.com/GoogleCloudPlatform/proto-gen-md-diagrams/pkg/logging"
  "regexp"
  "strings"
)

var RpcLinePattern = `rpc\s+(.*?)\((.*?)\)\s+returns\s+\((.*?)\)(.*)`

func NewRpcVisitor(debug bool) *RpcVisitor {
  return &RpcVisitor{
    RpcLineMatcher: regexp.MustCompile(RpcLinePattern),
    Log:            logging.NewLogger(debug, "rpc_visitor"),
    Visitors: []Visitor{
      &CommentVisitor{},
      NewRpcOptionVisitor(false)},
  }
}

type RpcVisitor struct {
  Log            *logging.Logger
  Visitors       []Visitor
  RpcLineMatcher *regexp.Regexp
}

func (rv *RpcVisitor) CanVisit(line *Line) bool {
  rv.Log.Debugf("Checking: %s - Status : %v", line.Syntax, rv.RpcLineMatcher.MatchString(line.Syntax))
  return rv.RpcLineMatcher.MatchString(line.Syntax)
}

func (rv *RpcVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {
  rv.Log.Debugf("Visiting RPC: %v\n", in.Syntax)

  values := rv.RpcLineMatcher.FindStringSubmatch(in.Syntax)
  out := ProtobufFactory.NewRPC(namespace, values[1], in.Comment)
  ParseInArgs(values, out)
  ParseReturnArgs(values, out)

  var comment = ""
  for scanner.Scan() {
    line := scanner.ReadLine()
    if line.Token == ClosedBrace {
      break
    }
    for _, v := range rv.Visitors {
      if v.CanVisit(line) {
        rt := v.Visit(scanner, in, values[1])
        switch t := rt.(type) {
        case api.RPCOption:
          out.AddRPCOption(t.Name(), Join(Space, comment, t.Comment()), t.Body())
        case string:
          comment += t
        }
      }
    }
  }
  return out
}

// HELPER FUNCTIONS

func ParseInArgs(values []string, rpc api.RPC) {
  inArgs := strings.Split(values[2], Comma)
  for _, i := range inArgs {
    if strings.HasPrefix(i, "stream") {
      rpc.AddInputParameter(true, strings.TrimSpace(i[strings.Index(i, Space):]))
    } else {
      rpc.AddInputParameter(false, strings.TrimSpace(i))
    }
  }
}

func ParseReturnArgs(values []string, rpc api.RPC) {
  returnArgs := strings.Split(values[3], Comma)
  for _, i := range returnArgs {
    if strings.HasPrefix(i, "stream") {
      rpc.AddReturnParameter(true, strings.TrimSpace(i[strings.Index(i, Space):]))
    } else {
      rpc.AddReturnParameter(false, strings.TrimSpace(i))
    }
  }
}
