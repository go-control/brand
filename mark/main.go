// Command mark writes the go-control mark, svg/color/go-control.svg:
//
//	cd mark && go run . > ../svg/color/go-control.svg
//
// go-control is control systems, and the mark is the closed loop every control
// textbook draws: a summing junction, a block, the output, and the feedback
// path that carries the output back to the junction.
//
// Nothing is placed by hand. The constants below give the sizes of the parts;
// the program derives every coordinate from them, measures the extent of the
// drawing and centres it on the square. Change a size and the mark redraws
// itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#33BF66" // the gradient's light stop
	dark  = "#0D7131" // its dark stop

	stroke = 13.0 // the family's line width

	inLen  = 12.0 // the reference line entering the junction
	radius = 20.0 // the summing junction
	link   = 24.0 // clear space from the junction to the block, between stroke centres
	blockW = 60.0 // the plant block
	blockH = 56.0
	corner = 10.0 // the block's corner radius
	outLen = 56.0 // the output line leaving the block
	tap    = 28.0 // where the feedback leaves the output line, from the block's right edge
	drop   = 56.0 // how far below the forward line the feedback path runs
)

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	// Forward path along y = 0, from x = 0 to the end of the output.
	cx := inLen + radius           // the junction's centre
	blockL := cx + radius + link   // the block's left edge
	blockR := blockL + blockW      // its right edge
	end := blockR + outLen         // the end of the output line
	fb := blockR + tap             // the feedback's x
	width := end                   // the drawing spans x in [0, end]
	top, bottom := -blockH/2, drop // and y in [top, drop]

	// Centre the drawing on the square.
	dx := 128 - width/2
	dy := 128 - (top+bottom)/2
	x := func(v float64) string { return num(v + dx) }
	y := func(v float64) string { return num(v + dy) }

	paths := []string{
		// the reference entering the junction, and the junction
		fmt.Sprintf(`<path d="M%s %sH%s"/>`, x(0), y(0), x(cx-radius)),
		fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, x(cx), y(0), num(radius)),
		// the junction to the block
		fmt.Sprintf(`<path d="M%s %sH%s"/>`, x(cx+radius), y(0), x(blockL)),
		// the block
		fmt.Sprintf(`<rect x="%s" y="%s" width="%s" height="%s" rx="%s"/>`,
			x(blockL), y(top), num(blockW), num(blockH), num(corner)),
		// the output
		fmt.Sprintf(`<path d="M%s %sH%s"/>`, x(blockR), y(0), x(end)),
		// the feedback: down from the output, back under the block, up into the junction
		fmt.Sprintf(`<path d="M%s %sV%sH%sV%s"/>`, x(fb), y(0), y(drop), x(cx), y(radius)),
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-control">
  <defs><linearGradient id="ctgrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#ctgrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(paths, "\n    "))
}
