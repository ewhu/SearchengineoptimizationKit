// cmd/searchengineoptimizationkit/main.go
package main

import (
"flag"
"log"
"os"

"searchengineoptimizationkit/internal/searchengineoptimizationkit"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := searchengineoptimizationkit.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
