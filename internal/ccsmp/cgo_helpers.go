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

package ccsmp

/*
#include <stdio.h>
#include <stdlib.h>

#include "./ccsmp_helper.h"
*/
import "C"
import "unsafe"

// This file contains a few helper functions that let us manipulate C data without
// importing C directly. These are used in testing as well as in demultiplexing user
// pointers.

// ToGoBytes convert a block of C data into a byte array using the given length
func ToGoBytes(pointer unsafe.Pointer, length int) []byte {
	return C.GoBytes(pointer, C.int(length))
}

// ToCArray convert the given go string array into a c string array and returns a function
// that can be deferred to free the array
func ToCArray(arr []string, nullTerminated bool) (cArray **C.char, freeArray func()) {
	length := len(arr)
	if nullTerminated {
		length++
	}
	cArr := make([]*C.char, length)
	for i, val := range arr {
		cArr[i] = C.CString(val)
	}
	if nullTerminated {
		cArr = append(cArr, nil)
	}
	freeFunction := func() {
		for i := 0; i < len(cArr); i++ {
			C.free(unsafe.Pointer(cArr[i]))
		}
	}
	return (**C.char)(unsafe.Pointer(&cArr[0])), freeFunction
}

func GenerateTestSolClientErrorInfoWrapper(returnCode SolClientReturnCode, subCode SolClientSubCode, responseCode SolClientResponseCode, errorInfo string) *SolClientErrorInfoWrapper {
	errorInfoWrapper := SolClientErrorInfoWrapper{}
	errorInfoWrapper.ReturnCode = returnCode
	detailedErrorInfo := C.solClient_errorInfo_detailed_t{}
	detailedErrorInfo.SubCode = subCode
	detailedErrorInfo.ResponseCode = responseCode
	for i := 0; i < len(errorInfo) && i < len(detailedErrorInfo.ErrorStr)-1; i++ {
		detailedErrorInfo.ErrorStr[i] = (C.char)(errorInfo[i])
	}
	detailedErrorInfo.ErrorStr[len(detailedErrorInfo.ErrorStr)-1] = '\x00'
	errorInfoWrapper.DetailedErrorInfo = &detailedErrorInfo
	return &errorInfoWrapper
}
