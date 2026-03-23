package aoc2025

import (
	"sort"
	"strconv"
	"strings"
)

type pgPoint struct{ x, y, z int }

type pgEdge struct {
	i, j int
	dist int
}

func parsePlayground(boxes []string) ([]pgPoint, []pgEdge) {
	points := make([]pgPoint, len(boxes))
	for i, line := range boxes {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		points[i] = pgPoint{x, y, z}
	}

	n := len(points)
	edges := make([]pgEdge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[i].x - points[j].x
			dy := points[i].y - points[j].y
			dz := points[i].z - points[j].z
			edges = append(edges, pgEdge{i, j, dx*dx + dy*dy + dz*dz})
		}
	}

	sort.Slice(edges, func(a, b int) bool {
		return edges[a].dist < edges[b].dist
	})

	return points, edges
}

type pgUF struct {
	parent []int
	rank   []int
	sz     []int
	comps  int
}

func newPgUF(n int) *pgUF {
	parent := make([]int, n)
	sz := make([]int, n)
	for i := range parent {
		parent[i] = i
		sz[i] = 1
	}
	return &pgUF{parent: parent, rank: make([]int, n), sz: sz, comps: n}
}

func (uf *pgUF) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// union returns true if a real merge happened (different components).
func (uf *pgUF) union(a, b int) bool {
	ra, rb := uf.find(a), uf.find(b)
	if ra == rb {
		return false
	}
	if uf.rank[ra] < uf.rank[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	uf.sz[ra] += uf.sz[rb]
	if uf.rank[ra] == uf.rank[rb] {
		uf.rank[ra]++
	}
	uf.comps--
	return true
}

func Playground(boxes []string, connections int) int {
	_, edges := parsePlayground(boxes)
	n := len(boxes)
	uf := newPgUF(n)

	limit := connections
	if limit > len(edges) {
		limit = len(edges)
	}
	for i := 0; i < limit; i++ {
		uf.union(edges[i].i, edges[i].j)
	}

	seen := make(map[int]bool)
	var sizes []int
	for i := 0; i < n; i++ {
		r := uf.find(i)
		if !seen[r] {
			seen[r] = true
			sizes = append(sizes, uf.sz[r])
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	product := 1
	for i := 0; i < 3; i++ {
		if i < len(sizes) {
			product *= sizes[i]
		}
	}

	return product
}

func Playground2(boxes []string) int {
	points, edges := parsePlayground(boxes)
	n := len(points)

	if n <= 1 {
		return 0
	}

	uf := newPgUF(n)

	for _, e := range edges {
		if uf.union(e.i, e.j) && uf.comps == 1 {
			return points[e.i].x * points[e.j].x
		}
	}

	return 0
}
