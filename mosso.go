// Copyright 2021 fishfinal <contact.fishfinal@gmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mosso

import (
	"bytes"
	"fmt"
	"os"
	"runtime"

	"github.com/fishfinal/fsx"
)

const (
	DefaultWriteFileName = "mosso.txt"
)

// DebugShowContentWithJSON Output debugging data information,
// which can be customized by option type method.
func DebugShowContentWithJSON(data interface{}, ops ...Option) {
	fmt.Print(DebugContentWithJSON(data, ops...))
}

// DebugContentWithJSON Get debugging data information,
// which can be customized by option type method.
func DebugContentWithJSON(data interface{}, ops ...Option) string {
	options := options{
		format:          true,
		visual:          true,
		typ:             true,
		fileInfo:        true,
		writeFile:       false,
		bottomSpaceLine: 1,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	var show, storage = bytes.Buffer{}, bytes.Buffer{}
	defer show.Reset()
	defer storage.Reset()

	if options.fileInfo {
		_, file, line, _ := runtime.Caller(3)
		show.WriteString(fmt.Sprintf("%s:%d\n", file, line))
	}

	// Main content.
	content := DebugConvertInterfaceToJSONString(data, options.format)

	if options.visual {
		if options.typ {
			start := fmt.Sprintf("--- CurrentType(%T) | Start ------------------\n", data)
			if options.info != nil {
				start = fmt.Sprintf("--- CurrentType(%T) (%s) | Start ------------------\n", data, *options.info)
			}
			show.WriteString(start)

			show.WriteString(fmt.Sprintf("%s\n", content))

			end := fmt.Sprintf("--- CurrentType(%T) | Finish ------------------\n", data)
			if options.info != nil {
				end = fmt.Sprintf("--- CurrentType(%T) (%s) | Finish ------------------\n", data, *options.info)
			}
			show.WriteString(end)
		} else {
			show.WriteString(content)
		}
	} else {
		show.WriteString(content)
	}

	// Space line.
	show.WriteString(fmt.Sprintf("%s", getBottomLinesSpaceStr(options.bottomSpaceLine)))

	// Write content to file.
	writerFileContent(options, storage, content)

	return show.String()
}

func writerFileContent(options options, storage bytes.Buffer, content string) {
	if options.writeFile {
		storage.WriteString(fmt.Sprintf("%s\n", content))

		filename := DefaultWriteFileName
		if options.filename != "" {
			filename = options.filename
		}

		if !fsx.Exists(filename) {
			err := fsx.Touch(filename)
			if err != nil {
				filename = DefaultWriteFileName
			}
		}

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err == nil && file != nil {
			defer file.Close()
			_ = file.Truncate(0)
			_, _ = file.Write(storage.Bytes())
		}
	}
}
