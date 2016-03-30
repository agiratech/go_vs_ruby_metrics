package psqlConn

import (
  "database/sql"
   _ "github.com/lib/pq"
  "fmt"
  "encoding/json"
)

var db *sql.DB
var dbErr error

const timeLayout = "Jan 2, 2006"

const tableName = "testcsv"

func InitDB(){
  db, dbErr = sql.Open("postgres", "user=postgres dbname=godb password=root")

  if dbErr != nil {
    fmt.Println(dbErr)
  }
  fmt.Println("DB Connected")
}

func checkLast(index,len int) string{
  var operator string
  if operator = ")"; index < len { operator = ","}
  return operator
}

func CreateTable(headers []string){
  check_table := fmt.Sprintf("create table IF NOT EXISTS %v",tableName)
  set_primary_key := "("
  set_variables := ""
  length  := len(headers)
  for i,header := range headers {
    value := fmt.Sprintf("%v varchar(255)%v",header,checkLast(i+1,length))
    set_variables += value
  }
  finalStr := check_table+set_primary_key+set_variables
  fmt.Println(finalStr)
  _, err := db.Exec(finalStr)
  if err != nil {
    fmt.Println(err)
  }
}

func InsertRec(rec string) {
  insertQuery := fmt.Sprintf("INSERT INTO %v (policyid,statecode,county,eq_site_limit,hu_site_limit,fl_site_limit,fr_site_limit,tiv_2011,tiv_2012,eq_site_deductible,hu_site_deductible,fl_site_deductible,fr_site_deductible,point_latitude,point_longitude,line,construction,point_granularity) values %v",tableName,rec)
  // fmt.Println(insertQuery)
  _, err := db.Exec(insertQuery)
  if err != nil {
    fmt.Println(err)
  }
}


type info struct {
  Policyid string  `json:"id"`
  Statecode string `json:"statecode"`
  County string `json:"county"`
  // eq_site_limit string
  // hu_site_limit string
  // fl_site_limit string
  // fr_site_limit string
  // tiv_2011 string
  // tiv_2012 string
  // eq_site_deductible string
  // hu_site_deductible string
  // fl_site_deductible string
  // fr_site_deductible string
  // point_latitude string
  // point_longitude string
  // tiv_2012eq_site_deductible string
}


func SearchByNmae(name string)(error,string){
  data_param := []info{}
  like_name := "%"+name+"%"
  query := fmt.Sprintf("SELECT policyid,statecode,county FROM %v WHERE county iLIKE $1",tableName)
  // fmt.Println(query)
  rows, err  := db.Query(query,like_name)
  if err != nil {
    fmt.Println(err)
  }
  for rows.Next() {
    var r info
    err = rows.Scan(&r.Policyid,&r.Statecode,&r.County)
    if err != nil {
      fmt.Println("Scan: %v", err)
    }
    data_param = append(data_param,r)
  }
  res, _ := json.Marshal(data_param)
  return err,string(res)
}