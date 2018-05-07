package main

import "fmt"

func main() {
	name := "Pranay Wankhede"

	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset= "UTF-8">
<title>Hello World!!!</title>
<body>
<h1>` + name + ` says "Hello"</h1>
</body>
</html>
`
	fmt.Println(tpl)

}
