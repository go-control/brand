# go-control — brand assets

Official logos for the **go-control** organization.

<img src="avatar/go-control.png" width="160" alt="go-control">

The mark is the closed loop every control textbook draws: a summing junction, a
block, the output, and the feedback path that carries the output back to the
junction. It is written by `mark/main.go`, which derives every coordinate from
the sizes of the parts and centres the drawing; nothing is placed by hand. Its
gradient (`#33BF66` to `#0D7131`) is a blue-green (hue 142) that sets go-control apart
from the emerald of go-ndarray, go-fft and go-notebook (163), the yellow-green
of go-fem (115), the blue of go-plot, the violet of go-numerics and go-symbolic,
the magenta of go-bignum, the coral of go-sparse, the orange of go-approx and
the rose of go-dsp; its dark end reads on white at 6.13:1 and its light end on
a dark page at 8.10:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-control.svg` differs.

---
*Generated from `svg/color/go-control.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
