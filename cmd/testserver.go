// Copyright © 2017 Christian Kniep <christian@qnib.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "log"
    "fmt"
    "net/http"
    "strings"
    "strconv"

    "github.com/urfave/cli"
)

// HTTPServer holds the codes and the channel
type HTTPServer struct {
    StatusSequence []string
    StatusChan chan string
    Quiet bool
}

// NewHTTPServer returns HTTPServer
func NewHTTPServer(quiet bool, sc []string) HTTPServer {
    return HTTPServer{
        StatusSequence: sc,
        StatusChan: make(chan string),
        Quiet: quiet,
    }
}

func (hs *HTTPServer) fillStatusCodeChan() string {
    for {
        for _, c := range hs.StatusSequence {
            hs.StatusChan <- c
        }
    }
}

func (hs *HTTPServer) httpHandler(w http.ResponseWriter, r *http.Request) {
    sc := <-hs.StatusChan
    if ! hs.Quiet {
        log.Println("Statuscode: ", sc)
    }
    ci, _ := strconv.Atoi(sc)
    if strings.HasPrefix(sc, "2") {
        w.WriteHeader(ci)
    } else if strings.HasPrefix(sc, "4") || strings.HasPrefix(sc, "5") {
        http.Error(w, http.StatusText(ci), ci)
    }
    fmt.Fprintf(w, fmt.Sprintf("%s, %s\n", sc, http.StatusText(ci)))
}

// Run starts the webserver
func (hs *HTTPServer) Run() {
    go hs.fillStatusCodeChan()
    log.Printf("Start Testserver. Quiet:%v, Seq:%v\n", hs.Quiet, hs.StatusSequence)
    http.HandleFunc("/", hs.httpHandler)
    http.ListenAndServe(":8080", nil)

}


// TestServer serves the status-sequence of HTTP codes
func TestServer(c *cli.Context) error {
    hserver := NewHTTPServer(c.Bool("quiet"), strings.Split(c.String("status-sequence"), ","))
    hserver.Run()
    return nil
}
