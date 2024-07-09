package lem

import (
	"fmt"
	"strings"
)

// Adds an edge to the graph
func AddEdge(graph *Graph, from, to string) error {
	// get Room
	fromRoom := GetRoom(graph, from)
	toRoom := GetRoom(graph, to)

	// error handling
	if (fromRoom == nil || toRoom == nil || fromRoom == toRoom) {
		err := fmt.Errorf("😞 ERROR: invalid edge(%v --> %v)", strings.Trim(from, "\r\n"), strings.Trim(to, "\r\n"))
		return err
	} else if (RoomExists(fromRoom.Children, to)) {
		err := fmt.Errorf("😞 ERROR: existing edge(%v --> %v)", from, to)
		return err
	} else {
		// add the edge
		fromRoom.Children = append(fromRoom.Children, toRoom)
		toRoom.Children = append(toRoom.Children, fromRoom)
	}
	return nil
}
