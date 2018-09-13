package prototype

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	_ "npd/queue"
	"os"
	"strings"
)

type Npd struct {
	ErrorLog *log.Logger

	HTTPServer  *http.Server
	HTTPSServer *http.Server

	BackendProxy *httputil.ReverseProxy
}

func (w *Npd) Start() error {
	//If no error log is specified use the stderr as log
	if w.ErrorLog == nil {
		w.ErrorLog = log.New(os.Stderr, "Npd error: ", 0)
	}

	//If no backend proxy is set create a default one
	if w.BackendProxy == nil {
		w.BackendProxy = &httputil.ReverseProxy{
			Director:  w.modifyRequestForBalancing,
			Transport: http.DefaultTransport,
			ErrorLog:  w.ErrorLog,
		}
	}

	httpServer := w.HTTPServer
	if httpServer == nil {
		httpServer = &http.Server{
			Addr:     ":8080",
			Handler:  w,
			ErrorLog: w.ErrorLog,
		}
	}

	httpsServer := w.HTTPSServer
	if httpsServer == nil {
		httpsServer = &http.Server{
			Addr:     ":4443",
			Handler:  w,
			ErrorLog: w.ErrorLog,
		}
	}

	return httpServer.ListenAndServe()
}

func (w *Npd) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	w.BackendProxy.ServeHTTP(rw, req)
}

//Modifies and prepares the http request for trannsemission to the backend
// here loadbalancing logic is applied
func (w *Npd) modifyRequestForBalancing(req *http.Request) {
	//TODO make ip dynamic
	target, err := url.Parse("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	targetQuery := target.RawQuery

	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
	if _, ok := req.Header["User-Agent"]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		req.Header.Set("User-Agent", "")
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

type requestPhaseFilterResponse struct {
	//The action to be taken
	//TODO make this a proper struct
	Action func(http.ResponseWriter)

	//Should the request be interupted?
	Interupt bool
}
