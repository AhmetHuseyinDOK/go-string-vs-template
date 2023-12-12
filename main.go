package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func main() {
	strs := generateRandomStrings(2000000)
	// HTTP TEST
	templ := template.Must(template.New("test").Parse(`{{range .}}{{.}}{{end}}`))
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		templ.Execute(w, strs)
	})

	http.HandleFunc("/string-builder", func(w http.ResponseWriter, r *http.Request) {
		var builder strings.Builder

		for _, word := range strs {
			builder.WriteString(word)
		}

		result := builder.String()

		w.Write([]byte(result))
	})

	// STDOUT TEST

	// templ := template.Must(template.New("test").Parse(`{{range .}}{{.}}{{end}}`))

	// durtempl, memtempl := measurePerformance(func() {
	// 	templ.Execute(os.Stdout, strs)
	// })

	// durstr, memstr := measurePerformance(func() {
	// 	var builder strings.Builder

	// 	for _, word := range strs {
	// 		builder.WriteString(word)
	// 	}

	// 	result := builder.String()
	// 	fmt.Println(result)
	// })

	// cmd := exec.Command("clear") //Linux example, its tested
	// cmd.Stdout = os.Stdout
	// cmd.Run()
	// fmt.Printf("Template:\n")
	// fmt.Printf("\n\n\n\nExecution Time: %s\n", durtempl)
	// fmt.Printf("Memory Usage: %d bytes\n\n\n\n", memtempl)

	// fmt.Printf("String Birleştirme:\n")
	// fmt.Printf("\n\n\n\nExecution Time: %s\n", durstr)
	// fmt.Printf("Memory Usage: %d bytes\n\n\n\n", memstr)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

func measurePerformance(function func()) (time.Duration, uint64) {
	// Start timer
	startTime := time.Now()

	// Record memory usage before execution
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Execute the function
	function()

	// Record memory usage after execution
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate execution time
	executionTime := time.Since(startTime)

	return executionTime, m2.TotalAlloc - m1.TotalAlloc
	// Print execution time and memory usage

}

// randomString generates a random string of length n using a set of characters.
func randomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// generateRandomStrings generates a slice of random strings.
func generateRandomStrings(numStrings int) []string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	strings := make([]string, numStrings)
	for i := 0; i < numStrings; i++ {
		strings[i] = randomString(seededRand.Intn(40))
	}
	return strings
}
