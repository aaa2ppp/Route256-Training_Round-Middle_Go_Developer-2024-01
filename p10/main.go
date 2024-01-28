package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// func solution(...) (...) {
// 	// TODO
// }

type treeNode struct {
	id     int
	text   string
	childs []*treeNode
}

type writer interface {
	io.Writer
	// io.ByteWriter
	// io.StringWriter
}

func writeTree(out writer, indent []byte, root *treeNode) {
	fmt.Fprintf(out, "%s%s\n", indent, root.text)
	writeChilds(out, indent, root.childs)
}

func writeChilds(out writer, indent []byte, childs []*treeNode) {
	for i := 0; i < len(childs)-1; i++ {
		writeChild(out, indent, childs[i])
	}
	if i := len(childs) - 1; i >= 0 {
		writeLastChild(out, indent, childs[i])
	}
}

func writeChild(out writer, indent []byte, node *treeNode) {
	fmt.Fprintf(out, "%s|\n", indent)
	fmt.Fprintf(out, "%s|--%s\n", indent, node.text)
	writeChilds(out, append(indent, "|  "...), node.childs)
}

func writeLastChild(out writer, indent []byte, node *treeNode) {
	fmt.Fprintf(out, "%s|\n", indent)
	fmt.Fprintf(out, "%s|--%s\n", indent, node.text)
	writeChilds(out, append(indent, "   "...), node.childs)
}

func task(in *bufio.Reader, out writer) error {
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return err
	}

	// read forest

	forest := make(map[int]*treeNode, n+1)

	for i := 0; i < n; i++ {
		// read node
		var id, parentId int
		var text string
		if _, err := fmt.Fscan(in, &id, &parentId); err != nil {
			return err
		}
		// skip one space
		if _, err := in.ReadByte(); err != nil {
			return err
		}
		text, err := in.ReadString('\n')
		if err != nil {
			return err
		}
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSuffix(text, "\r")

		// get parent node
		parent := forest[parentId]
		if parent == nil {
			parent = &treeNode{id: parentId}
			forest[parentId] = parent
		}

		// get node
		node := forest[id]
		if node == nil {
			node = &treeNode{id: id}
			forest[id] = node
		}

		node.text = text
		parent.childs = append(parent.childs, node)
	}

	// sort childs by id for all nodes
	for _, node := range forest {
		childs := node.childs
		sort.Slice(childs, func(i, j int) bool {
			return childs[i].id < childs[j].id
		})
	}

	// write all trees of forest
	indent := make([]byte, 0, 256)
	for _, tree := range forest[-1].childs {
		writeTree(out, indent, tree)
		fmt.Fprintln(out)
	}

	return nil
}

func run(in io.Reader, out io.Writer) (err error) {
	bIn := bufio.NewReader(in)
	buf := &bytes.Buffer{}

	var t int
	if _, err := fmt.Fscan(bIn, &t); err != nil {
		return err
	}
	// skip eol
	if _, err := bIn.ReadString('\n'); err != nil {
		return err
	}

	for i := 0; i < t; i++ {
		if err := task(bIn, buf); err != nil {
			return fmt.Errorf("task%d: %w", i+1, err)
		}
	}

	if buf.Len() > 0 {
		// cut off the last delimiting blank line
		_, err = out.Write(buf.Bytes()[:buf.Len()-1])
	}

	return err
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")

	if err := run(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
