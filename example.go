package main

import (
	_ "expvar" // Imported for side-effect of handling /debug/vars.
	"flag"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-tigertonic"
	"log"
	"net/http"
	_ "net/http/pprof" // Imported for side-effect of handling /debug/pprof.
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

var (
	cert   = flag.String("cert", "", "certificate pathname")
	key    = flag.String("key", "", "private key pathname")
	config = flag.String("config", "", "pathname of JSON configuration file")
	listen = flag.String("listen", ":8000", "listen address")

	hMux       tigertonic.HostServeMux
	mux, nsMux *tigertonic.TrieServeMux
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: example [-cert=<cert>] [-key=<key>] [-config=<config>] [-listen=<listen>]")
		flag.PrintDefaults()
	}
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// Register endpoints defined in top-level functions below with example
	// uses of Timed go-metrics wrapper.
	mux = tigertonic.NewTrieServeMux()
	mux.Handle(
		"GET",
		"/incidents",
		tigertonic.Timed(tigertonic.Marshaled(get), "incidents", nil),
	)

	// Example use of namespaces.
	nsMux = tigertonic.NewTrieServeMux()
	nsMux.HandleNamespace("", mux)
	nsMux.HandleNamespace("/1.0", mux)

}

func main() {
	flag.Parse()

	// Example use of go-metrics.
	go metrics.Log(
		metrics.DefaultRegistry,
		60e9,
		log.New(os.Stderr, "metrics ", log.Lmicroseconds),
	)

	// Example of parsing a configuration file.
	c := &Config{}
	if err := tigertonic.Configure(*config, c); nil != err {
		log.Fatalln(err)
	}

	server := tigertonic.NewServer(
		*listen,

		// Example use of go-metrics to track HTTP status codes.
		tigertonic.CountedByStatus(
			// Example use of request logging
			tigertonic.Logged(nsMux, nil),
			"http",
			nil,
		),
	)

	// Example use of server.Close to stop gracefully.
	go func() {
		var err error
		if "" != *cert && "" != *key {
			err = server.ListenAndServeTLS(*cert, *key)
		} else {
			err = server.ListenAndServe()
		}
		if nil != err {
			log.Println(err)
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	log.Println(<-ch)
	server.Close()

}

// GET /incidents
func get(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *MyResponse, error) {
	return http.StatusOK, nil, &MyResponse{"FIRE", "STUFF"}, nil
}
