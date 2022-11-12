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

// Prompt user to input text data
name := s.Input("Enter your name: ")

// Parse splitted text by specified string
// Ignores empty elements
// Trim each string in a result
arr := s.SplitStringBy("a, b, c", ",") // ["a","b","c"]
```

# Argument Parsing

Check https://github.com/AldieNightStar/argulo for more help

```go
a := s.ArgParser("name").
	Param("param", "This is my param").
	Sample("-param a").
	Sample("-param b").
	RequiredParam("param2", "This is required param").
	Sample("-param2 data").
	Build().
	ParseOs() // or Parse(args)

if a.ValidateOk() {
	// Do something here
}
```