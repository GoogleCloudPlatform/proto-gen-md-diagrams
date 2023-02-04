/*
 * Copyright 2023 Google LLC
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
 */

package proto

import (
	"bufio"
	"strings"
)

type TestScanner struct {
	internalScanner *bufio.Scanner
}

func (ts *TestScanner) Scan() bool {
	return ts.internalScanner.Scan()
}

func (ts *TestScanner) Text() string {
	return ts.internalScanner.Text()
}

func (ts *TestScanner) Split(splitFunction bufio.SplitFunc) {
	ts.internalScanner.Split(splitFunction)
}

func (ts *TestScanner) Buffer(buf []byte, max int) {
	ts.internalScanner.Buffer(buf, max)
}

func (ts *TestScanner) Err() error {
	return ts.Err()
}

func (ts *TestScanner) Bytes() []byte {
	return ts.Bytes()
}

func (ts *TestScanner) ReadLine() *Line {
	return NewLine(ts.internalScanner.Text())
}

func NewTestScanner(in string) *TestScanner {
	return &TestScanner{
		internalScanner: bufio.NewScanner(strings.NewReader(in)),
	}
}
