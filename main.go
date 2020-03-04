package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"howett.net/plist"
)

type color struct {
	Alpha float64 `plist:"Alpha Component"`
	Space string  `plist:"Color Space"`
	R     float64 `plist:"Red Component"`
	G     float64 `plist:"Green Component"`
	B     float64 `plist:"Blue Component"`
}

func (c *color) toHex() string {
	return fmt.Sprintf("#%.2x%.2x%.2x", int(c.R*255), int(c.G*255), int(c.B*255))
}

type iterm struct {
	Color0      color `plist:"Ansi 0 Color"`
	Color1      color `plist:"Ansi 1 Color"`
	Color2      color `plist:"Ansi 2 Color"`
	Color3      color `plist:"Ansi 3 Color"`
	Color4      color `plist:"Ansi 4 Color"`
	Color5      color `plist:"Ansi 5 Color"`
	Color6      color `plist:"Ansi 6 Color"`
	Color7      color `plist:"Ansi 7 Color"`
	Color8      color `plist:"Ansi 8 Color"`
	Color9      color `plist:"Ansi 9 Color"`
	Color10     color `plist:"Ansi 10 Color"`
	Color11     color `plist:"Ansi 11 Color"`
	Color12     color `plist:"Ansi 12 Color"`
	Color13     color `plist:"Ansi 13 Color"`
	Color14     color `plist:"Ansi 14 Color"`
	Color15     color `plist:"Ansi 15 Color"`
	Background  color `plist:"Background Color"`
	Badge       color `plist:"Badge Color"`
	Bold        color `plist:"Bold Color"`
	Cursor      color `plist:"Cursor Color"`
	CursorGuide color `plist:"Cursor Guide Color"`
	CursorText  color `plist:"Cursor Text Color"`
	Forgeground color `plist:"Foreground Color"`
	Link        color `plist:"Link Color"`
	Selected    color `plist:"Selected Text Color"`
	Selection   color `plist:"Selection Color"`
}

func (i *iterm) toKittyConfig() string {
	return fmt.Sprintf(`
foreground %s
background %s
selection_foreground %s
selection_background %s
color0 %s
color1 %s
color2 %s
color3 %s
color4 %s
color5 %s
color6 %s
color7 %s
color8 %s
color9 %s
color10 %s
color11 %s
color12 %s
color13 %s
color14 %s
color15 %s

# URL styles
url_color %s
url_style single

# Cursor styles
cursor %s
  `, i.Forgeground.toHex(),
		i.Background.toHex(),
		i.Selected.toHex(),
		i.Selection.toHex(),
		i.Color0.toHex(),
		i.Color1.toHex(),
		i.Color2.toHex(),
		i.Color3.toHex(),
		i.Color4.toHex(),
		i.Color5.toHex(),
		i.Color6.toHex(),
		i.Color7.toHex(),
		i.Color8.toHex(),
		i.Color9.toHex(),
		i.Color10.toHex(),
		i.Color11.toHex(),
		i.Color12.toHex(),
		i.Color13.toHex(),
		i.Color14.toHex(),
		i.Color15.toHex(),
		i.Link.toHex(),
		i.Cursor.toHex(),
	)
}

func main() {
	flag.Parse()
	var in io.ReadSeeker
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	var data iterm
	decoder := plist.NewDecoder(in)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.toKittyConfig())
}
