package goscriptable

import (
	"strings"
	"testing"
)

func TestArgsParser(t *testing.T) {
	data := "-name Ihor -age 18 -lastName Fox -numbers 123 321 444 -git -java -golang"
	params := ParseArgs(strings.Split(data, " "))
	if !params.IsPresent("name") {
		t.Fatal("name is not present")
	}
	for _, name := range []string{"name", "age", "lastName", "numbers", "git", "java", "golang"} {
		if !params.IsPresent(name) {
			t.Fatalf("%s is not present as a param", name)
		}
	}
	for _, kn := range []string{"name=Ihor", "age=18", "lastName=Fox", "git=true", "java=true", "golang=true"} {
		arr := strings.Split(kn, "=")
		key := arr[0]
		val := arr[1]
		if params.Get(key).First() != val {
			t.Fatalf("%s is not %s", key, val)
		}
	}
	numbers := params.Get("numbers")
	for i, n := range []string{"123", "321", "444"} {
		if numbers.All()[i] != n {
			t.Fatal("numbers array is wrong")
		}
	}
}
