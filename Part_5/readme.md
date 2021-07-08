## Building Microservices With Golang

#### Part Five
In this part we are going to continue building a simple RESTfulservice using the standard packages in Golang.<br>

<b>Also we are going to learn how to use the Gorilla framework</b>

## Gorilla Framework
[![GoDoc](https://godoc.org/github.com/gorilla/mux?status.svg)](https://godoc.org/github.com/gorilla/mux)
[![CircleCI](https://circleci.com/gh/gorilla/mux.svg?style=svg)](https://circleci.com/gh/gorilla/mux)
[![Sourcegraph](https://sourcegraph.com/github.com/gorilla/mux/-/badge.svg)](https://sourcegraph.com/github.com/gorilla/mux?badge)

![Gorilla Logo](https://cloud-cdn.questionable.services/gorilla-icon-64.png)

https://www.gorillatoolkit.org/pkg/mux

Package `gorilla/mux` implements a request router and dispatcher for matching incoming requests to
their respective handler.

The name mux stands for "HTTP request multiplexer". Like the standard `http.ServeMux`, `mux.Router` matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. The main features are:

* It implements the `http.Handler` interface so it is compatible with the standard `http.ServeMux`.
* Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
* URL hosts, paths and query values can have variables with an optional regular expression.
* Registered URLs can be built, or "reversed", which helps maintaining references to resources.
* Routes can be used as subrouters: nested routes are only tested if the parent route matches. This is useful to define groups of routes that share common conditions like a host, a path prefix or other repeated attributes. As a bonus, this optimizes request matching.

---

* [Install](#install)
* [Examples](#examples)
* [Matching Routes](#matching-routes)
* [Static Files](#static-files)
* [Serving Single Page Applications](#serving-single-page-applications) (e.g. React, Vue, Ember.js, etc.)
* [Registered URLs](#registered-urls)
* [Walking Routes](#walking-routes)
* [Graceful Shutdown](#graceful-shutdown)
* [Middleware](#middleware)
* [Handling CORS Requests](#handling-cors-requests)
* [Testing Handlers](#testing-handlers)
* [Full Example](#full-example)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/gorilla/mux
```