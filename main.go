package sensemicroservice

import (
	"encoding/json"
	"log"

	"github.com/buaazp/fasthttprouter"
	jtltojson "github.com/nicholasvuono/jtl-to-json"
	"github.com/valyala/fasthttp"
)

func getProtocolLevelResultsList(ctx *fasthttp.RequestCtx) {
	data := db.readAll("plu")
	ctx.Write(data)
}

func getBrowserLevelResultsList(ctx *fasthttp.RequestCtx) {
	data := db.readAll("blu")
	ctx.Write(data)
}

func addProtocolLevelResult(ctx *fasthttp.RequestCtx) {
	var r *jtltojson.Result
	b := ctx.PostBody()
	err := json.Unmarshal(b, r)
	checkErr(err)
	db.write("plu", r)
}

func addBrowserLevelResult(ctx *fasthttp.RequestCtx) {
	var r *jtltojson.Result
	b := ctx.PostBody()
	err := json.Unmarshal(b, r)
	checkErr(err)
	db.write("blu", r)
}

func routes() *fasthttprouter.Router {
	r := fasthttprouter.New()
	r.GET("/results/plu/list", getProtocolLevelResultsList)
	r.GET("/results/blu/list", getBrowserLevelResultsList)
	r.POST("/results/plu/add", addProtocolLevelResult)
	r.POST("/results/blu/add", addBrowserLevelResult)
	return r
}

func main() {
	log.Fatal(fasthttp.ListenAndServe(":10000", routes().Handler))
}
