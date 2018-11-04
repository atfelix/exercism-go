package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	if hasDuplicateRecords(records) {
		return nil, fmt.Errorf("duplicate records")
	}
	if !areRecordsValid(records) {
		return nil, fmt.Errorf("records are not valid")
	}

	sort.Slice(records, func(i, j int) bool {
		return (records[i].Parent < records[j].Parent ||
					(records[i].Parent == records[j].Parent && records[i].ID <= records[j].ID))
	})
	
	bucketsOfRecords := make([][]Record, len(records))

	for _, record := range records[1:] {
		bucketsOfRecords[record.Parent] = append(bucketsOfRecords[record.Parent], record)
	}

	nodesOfRecords := make([]*Node, len(records))

	for index := range records {
		if len(bucketsOfRecords[index]) == 0 {
			nodesOfRecords[index] = &Node{ID: index}
		} else {
			nodesOfRecords[index] = &Node{ID: index, Children:[]*Node{}}
		}
	}

	for index, node := range nodesOfRecords {
		for _, child := range bucketsOfRecords[index] {
			node.Children = append(node.Children, nodesOfRecords[child.ID])
		}
	}
	
	return nodesOfRecords[0], nil
}

func hasDuplicateRecords(records []Record) bool {
	m := map[Record]bool{}
	for _, record := range records {
		if _, ok := m[record]; ok {
			return true 
		}
		m[record] = true
	}
	return false
}

func areRecordsValid(records []Record) bool {
	for _, record := range records {
		if record.ID < 0 || record.ID >= len(records) {
			return false
		}
		if record.ID == 0 && record.Parent == 0 {
			continue
		}
		if record.Parent >= record.ID {
			return false
		}
	}
	return true
}
