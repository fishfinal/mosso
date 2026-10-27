<div align="center">
  <img src="assets/banner/mosso.jpg" alt="mosso banner" />
</div>

<div align="center">

<img src="https://img.shields.io/badge/Go-%3E%3D1.18-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
<img src="https://img.shields.io/badge/license-Apache--2.0-blue?style=for-the-badge" alt="License" />
<img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white" alt="Build Status" />
<img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge" alt="PRs Welcome" />

</div>

**MOSSO** is a lightweight, out-of-the-box development kit designed for rapid debugging and logging in Go. It helps you visualize complex data structures as formatted JSON directly in your console or save them to a file.

## :rocket: Installation

```shell
go get -u github.com/channellabs/mosso
```

## :rainbow: Features

- [x] Convert and print any `interface{}` to formatted JSON.
- [x] Automatically trace and display the caller's file name and line number.
- [x] Add custom informational tags to distinguish different debug outputs.
- [x] Toggle visual borders (Start/End markers) for better readability.
- [x] Write the debug output directly to a specific file.
- [x] Customize the number of blank lines appended at the bottom of the output.

## :ship: Quick Start

```go
package main

import (
  "github.com/channellabs/mosso"
)

type Sex uint8

const (
  SexOfUnknown Sex = iota
  SexOfMale
  SexOfFemale
)

type User struct {
  Name string `json:"name"`
  Age  uint8  `json:"age"`
  Sex  Sex    `json:"sex"`
}

func main() {
  var users = []User{
    {
      Name: "tom",
      Age:  18,
      Sex:  SexOfMale,
    },
    {
      Name: "tom",
      Age:  18,
      Sex:  SexOfFemale,
    },
  }

  // Console Output Example:
  // main.go:33
  // --- CurrentType([]main.User) | Start ------------------
  // [
  //   {
  //     "name": "tom",
  //     "age": 18,
  //     "sex": 1
  //   },
  //   {
  //     "name": "tom",
  //     "age": 18,
  //     "sex": 2
  //   }
  // ]
  // --- CurrentType([]main.User) | Finish ------------------
  //
  mosso.DebugShowContentWithJSON(users)
}
```

## :wrench: Advanced Usage & Options

`mosso` uses the functional options pattern, allowing you to highly customize the output to fit your debugging needs.

### Customizing Output

```go
data := map[string]string{"username": "fishfinal", "email": "contact.fishfinal@gmail.com"}

// Example: Add an info tag, hide type borders, and write to a file
mosso.DebugShowContentWithJSON(data,
  mosso.WithInfo("UserInfo"),                   // Add a custom tag/description
  mosso.WithSpecificOfVisual(true),             // Show the "Start" and "Finish" borders
  mosso.WithSpecificOfWriteFile(true),          // Enable writing to file
  mosso.WithSpecificOfFilename("hello.json"),   // Specify the output filename (default is mosso.txt)
)
```

### Getting Output as a String

If you don't want to print directly to the console but need the string representation:

```go
jsonStr := mosso.DebugContentWithJSON(data, 
  mosso.WithSpecificOfVisual(false), // Disable borders for clean JSON string
)
```

## :gear: API Overview

| Function | Description |
| --- | --- |
| `DebugShowContentWithJSON(data interface{}, ops ...Option)` | Prints the formatted JSON directly to the standard output. |
| `DebugContentWithJSON(data interface{}, ops ...Option) string` | Returns the formatted debug string instead of printing it. |

### Available Options

| Option | Description |
| --- | --- |
| `WithSpecificOfFormat(bool)` | Enable/disable JSON indentation (default: `true`). |
| `WithInfo(string)` | Add a custom description/info tag next to the type. |
| `WithSpecificOfVisual(bool)` | Enable/disable "Start/Finish" visual borders (default: `true`). |
| `WithSpecificOfShowType(bool)` | Show/hide the data type in the output (default: `true`). |
| `WithSpecificOfLineNumber(bool)` | Show/hide the caller's file name and line number (default: `true`). |
| `WithSpecificOfBottomSpaceLine(int)` | Set the number of blank lines at the bottom (default: `1`). |
| `WithSpecificOfWriteFile(bool)` | Enable/disable writing the output to a file (default: `false`). |
| `WithSpecificOfFilename(string)` | Set the target filename when writing to a file (default: `"mosso.txt"`). |

## :memo: License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
