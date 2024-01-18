// pubsubplus-go-client
//
// Copyright 2021-2024 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"solace.dev/go/messaging/internal/generator"
)

// This file is not part of the exported library, and is used to generate subcode_generated.go.
// It can be invoked by running `go generate` in the the errors directory.
// The environment variable SOLCLIENT_H must be set to the absolute path to include/solClient.h.

const outputFile string = "subcode_generated.go"
const header string = "package subcode\n\n// Code generated by subcode_generator.go via go generate. DO NOT EDIT.\nconst (\n"
const footer string = ")\n"
const namePrefix string = ""
const subcodeType string = "Code"
const enumName string = "solClient_subCode"
const enumPrefix string = "SOLCLIENT_SUBCODE"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must specify SOLCLIENT_H as first argument")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	generator.FromEnum(inputFile, outputFile, enumName, enumPrefix, namePrefix, subcodeType, header, footer, false)
}
