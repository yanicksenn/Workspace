package gtasks

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"testing"

	"google.golang.org/api/tasks/v1"
)

// captureOutput is a helper to capture stdout from a function.
func captureOutput(t *testing.T, f func()) string {
	t.Helper()
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// extractID is a helper to extract an ID from the CLI's output.
func extractID(output string) string {
	re := regexp.MustCompile(`\((.*?)\)`) 
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func printTaskLists(lists *tasks.TaskLists) {
	if len(lists.Items) == 0 {
		fmt.Println("No task lists found.")
		return
	}

	fmt.Println("Task Lists:")
	for _, item := range lists.Items {
		fmt.Printf("- %s (%s)\n", item.Title, item.Id)
	}
}

func printTaskList(list *tasks.TaskList) {
	fmt.Printf("ID:    %s\n", list.Id)
	fmt.Printf("Title: %s\n", list.Title)
	fmt.Printf("Self:  %s\n", list.SelfLink)
}