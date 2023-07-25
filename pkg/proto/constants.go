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

// Constants used for parsing and interpretation
const (
	Protobuf3Types = "double,float,int32,int64,uint32,uint64,sint32,sint64,fixed32,fixed64,sfixed32,sfixed64,bool,string,bytes"

	PrefixRepeated = "repeated"
	PrefixMap      = "map"
	PrefixReserved = "reserved"
	PrefixOptional = "optional"

	SpaceRemovalRegex = `\s+`
	Period            = "."
	Empty             = ""
	Space             = " "
	OpenBrace         = "{"
	CloseBrace        = "}"
	OpenBracket       = "["
	ClosedBracket     = "]"
	Semicolon         = ";"
	Comma             = ","
	Pipe              = "|"
	Hyphen            = "-"

	InlineCommentPrefix        = "//"
	MultiLineCommentInitiator  = "/*"
	MultilineCommentTerminator = "*/"
	OpenMap                    = "map<"
	CloseMap                   = ">"
	DoubleQuote                = `"`
	SingleQuote                = `'`
	EndL                       = "\n"
	CommentNewLine             = `:~:`
)

// From gist: https://gist.github.com/ik5/d8ecde700972d4378d87
const (
	InfoColor  = "\033[1;32mINFO: %s\033[0m"
	ErrorColor = "\033[1;31mERROR: %s\033[0m"
	DebugColor = "\033[1;36mDEBUG: %s\033[0m"
)
