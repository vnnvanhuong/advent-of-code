package aoc2025

import (
	"strings"
)

const reactorStart = "you"
const reactorSink = "out"

// ParseReactorGraph builds an adjacency list from puzzle lines like "name: a b c".
func ParseReactorGraph(lines []string) map[string][]string {
	g := make(map[string][]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		colon := strings.Index(line, ":")
		if colon < 0 {
			continue
		}
		name := strings.TrimSpace(line[:colon])
		rest := strings.TrimSpace(line[colon+1:])
		if rest == "" {
			g[name] = nil
			continue
		}
		g[name] = strings.Fields(rest)
	}
	return g
}

// Reactor returns the number of distinct directed paths from "you" to "out".
// The graph is assumed to be acyclic on paths that matter; a cycle reachable
// from you causes a panic.
func Reactor(g map[string][]string) int64 {
	memo := make(map[string]int64)
	onStack := make(map[string]bool)

	var dfs func(string) int64
	dfs = func(v string) int64 {
		if v == reactorSink {
			return 1
		}
		if onStack[v] {
			panic("reactor: cycle on path from you")
		}
		if n, ok := memo[v]; ok {
			return n
		}
		onStack[v] = true
		var total int64
		for _, w := range g[v] {
			total += dfs(w)
		}
		onStack[v] = false
		memo[v] = total
		return total
	}

	return dfs(reactorStart)
}

// ReactorFromLines parses lines and counts paths from you to out.
func ReactorFromLines(lines []string) int64 {
	return Reactor(ParseReactorGraph(lines))
}
