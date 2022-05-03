# Functions to use Golang as Script
```go
import s "github.com/AldieNightStar/goscriptable"
```
* Scripts
```go
if s.IsFileExist("test.txt") { }

// Will delete folder Recursively as well
s.Delete("folder2")

// Copy file
s.CopyFile(file1, file2)

// Copy folder
s.CopyDirectory(dir1, dir2)

// HexBased time string
s.TimeString()

// Unpack ZIP
s.UnpackZip("files.zip", "files")

// Завантажити з інтернету
data := s.HttpGet(url)

// Generate code block with {} body
s.Block("func (a, b)"
	"line1",
	"line2",
	"etc..."
)

// Each line in text is tabulated by n symbols
s.Tabulated(4, text)

s.MultiplyString(4, "FLIP") // FLIPFLIPFLIPFLIP

// Run Command with workdir specified
s.Run(workDir, "rm", "-rf", "./files")

// Run Command with workdir specified
// Will also stdout return as text
text := s.RunAndReturn(workDir, "rm", "-rf", "./files")

// True if element is in a list
s.IsIn("dir", []string{"a", "b", "c", "dir"})

// Create folder
s.CreateFolder("folder")

// Get program args
args := s.GetOsArgs()

// Create file
file := s.CreateFile(fileName)

// Write to a file
s.WriteToFile("file.txt", data)

// Open file or return nil
file := s.Open(fileName)

// Parse args into map[string]string
a := s.ParseArgs(s.GetOsArgs())

// Prompt user to input text data
name := s.Input("Enter your name: ")

// Parse splitted text by specified string
// Ignores empty elements
// Trim each string in a result
arr := s.SplitStringBy("a, b, c", ",") // ["a","b","c"]
```
* Data / Config
```go
// Load config
d := s.LoadData("x.dat")

// Write config
d["a"] = "Alex"
d["i"] = "Ihor"
d["l"] = "Liza"

// Save config
s.SaveData("x.dat", d)
```

* String Divider
```go
d := s.NewDivString(", ")
d.Append("Hello")
d.Append("Hi")
d.Append("123")

d.String() // Hello, Hi, 123

d.Divide("+") // Hello+Hi+123
```

* Find inside function
	* Find's tags `[[name: ...]]` in a lot of texts or sites or some other sources
```go
// Will return data inside tag. If tag is "abc" then it will return data inside "[[abc: ...]]"
s.FindInside(data, tagname)
```
