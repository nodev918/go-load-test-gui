# FYNE

## enviroment
- ubuntu 20.04 (focal)

## getting started
1. install go
2. play tutorial example
```
pckage main

import (
  "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    )

func main() {
    a := app.New()
      w := a.NewWindow("Hello World")

        w.SetContent(widget.NewLabel("Hello World!"))
          w.ShowAndRun()
}ackage
```
$ go run .
$ FYNE_THEME=light go run . //for light theme
3. install fyne command tool and do package
- $ go install fyne.io/fyne/v2/cmd/fyne@latest
- $ fyne package  or  $ fyne package -icon drop_image_path
- parameters [-o windows]
- fyne install 
