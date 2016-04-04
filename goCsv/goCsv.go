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
  val1 := fmt.Sprintf("%csvDataContainer",status)
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
  csvDataContainer := &VariableInit{}
  file, err := os.Open("../samples/myfile_sample.csv")
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
  csvDataContainer.remaining = records[1:]
  c.assignValue(header)
  csvDataContainer.Lenght = 200
  psqlConn.CreateTable(c.returnValue())
  csvDataContainer.GenertaeString()
  jsonMsg,err := getResponse("200", "Success!!")
  return jsonMsg,err
}

func (csvDataContainer *VariableInit) GenertaeString() {
  // for example csvDataContainer.Lenght = 100
  csvDataContainer.stringVal = ""// store the rows of data (as a string) to be inserted in table
  csvDataContainer.str = ""// used for comma and semicolon separator
  if len(csvDataContainer.remaining) > csvDataContainer.Lenght{
    csvDataContainer.first_100 =  csvDataContainer.remaining[:csvDataContainer.Lenght]
    csvDataContainer.remaining =  csvDataContainer.remaining[csvDataContainer.Lenght:]
  }else{
    csvDataContainer.first_100 = csvDataContainer.remaining
    csvDataContainer.remaining = nil
  }
  for j,rec_values := range csvDataContainer.first_100 {
    csvDataContainer.stringVal += "("
    for i,rec := range rec_values {
      if rec = rec; rec == "" { rec = "NULL"}
      checkAlphanumeric, _ := regexp.Compile("([a-zA-Z]+)")
      if (checkAlphanumeric.MatchString(rec)) {rec = "'"+rec+"'"}
      if csvDataContainer.str  = ","; i == 0 { csvDataContainer.str = "" }
      csvDataContainer.stringVal += (csvDataContainer.str+rec)
    }
    if (len(csvDataContainer.first_100) == j+1) {  csvDataContainer.stringVal += ");"
    } else {csvDataContainer.stringVal += "),"}
  }
  psqlConn.InsertRec(csvDataContainer.stringVal)
  csvDataContainer.first_100 = nil
  fmt.Printf("--------------> %d\n", len(csvDataContainer.remaining))
  if (len(csvDataContainer.remaining) !=0) { csvDataContainer.GenertaeString() }
}
