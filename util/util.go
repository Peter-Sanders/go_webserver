package util 

import "os"

func Get_sql(file string) (string){
  b, err := os.ReadFile("sql/" + file + ".sql")
  if err != nil {
    return " "
  }
  str := string(b)
  return str
}
