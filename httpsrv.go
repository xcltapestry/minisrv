package minisrv

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
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

type HTTPServer struct {
	srv        *http.Server
	router     *mux.Router
	middleware *negroni.Negroni

	writeTimeout time.Duration
	readTimeout  time.Duration
}

type RouteFunc func(m *mux.Router)

func NewHTTPServer() *HTTPServer {
	srv := &HTTPServer{}
	srv.readTimeout = time.Duration(10 * time.Second)
	srv.writeTimeout = time.Duration(10 * time.Second)
	return srv
}

func (s *HTTPServer) AddRoute(f RouteFunc) *HTTPServer {
	f(s.mux())
	return s
}

func (s *HTTPServer) mux() *mux.Router {
	if s.router != nil {
		return s.router
	}
	s.router = mux.NewRouter().StrictSlash(false)
	return s.router
}

type MiddlewareFunc func(m *negroni.Negroni)

func (s *HTTPServer) AddMiddleware(f MiddlewareFunc) *HTTPServer {
	f(s.negroini())
	return s
}

func (s *HTTPServer) negroini() *negroni.Negroni {
	if s.middleware != nil {
		return s.middleware
	}
	s.middleware = negroni.Classic()
	return s.middleware
}

const _defaultAddr = ":8082"

func (s *HTTPServer) ListenAndServe(addrs ...string) error {

	s.middleware.UseHandler(s.router)

	addr := _defaultAddr
	if len(addrs) == 0 {
		for _, v := range addrs {
			addr = v
		}
	}
	svc := &http.Server{
		Addr:           addr,
		Handler:        s.middleware,
		ReadTimeout:    s.readTimeout,
		WriteTimeout:   s.writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return svc.ListenAndServe()
}

func (s *HTTPServer) WithReadTimeout(dur time.Duration) {
	s.readTimeout = dur
}

func (s *HTTPServer) WithWriteTimeout(dur time.Duration) {
	s.writeTimeout = dur
}
