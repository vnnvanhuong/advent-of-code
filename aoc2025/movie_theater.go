package aoc2025

import (
	"strconv"
	"strings"
)

// MovieTheater calculates the maximum area of a rectangle
// formed by any two red tiles as opposite corners.
func MovieTheater(input []string) int {
	type point struct{ x, y int }
	var points []point

	for _, line := range input {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, point{x, y})
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			dx := p1.x - p2.x
			if dx < 0 {
				dx = -dx
			}
			dy := p1.y - p2.y
			if dy < 0 {
				dy = -dy
			}

			area := (dx + 1) * (dy + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

// MovieTheater2 calculates the maximum area of a rectangle
// formed by any two red tiles as opposite corners, such that
// the entire rectangle is inside the rectilinear polygon formed by the tiles.
func MovieTheater2(input []string) int {
	type point struct{ x, y int }
	var points []point

	for _, line := range input {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, point{x, y})
	}

	isInside := func(x, y float64) bool {
		// Check if point is on any boundary edge
		for i := 0; i < len(points); i++ {
			p1 := points[i]
			p2 := points[(i+1)%len(points)]

			if float64(p1.x) == x && float64(p2.x) == x {
				minY := float64(p1.y)
				maxY := float64(p2.y)
				if p1.y > p2.y {
					minY, maxY = maxY, minY
				}
				if y >= minY && y <= maxY {
					return true
				}
			}
			if float64(p1.y) == y && float64(p2.y) == y {
				minX := float64(p1.x)
				maxX := float64(p2.x)
				if p1.x > p2.x {
					minX, maxX = maxX, minX
				}
				if x >= minX && x <= maxX {
					return true
				}
			}
		}

		intersections := 0
		for i := 0; i < len(points); i++ {
			p1 := points[i]
			p2 := points[(i+1)%len(points)]

			if p1.x == p2.x && float64(p1.x) > x {
				minY := float64(p1.y)
				maxY := float64(p2.y)
				if p1.y > p2.y {
					minY, maxY = maxY, minY
				}
				if y >= minY && y < maxY {
					intersections++
				}
			}
		}
		return intersections%2 == 1
	}

	isValidRectangle := func(p1, p2 point) bool {
		minX := float64(p1.x)
		maxX := float64(p2.x)
		if p1.x > p2.x {
			minX, maxX = maxX, minX
		}
		minY := float64(p1.y)
		maxY := float64(p2.y)
		if p1.y > p2.y {
			minY, maxY = maxY, minY
		}

		centerX := (minX + maxX) / 2.0
		centerY := (minY + maxY) / 2.0

		if !isInside(centerX, centerY) {
			return false
		}

		for i := 0; i < len(points); i++ {
			e1 := points[i]
			e2 := points[(i+1)%len(points)]

			if e1.y == e2.y { // Horizontal edge
				y := float64(e1.y)
				if y > minY && y < maxY {
					edgeMinX := float64(e1.x)
					edgeMaxX := float64(e2.x)
					if e1.x > e2.x {
						edgeMinX, edgeMaxX = edgeMaxX, edgeMinX
					}
					if edgeMaxX > minX && edgeMinX < maxX {
						return false
					}
				}
			} else { // Vertical edge
				x := float64(e1.x)
				if x > minX && x < maxX {
					edgeMinY := float64(e1.y)
					edgeMaxY := float64(e2.y)
					if e1.y > e2.y {
						edgeMinY, edgeMaxY = edgeMaxY, edgeMinY
					}
					if edgeMaxY > minY && edgeMinY < maxY {
						return false
					}
				}
			}
		}
		return true
	}

	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			dx := p1.x - p2.x
			if dx < 0 {
				dx = -dx
			}
			dy := p1.y - p2.y
			if dy < 0 {
				dy = -dy
			}

			area := (dx + 1) * (dy + 1)
			if area > maxArea {
				if isValidRectangle(p1, p2) {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
