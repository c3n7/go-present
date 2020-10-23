package main

import (
  "context"
  "fmt"
  "os"
  "bufio"

  "github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
  "github.com/mum4k/termdash/container"
  "github.com/mum4k/termdash/linestyle"
  "github.com/mum4k/termdash/terminal/termbox"
  "github.com/mum4k/termdash/terminal/terminalapi"
  "github.com/mum4k/termdash/widgets/text"
)


func main() {
  // Read from file
  file, err := os.Open("test1.md")
  if err != nil {
      fmt.Println(err)
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      lines = append(lines, scanner.Text())
  }


  if err := scanner.Err(); err != nil {
      fmt.Println(err)
  }

  t, err := termbox.New()
  if err != nil {
    panic(err)
  }
  defer t.Close()

  ctx, cancel := context.WithCancel(context.Background())

    title, err := text.New()
  if err != nil {
    panic(err)
  }
  if err := title.Write("The Horseman.", text.WriteCellOpts(cell.FgColor(cell.ColorYellow))); err != nil {
    panic(err)
  }

  content, err := text.New(text.WrapAtRunes())
  if err != nil {
    panic(err)
  }
  // Write the read content
  for _, l := range lines {
    if err := content.Write(fmt.Sprintln(l)); err != nil {
      panic(err)
    }
  }

  slide, err := container.New(
    t,
    container.Border(linestyle.Light),
    container.BorderTitle("PRESS Q TO QUIT"),
    container.SplitHorizontal(
      container.Top(
        container.Border(linestyle.Light),
        // container.BorderTitle("Title"),
        container.PlaceWidget(title),
        container.BorderTitleAlignCenter(),
      ),
      container.Bottom(
        container.Border(linestyle.Light),
        // container.BorderTitle("Content"),
        container.PlaceWidget(content),
      ),
			container.SplitPercent(20),
    ),
  )
	if err != nil {
		panic(err)
	}

    quitter := func(k *terminalapi.Keyboard) {
    if k.Key == 'q' || k.Key == 'Q' {
      cancel()
    } else if k.Key == 'g' || k.Key == 'G' {
      if err := title.Write("Judgement.", text.WriteReplace()); err != nil {
        panic(err)
      }
    }
  }

  if err := termdash.Run(ctx, t, slide, termdash.KeyboardSubscriber(quitter)); err != nil {
    panic(err)
  }
}
