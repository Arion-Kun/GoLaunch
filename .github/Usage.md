## 📝 Usage

`Check if a launch arg is present, and print its values if present`

Example of printing a single argument: `./example.exe --test a b c -d -e`

```go
import (
	"fmt"
	"github.com/Arion-Kun/GoLaunch"
)

func main() {
	contains, val := GoLaunch.TryGetValue("--test")
	if !contains {
		println("Test is not found")
		return
	}

	fmt.Println(val)

}
```
Output:
```
["a", "b", "c"]
```
---
Example of printing all launch arguments: `./example.exe --test a b c -d something -e`
```go
import (
	"fmt"
	"github.com/Arion-Kun/GoLaunch"
)

func main() {
  for m := range GoLaunch.GetSanitizedArgs() {
    for s, strings := range m {
      fmt.Println(s, strings)
    }
  }
}
```
Output:
```
--test ["a", "b", "c"]
-d ["something"]
-e []
```
