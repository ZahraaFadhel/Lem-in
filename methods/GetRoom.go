package lem

import "strings"

// Returns a pointer to the Room with a key
func GetRoom(graph *Graph, k string) *Room {
	k = strings.Trim(k, "\r\n")
	for i, v := range graph.Rooms {
		if v.Key == k {
			return graph.Rooms[i]
		}
	}
	return nil
}
