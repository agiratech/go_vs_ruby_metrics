package goCsv

import (
  "encoding/csv"
  "fmt"
  "os"
  "encoding/json"
  "time"
  "regexp"
  "github.com/agiratech/go_vs_ruby_metrics/psqlConn"
  // "unsafe"
  // "log"
)

var start_time time.Time

func init() {
  psqlConn.InitDB()
}

type Message struct {
  Status string
  Body string
  ProcessTime string
}

type VariableInit struct {

  first_100 [][]string
  remaining [][]string
  stringVal string
  str string
  Lenght int
}

type CsvHeader struct {
  header []string
}

func executionTime() string{
  end_time := time.Now()
  delta := end_time.Sub(start_time)
  return delta.String()
}

func SearchByName(name string) (error,string) {
  start_time = time.Now()
  err,rec := psqlConn.SearchByNmae(name)
  fmt.Println(executionTime())
  return err,rec
}

func (c *CsvHeader) assignValue(a []string)  {
  c.header = a
}

func (c *CsvHeader) returnValue() []string {
  return c.header
}

func getResponse(status string, msg string) (string, error){
  val1 := fmt.Sprintf("%v",status)
  val2 := fmt.Sprintf("%v",msg)
  time := (executionTime())
  fmt.Println(time)
  message := Message{val1, val2,time}
  jbMsg, err := json.Marshal(message)

  if err != nil {
    return "", err
  }

  jsonMsg := string(jbMsg[:]) // converting byte array to string
  return jsonMsg, nil
}

func Import() (string, error) {
  start_time = time.Now()
  v := &VariableInit{}
  file, err := os.Open("samples/myfile_sample.csv")
  c := &CsvHeader{}
  if err != nil {
    pwd, _ := os.Getwd()
    fmt.Println(err, pwd)
    jsonMsg,err := getResponse("400", "File Parse Error!! 1")
    return jsonMsg,err
  }
  // automatically call Close() at the end of current method
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = ','
  records, err := reader.ReadAll()
  if err != nil {
    fmt.Println("CSV Parse Error:", err)
    jsonMsg,err := getResponse("400", "File Parse Error!! 2")
    return jsonMsg,err
  }
  header := records[0]
  v.remaining = records[1:]
  c.assignValue(header)
  v.Lenght = 200
  psqlConn.CreateTable(c.returnValue())
  v.GenertaeString()
  jsonMsg,err := getResponse("200", "Success!!")
  return jsonMsg,err
}

// func (v *VariableInit) GenertaeString(record [][]string) {
//   v.stringVal = ""
//   v.str = ""
//   var lenght int = 200
//   if len(record) > lenght{
//     temp_1 := [][]string{}
//     temp_2 := [][]string{}
//     v.first_100 = append(temp_1,record[1:lenght]...)
//     v.remaining = append(temp_2,record[lenght:]...)
//   }else{
//     v.first_100 = record
//   }
//   for j,rec_values := range v.first_100 {
//     v.stringVal += "("
//     for i,rec := range rec_values {
//       if rec = rec; rec == "" { rec = "NULL"}
//       r, _ := regexp.Compile("([a-zA-Z]+)")
//       if (r.MatchString(rec)) {rec = "'"+rec+"'"}
//       if v.str  = ","; i == 0 { v.str = "" }
//       v.stringVal += (v.str+rec)
//     }
//     if (len(v.first_100) == j+1) {  v.stringVal += ");"
//     } else {v.stringVal += "),"}
//   }
//   psqlConn.InsertRec(v.stringVal)
//   v.first_100 = nil
//   fmt.Printf("record: %T, %d\n", record, unsafe.Sizeof(record))
//   fmt.Printf("Remaining: %T, %d\n", v.remaining, unsafe.Sizeof(v.remaining))

//   if (len(v.remaining) !=0) { v.GenertaeString(v.remaining) }
// }

func (v *VariableInit) GenertaeString() {
  v.stringVal = ""
  v.str = ""
  // var lenght int = 200
  if len(v.remaining) > v.Lenght{
    v.first_100 =  v.remaining[:v.Lenght]
    v.remaining =  v.remaining[v.Lenght:]

  }else{
    v.first_100 = v.remaining
    v.remaining = nil
  }
  for j,rec_values := range v.first_100 {
    v.stringVal += "("
    for i,rec := range rec_values {
      if rec = rec; rec == "" { rec = "NULL"}
      r, _ := regexp.Compile("([a-zA-Z]+)")
      if (r.MatchString(rec)) {rec = "'"+rec+"'"}
      if v.str  = ","; i == 0 { v.str = "" }
      v.stringVal += (v.str+rec)
      // fmt.Printf("V ---> %v\n",unsafe.Sizeof(v))

    }
    if (len(v.first_100) == j+1) {  v.stringVal += ");"
    } else {v.stringVal += "),"}
  }
  psqlConn.InsertRec(v.stringVal)
  v.first_100 = nil
  fmt.Printf("--------------> %d\n", len(v.remaining))
  if (len(v.remaining) !=0) { v.GenertaeString() }
}
