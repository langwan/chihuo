package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/langwan/langgo/helpers/os"
	"math"
	"strings"
)

type Listener struct {
}

func (l Listener) ProgressChanged(event *helper_os.ProgressEvent) {
	if event.EventType == helper_os.TransferCompletedEvent {
		fmt.Println("\nfinish")
	} else {
		progress(float64(event.ConsumedBytes) / float64(event.TotalBytes) * 100.0)
	}

}

func main() {
	listener := Listener{}
	helper_os.CopyFileWatcher("samples/2.mp4", "samples/1.mp4", nil, &listener)
}

func progress(n float64) {
	number := int(math.Round(n))
	bar := strings.Repeat(" ", 100)
	bar = strings.Replace(bar, " ", "=", number)
	tag := fmt.Sprintf("%d%%", number)
	if number >= 100 {
		tag = "ok."
	}
	progress := fmt.Sprintf("%s [ %s%s%s ] %s%s\r", color.Green, color.Yellow, bar, color.Green, color.Purple, tag)
	fmt.Print(progress)
}
