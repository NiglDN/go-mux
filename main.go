// main.go

package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"postgres",
		"postgres")

	a.Run(":8010")
}
