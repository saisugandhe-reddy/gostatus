package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/saisugandhe-reddy/gostatus/internal/checker"
)

func main() {
	timeout := flag.Duration("timeout", 5*time.Second, "request timeout")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "usage: gostatus [--timeout 5s] <url> [url...]")
		os.Exit(1)
	}

	results := checker.CheckURLs(urls, *timeout)

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "URL\tSTATUS\tLATENCY\tERROR")
	for _, result := range results {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\n", result.URL, result.StatusCode, result.Latency.Round(time.Millisecond), result.Error)
	}
	_ = w.Flush()
}
