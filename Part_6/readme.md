## Building Microservices With Golang

#### Part Six
In this part we are going to look at the Go Validator package and how it can be used to validate <b>JSON.</b>br>

## Package Validator

<img align="right" src="https://raw.githubusercontent.com/go-playground/validator/v9/logo.png">[![Join the chat at https://gitter.im/go-playground/validator](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-playground/validator?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
![Project status](https://img.shields.io/badge/version-10.7.0-green.svg)
[![Build Status](https://travis-ci.org/go-playground/validator.svg?branch=master)](https://travis-ci.org/go-playground/validator)
[![Coverage Status](https://coveralls.io/repos/go-playground/validator/badge.svg?branch=master&service=github)](https://coveralls.io/github/go-playground/validator?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/validator)](https://goreportcard.com/report/github.com/go-playground/validator)
[![GoDoc](https://godoc.org/github.com/go-playground/validator?status.svg)](https://pkg.go.dev/github.com/go-playground/validator/v10)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package validator implements value validations for structs and individual fields based on tags.

It has the following **unique** features:

-   Cross Field and Cross Struct validations by using validation tags or custom validators.
-   Slice, Array and Map diving, which allows any or all levels of a multidimensional field to be validated.
-   Ability to dive into both map keys and values for validation
-   Handles type interface by determining it's underlying type prior to validation.
-   Handles custom field types such as sql driver Valuer see [Valuer](https://golang.org/src/database/sql/driver/types.go?s=1210:1293#L29)
-   Alias validation tags, which allows for mapping of several validations to a single tag for easier defining of validations on structs
-   Extraction of custom defined Field Name e.g. can specify to extract the JSON name while validating and have it available in the resulting FieldError
-   Customizable i18n aware error messages.
-   Default validator for the [gin](https://github.com/gin-gonic/gin) web framework; upgrading from v8 to v9 in gin see [here](https://github.com/go-playground/validator/tree/master/_examples/gin-upgrading-overriding)

Installation
------------

Use go get.

	go get github.com/go-playground/validator

Then import the validator package into your own code.

	import "github.com/go-playground/validator"

-------