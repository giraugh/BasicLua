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
        ind++
      }

      //PRINT + EXPR
      if tok.name == "PRINT" && ntok.name == "EXPRESSION" {
        s += "print( "+ntok.value+" )\n"
        ind++
      }

      //if we arent the second last either
      if ind < len(toks)-2 {
        nntok := toks[ind+2]

        //SET + VARNAME + VARARG
        if tok.name == "LET" && ntok.name == "VARNAME" && nntok.name == "VARVAL" {
          s += ntok.value + " = " + nntok.value + "\n"
          ind++
          ind++
        }
      }
    }

    //Increase Index Naturally
    ind++
  }


  return s
}
