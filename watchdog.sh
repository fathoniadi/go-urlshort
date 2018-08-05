#!/bin/bash

watchmedo auto-restart --patterns="*.html;*.go" --ignore-directories --recursive -- go run main.go