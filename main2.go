package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
)


//Stucts usadas para código
type Response struct {
  Height int `json:"height"`
  Hash string `json:"hash"`
  Timestamp string `json:"timestamp"`
}

type Price struct {
  Time int `json:"time"`
  USD int `json:USD`
  EUR int `json:EUR`
}

type Adjustment struct {
  ProgressPercent float32 `json:"progressPercent"`
  DifficultyChange float32 `json:"difficultyChange"`
  EstimatedRetargetDate float32 `json:"estimatedRetargetDate"`
  RemainingBlocks float32 `json:"remainingBlocks"`
  RemainingTime float32 `json:"remainingTime "`
  PreviousRetarget float32 `json:"previousRetarget"`
  NextRetargetHeight float32 `json:"nextRetargetHeight"`
  TimeAvg float32 `json:"timeAvg"`
  AdjustedTimeAvg float32 `json:"adjustedTimeAvg"`
  TimeOffset float32 `json:"timeOffset"`
}

/*type GeneralAddress struct {

}

type TransactionsAddress struct {

}

type TransactionsChainAddress struct {

}

type TransactionsMempoolAddress struct {

}

type UtxoAddress struct {

}

type ValidationAddress struct {

}*/



func main() {

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  }) 

  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Render("index", fiber.Map{})
  })
 
  app.Post("/block", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)
    block := ctx.FormValue("block")

    fmt.Println(block)
    
     resp, err := http.Get("https://mempool.space/api/v1/mining/blocks/timestamp/"+block)

    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result Response

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Can not unmarshal JSON.")
    }

    fmt.Println(result) 
    
    return ctx.Render("comps/block", fiber.Map{
      "Height": result.Height,
      "Hash": result.Hash,
      "Timestamp": result.Timestamp,
    })
  })

  app.Post("/price", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)
    
    resp, err := http.Get("https://mempool.space/api/v1/prices")

    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result Price 

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Cannot unmarshal JSON.") 
    }

    fmt.Println(result)

    return ctx.Render("comps/price", fiber.Map{
      "Time": result.Time,
      "USD": result.USD,
      "EUR": result.EUR,
    })
  })

  app.Post("/adjustment", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)

    resp, err := http.Get("https://mempool.space/api/v1/difficulty-adjustment")

    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result Adjustment 

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Cannot unmarshal JSON.")
      fmt.Println(err)
    }

    fmt.Println(result)

    return ctx.Render("comps/adjustment", fiber.Map{
      "Diff": result.DifficultyChange,
    })
  })

  app.Post("/height", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)

    resp, err := http.Get("https://mempool.space/api/blocks/tip/height")

    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result int

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Cannot unmarshal JSON.")
      fmt.Println(err)
    }

    fmt.Println(result)

    return ctx.Render("comps/height", fiber.Map{
      "Height": result,
    })
  })

  app.Post("/tip-hash", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)

    resp, err := http.Get("https://mempool.space/api/blocks/tip/hash")
    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    stringfiedJson := string(body)
    fmt.Println(stringfiedJson)

    var result *string 
    result = &stringfiedJson

    /*if err := json.Unmarshal([]byte(stringfiedJson), &result); err != nil {
      fmt.Println("Cannot unmarshal JSON.")
      fmt.Println(err)
    }*/     

    fmt.Println(result)

    return ctx.Render("comps/tip-hash", fiber.Map{
      "Hash": result,
    })
  })

  /*app.Post("/get-address", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)

    address := ctx.FormValue("address")

    addressGeneral, err := http.Get("https://mempool.space/api/address/"+address)
    if err != nil {
      log.Fatalln(err)
    }

    agbody, err := ioutil.ReadAll(addressGeneral.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resAG GeneralAddress

    addressTransactions, err := http.Get("https://mempool.space/api/address/"+address+"/txs")
    if err != nil {
      log.Fatalln(err)
    }

    atbody, err := ioutil.ReadAll(addressTransactions.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resAT TransactionsAddress

    addressTransaChain, err := http.Get("https://mempool.space/api/address/"+address+"/txs/chain")
    if err != nil {
      log.Fatalln(err)
    }

    atcbody, err := ioutil.ReadAll(addressTransaChain.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resATC TransactionsChainAddress

    addressTransaMempool, err := http.Get("https://mempool.space/api/address/"+address+"/txs/mempool")
    if err != nil {
      log.Fatalln(err)
    }

    atnbody, err := ioutil.ReadAll(addressTransaMempool.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resATM TransactionsMempoolAddress

    addressUtxo, err := http.Get("https://mempool.space/api/address/"+address+"/utxo")
    if err != nil {
      log.Fatalln(err)
    }

    aubody, err := ioutil.ReadAll(addressUtxo.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resAutxo UtxoAddress

    addressValidation, err := http.Get("https://mempool.space/api/v1/validate-address/"+address)
    if err != nil {
      log.Fatalln(err)
    }

    avbody, err := ioutil.ReadAll(addressValidation.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var resAV ValidationAddress

    return ctx.Render("comps/adressInfo", fiber.Map{

    })
  })*/

  log.Fatal(app.Listen(":9000"))
}
