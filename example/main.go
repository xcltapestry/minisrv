package main

/**
 * Copyright 2022 minisrv Author. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @Project golibs
 * @Description
 * @author XiongChuanLiang<br/>(xcl_168@aliyun.com)
 * @license http://www.apache.org/licenses/  Apache v2 License
 * @version 1.0
 */

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/xcltapestry/minisrv"
)

func main() {
	srv := minisrv.NewHTTPServer().
		AddRoute(route).
		AddMiddleware(middleware)
		// ListenAndServe() // or ListenAndServe(":8082")

	go func() {
		fmt.Println("starting the server at port :8082")
		err := srv.ListenAndServe() // or ListenAndServe(":3000")
		if err != nil {
			fmt.Println("[error] could not start the server. ", "error:", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c

	fmt.Println("shutting down the server. ", "received signal", sig)

	//gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)

}

func route(m *mux.Router) {
	m.HandleFunc("/", indexHandler)
	m.HandleFunc("/health", healthHandler)
	m.HandleFunc("/api/v1/actid/%d", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the actid page!")
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index : http://127.0.0.1:8082")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "health : https://github.com/xcltapestry/minisrv")
}

func middleware(n *negroni.Negroni) {
	n.Use(negroni.HandlerFunc(Authorizer))
	n.Use(negroni.HandlerFunc(APIMiddleware))
}

func Authorizer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}

func APIMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(w, r)
}
