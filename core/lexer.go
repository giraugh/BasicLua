/*
Lexing: A basic pass sweeps the data and picks up syntax errors while
creating an array of 'tokens' with optional values, i.e a "STRING" token with the value being the contents
of the string.
*/

package basic

import (
  "strconv"
  "strings"
)

type tokenList []token

func (t tokenList) add(input, value string) []token {
  tok := newToken(input, value)
  return append(t, *tok)
}

//Formats error information nicely
func err(input string, curLine, curLineIndex int) serror {
  return serror("Error: @("+strconv.Itoa(curLine)+":"+strconv.Itoa(curLineIndex)+")"+input)
}

/* Lexer Function */
func Lex(input string) (tokenList, serror) {
  //make input uppercase
  input = strings.ToLower(input)

  toks := make(tokenList, 0)
  buff := ""
  state := "default"
  curLine := 1 //keep track of line
  curLineIndex := 0 //keep track of line

  for _, c := range input {
    //For each character char
		char, _ := strconv.Unquote(strconv.QuoteRuneToASCII(c))

    //increase buffer
    buff += char

    switch state {
    case "default":
      if buff == " " {buff = ""}
      if char == "\n" {buff = ""}
      if buff == "end" {buff = ""} //ignored
      if buff == "print" {
        state = "parg"
        toks = toks.add("PRINT","")
        buff = ""
      }
      if buff == "let" {
        state = "larg1"
        toks = toks.add("LET","")
        buff = ""
      }
    case "larg1": //Set argument 1
      if char == "\n" {
        //Error, was expecting a variable name
        return nil, err("Expecting a string, number or expression, got a newline.", curLine, curLineIndex)
      }

      if char == " " {
        if len(buff) > 1 {
          state = "larg2"
          toks = toks.add("VARNAME", buff[:len(buff)-1])
          buff = ""
        } else {
          buff = ""
        }
      }
    case "larg2": //Set argument 1
      if char == " " || char == "\n" {
        state = "default"
        toks = toks.add("VARVAL", buff[:len(buff)-1])
        buff = ""
      }
    case "parg": //Print argument
      if char == "\n" {
        //Error, was expecting a string to start
        return nil, err("Expecting a string or expression, got a newline.", curLine, curLineIndex)
      }

      if buff == " " {
        buff = ""
      }

      //expression start char
      if char == "(" {
        state = "expression"
        buff = ""
      }

      //string start char
      if char == "\"" {
        state = "string"
        buff = ""
      }
    case "string":
      if char == "\"" {
        toks = toks.add("STRING", buff[:len(buff)-1])
        state = "default"
        buff = ""
      }
    case "expression":
      if char == ")" {
        toks = toks.add("EXPRESSION", buff[:len(buff)-1])
        state = "default"
        buff = ""
      }
    }

    //Update line and index
    if char == "\n" {
      curLine++
      curLineIndex = 0
    } else {curLineIndex++}
  }

  return toks, ""
}
