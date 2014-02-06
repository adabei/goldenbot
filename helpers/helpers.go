package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func Tail(file string, ch chan string) {
	cmd := exec.Command("tail", "-f", "-n", "0", file)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)

		if err != nil {
			break
		}

		if n != 0 {
			ch <- string(buf[0:n])
		}

	}
	close(ch)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("tail: finished")
}

// TailFallback monitors and returns new lines added to a file by other processes.
// It sends every new line on the passed channel.
// Use this only if tail is not available on your OS.
func TailFallback(file string, ch chan string) {
	currentSize := getFileSize(file)

	for {
		// Only read if the file size has changed
		if getFileSize(file) == currentSize {

			time.Sleep(5000 * time.Millisecond)
			continue
		} else if getFileSize(file) < currentSize {
			currentSize = 0
		}

		fi, _ := os.Open(file)
		defer fi.Close()

		// Information
		//pos, _ := fi.Seek(0, os.SEEK_END)
		//pos, _ := fi.Seek(currentSize, 0)

		fi.Seek(currentSize, 0)
		scanner := bufio.NewScanner(fi)

		for scanner.Scan() {
			ch <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error scanning input: ", err)
		}

		currentSize = getFileSize(file)
	}
}

// getFileSize calculates the file size of a given file.
func getFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		log.Fatal("Could not get size of log file.")
	}

	return info.Size()
}
