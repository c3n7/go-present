package main

import (
  "context"

  "github.com/mum4k/termdash"
  "github.com/mum4k/termdash/container"
  "github.com/mum4k/termdash/linestyle"
  "github.com/mum4k/termdash/terminal/termbox"
  "github.com/mum4k/termdash/terminal/terminalapi"
  "github.com/mum4k/termdash/widgets/text"
)


func main() {
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
  if err := title.Write("The Horseman."); err != nil {
    panic(err)
  }

  content, err := text.New(text.WrapAtRunes())
  if err != nil {
    panic(err)
  }
	if err := content.Write("Mark me, the end is nigh, the rider on a pale white horse is coming, and his judgement is death. God save the republic, God help the helpless, God damn the lawless."); err != nil {
		panic(err)
	}

  slide, err := container.New(
    t,
    container.Border(linestyle.Light),
    container.BorderTitle("PRESS Q TO QUIT"),
    container.SplitHorizontal(
      container.Top(
        container.Border(linestyle.Light),
        container.BorderTitle("Title"),
        container.PlaceWidget(title),
      ),
      container.Bottom(
        container.Border(linestyle.Light),
        container.BorderTitle("Content"),
        container.PlaceWidget(content),
      ),
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
