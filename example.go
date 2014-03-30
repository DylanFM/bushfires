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
	listen = flag.String("listen", "127.0.0.1:8000", "listen address")

	hMux       tigertonic.HostServeMux
	mux, nsMux *tigertonic.TrieServeMux
)

// A version string that can be set with
//
//     -ldflags "-X main.Version VERSION"
//
// at compile-time.
var Version string

type context struct {
	Username string
}

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: example [-cert=<cert>] [-key=<key>] [-config=<config>] [-listen=<listen>]")
		flag.PrintDefaults()
	}
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	// We'll use this CORSBuilder to set Access-Control-Allow-Origin headers
	// on certain endpoints.
	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins("*")

	// Register endpoints defined in top-level functions below with example
	// uses of Timed go-metrics wrapper.
	mux = tigertonic.NewTrieServeMux()
	mux.Handle(
		"GET",
		"/stuff/{id}",
		cors.Build(tigertonic.Timed(
			tigertonic.Marshaled(get),
			"GET-stuff-id",
			nil,
		)),
	)

	// Example use of the version endpoint.
	mux.Handle("GET", "/version", tigertonic.Version(Version))

	// Example use of namespaces.
	nsMux = tigertonic.NewTrieServeMux()
	nsMux.HandleNamespace("", mux)
	nsMux.HandleNamespace("/1.0", mux)

	// Example use of virtual hosts.
	hMux = tigertonic.NewHostServeMux()
	hMux.Handle("example.com", nsMux)

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

			// Example use of request logging, redacting the word SECRET
			// wherever it appears.
			tigertonic.Logged(

				// Example use of WithContext, which is required in order to
				// use Context within any handlers.  The second argument is a
				// zero value of the type to be used for all actual request
				// contexts.
				tigertonic.WithContext(hMux, context{}),

				func(s string) string {
					return s
				},
			),
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

// GET /stuff/{id}
func get(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *MyResponse, error) {
	return http.StatusOK, nil, &MyResponse{u.Query().Get("id"), "STUFF"}, nil
}
