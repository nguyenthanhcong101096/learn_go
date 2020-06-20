package main

import (
  "unicode"
  "strings"
)

func isPalindrome(str string) bool {
  start := 0
  end := len(str) - 1

  str = strings.ToLower(str)

  for start < end {
  
    for !unicode.IsLetter([]rune(str)[start]) {
      start++
    }
    for !unicode.IsLetter([]rune(str)[end]) {
      end--
    }
    
    if []rune(str)[start] != []rune(str)[end] {
      return false
    }

    start++
    end--
  }

  return true
}

func main() {
  str := "A man, a plan, a canal: Panama"
  isPalindrome(str)
}