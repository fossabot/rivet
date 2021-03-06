/*
 * Copyright (c) 2019. ENNOO - All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package file

import (
	"fmt"
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	path := "/etc/profile"
	exist, err := PathExists(path)
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println(path, "exist =", exist)
	}

	path = "/etc/hello"
	exist, err = PathExists(path)
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println(path, "exist =", exist)
	}

	path = "/ha/oo"
	err = os.MkdirAll(path, os.ModePerm)
	if nil == err {
		exist, err = PathExists(path)
		if nil != err {
			fmt.Println(err.Error())
		} else {
			fmt.Println(path, "exist =", exist)
		}
		err = os.Remove(path)
		if nil != err {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func TestReadFileFirstLine(t *testing.T) {
	txt, err := ReadFileFirstLine("../../../a.txt")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("txt =", txt)
	}

	profile, err := ReadFileFirstLine("/etc/profile")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("profile =", profile)
	}

	hello, err := ReadFileFirstLine("/etc/hello")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hello =", hello)
	}
}

func TestReadFileByLine(t *testing.T) {
	hosts, err := ReadFileByLine("/etc/hostname")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hosts =", hosts)
	}

	profile, err := ReadFileByLine("/etc/profile")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("profile =", profile)
	}

	hello, err := ReadFileByLine("/etc/hello")
	if nil != err {
		fmt.Println(err.Error())
	} else {
		fmt.Println("hello =", hello)
	}
}
