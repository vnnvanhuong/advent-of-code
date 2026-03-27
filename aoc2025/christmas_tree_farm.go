package aoc2025

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ChristmasTreeFarmFromLines runs part 1 on raw puzzle lines.
func ChristmasTreeFarmFromLines(lines []string) int {
	return ChristmasTreeFarm(lines)
}

// ChristmasTreeFarm1 counts how many regions can fit all listed presents (polyomino packing).
func ChristmasTreeFarm(lines []string) int {
	shapes, regions := parseChristmasTreeFarm(lines)
	orientations := make([][][]point, len(shapes))
	for i := range shapes {
		orientations[i] = uniqueOrientations(shapes[i])
	}
	count := 0
	for _, reg := range regions {
		if regionFits(reg, orientations) {
			count++
		}
	}
	return count
}

type point struct{ r, c int }

func parseChristmasTreeFarm(lines []string) ([][]point, []regionSpec) {
	var shapes [][]point
	var regions []regionSpec

	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}
		if isRegionLine(line) {
			break
		}
		if idx, ok := parseShapeHeader(line); ok {
			i++
			var rows []string
			for i < len(lines) {
				s := lines[i]
				if strings.TrimSpace(s) == "" {
					i++
					break
				}
				if isRegionLine(strings.TrimSpace(s)) {
					break
				}
				if _, ok := parseShapeHeader(strings.TrimSpace(s)); ok {
					break
				}
				rows = append(rows, s)
				i++
			}
			for len(shapes) <= idx {
				shapes = append(shapes, nil)
			}
			shapes[idx] = gridToShape(rows)
			continue
		}
		i++
	}

	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		i++
		if line == "" {
			continue
		}
		if reg, ok := parseRegionLine(line); ok {
			regions = append(regions, reg)
		}
	}

	return shapes, regions
}

func parseShapeHeader(line string) (int, bool) {
	if !strings.HasSuffix(line, ":") {
		return 0, false
	}
	numStr := strings.TrimSuffix(line, ":")
	numStr = strings.TrimSpace(numStr)
	if numStr == "" {
		return 0, false
	}
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, false
	}
	return n, true
}

func isRegionLine(line string) bool {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return false
	}
	dim := strings.TrimSpace(parts[0])
	xIdx := strings.Index(dim, "x")
	if xIdx <= 0 || xIdx >= len(dim)-1 {
		return false
	}
	_, err1 := strconv.Atoi(strings.TrimSpace(dim[:xIdx]))
	_, err2 := strconv.Atoi(strings.TrimSpace(dim[xIdx+1:]))
	return err1 == nil && err2 == nil
}

func parseRegionLine(line string) (regionSpec, bool) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return regionSpec{}, false
	}
	dim := strings.TrimSpace(parts[0])
	xIdx := strings.Index(dim, "x")
	if xIdx <= 0 {
		return regionSpec{}, false
	}
	w, err1 := strconv.Atoi(strings.TrimSpace(dim[:xIdx]))
	h, err2 := strconv.Atoi(strings.TrimSpace(dim[xIdx+1:]))
	if err1 != nil || err2 != nil {
		return regionSpec{}, false
	}
	fields := strings.Fields(strings.TrimSpace(parts[1]))
	counts := make([]int, len(fields))
	for j, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return regionSpec{}, false
		}
		counts[j] = n
	}
	return regionSpec{w: w, h: h, counts: counts}, true
}

func gridToShape(rows []string) []point {
	var cells []point
	for r, row := range rows {
		for c, ch := range row {
			if ch == '#' {
				cells = append(cells, point{r, c})
			}
		}
	}
	return normalize(cells)
}

func normalize(cells []point) []point {
	if len(cells) == 0 {
		return cells
	}
	minR, minC := cells[0].r, cells[0].c
	for _, p := range cells[1:] {
		if p.r < minR {
			minR = p.r
		}
		if p.c < minC {
			minC = p.c
		}
	}
	out := make([]point, len(cells))
	for i, p := range cells {
		out[i] = point{p.r - minR, p.c - minC}
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].r != out[j].r {
			return out[i].r < out[j].r
		}
		return out[i].c < out[j].c
	})
	return out
}

func shapeKey(cells []point) string {
	var b strings.Builder
	for _, p := range cells {
		fmt.Fprintf(&b, "%d,%d;", p.r, p.c)
	}
	return b.String()
}

// rot90CW rotates (r,c) 90° clockwise around origin: (r,c) -> (c, -r).
func rot90CW(p point) point { return point{p.c, -p.r} }

func flipH(p point) point { return point{p.r, -p.c} }

func transformCells(cells []point, f func(point) point) []point {
	out := make([]point, len(cells))
	for i, p := range cells {
		out[i] = f(p)
	}
	return normalize(out)
}

func uniqueOrientations(base []point) [][]point {
	seen := map[string]struct{}{}
	var out [][]point
	try := func(cells []point) {
		k := shapeKey(cells)
		if _, ok := seen[k]; ok {
			return
		}
		seen[k] = struct{}{}
		cp := make([]point, len(cells))
		copy(cp, cells)
		out = append(out, cp)
	}
	cur := base
	for t := 0; t < 4; t++ {
		try(cur)
		cur = transformCells(cur, rot90CW)
	}
	cur = transformCells(base, flipH)
	for t := 0; t < 4; t++ {
		try(cur)
		cur = transformCells(cur, rot90CW)
	}
	return out
}

type regionSpec struct {
	w, h   int
	counts []int
}

func regionFits(reg regionSpec, orientations [][][]point) bool {
	var pieces []int
	var totalCells int
	for si, cnt := range reg.counts {
		if si >= len(orientations) {
			if cnt > 0 {
				return false
			}
			continue
		}
		for k := 0; k < cnt; k++ {
			pieces = append(pieces, si)
			totalCells += len(orientations[si][0])
		}
	}
	if totalCells > reg.w*reg.h {
		return false
	}
	sort.Slice(pieces, func(i, j int) bool {
		si := len(orientations[pieces[i]][0])
		sj := len(orientations[pieces[j]][0])
		if si != sj {
			return si > sj
		}
		return pieces[i] < pieces[j]
	})

	grid := make([][]bool, reg.h)
	for r := range grid {
		grid[r] = make([]bool, reg.w)
	}
	return packPieces(grid, pieces, 0, orientations)
}

func packPieces(grid [][]bool, pieces []int, idx int, orientations [][][]point) bool {
	if idx == len(pieces) {
		return true
	}
	h, w := len(grid), len(grid[0])
	sid := pieces[idx]
	for _, orient := range orientations[sid] {
		maxR, maxC := 0, 0
		for _, p := range orient {
			if p.r > maxR {
				maxR = p.r
			}
			if p.c > maxC {
				maxC = p.c
			}
		}
		for r0 := 0; r0 <= h-1-maxR; r0++ {
			for c0 := 0; c0 <= w-1-maxC; c0++ {
				if canPlace(grid, orient, r0, c0) {
					doPlace(grid, orient, r0, c0, true)
					if packPieces(grid, pieces, idx+1, orientations) {
						return true
					}
					doPlace(grid, orient, r0, c0, false)
				}
			}
		}
	}
	return false
}

func canPlace(grid [][]bool, cells []point, r0, c0 int) bool {
	for _, p := range cells {
		r, c := r0+p.r, c0+p.c
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] {
			return false
		}
	}
	return true
}

func doPlace(grid [][]bool, cells []point, r0, c0 int, v bool) {
	for _, p := range cells {
		grid[r0+p.r][c0+p.c] = v
	}
}
