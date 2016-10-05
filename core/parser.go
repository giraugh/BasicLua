package basic

import (
  //"fmt"
)

/* Parse */
func Parse(toks []token) string {
  s := ""
  ind := 0
  for {
    //break when done
    if ind >= len(toks) { break }

    //set local tok
    tok := toks[ind]

    //If we arent the last
    if ind < len(toks)-1 {
      ntok := toks[ind+1]

      //PRINT + STRING
      if tok.name == "PRINT" && ntok.name == "STRING" {
        s += "print( \""+ntok.value+"\" )\n"
      }

      //PRINT + EXPR
      if tok.name == "PRINT" && ntok.name == "EXPRESSION" {
        s += "print( "+ntok.value+" )\n"
      }
    }

    //Increase Index Naturally
    ind++
  }


  return s
}
