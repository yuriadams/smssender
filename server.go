package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yuriadams/smssender/api/controllers"
)

var (
	logOn   *bool
	port    *int
	urlBase string
)

func init() {
	domain := flag.String("d", "localhost", "domain")
	port = flag.Int("p", 8888, "port")
	logOn = flag.Bool("l", true, "log on/off")

	flag.Parse()
	urlBase = fmt.Sprintf("http://%s:%d", *domain, *port)
}

func main() {
	smsc := controllers.NewSMSController()
	r := httprouter.New()
	r.GET("/", smsc.GetSMSHandler)
	r.POST("/api/sendSMS", smsc.SendSMSHandler)

	logging("Starting server %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), r))
}

func logging(format string, v ...interface{}) {
	if *logOn {
		log.Printf(fmt.Sprintf("%s\n", format), v...)
	}
}
