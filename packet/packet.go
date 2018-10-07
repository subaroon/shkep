package packet

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	Length int
	Name   string
}
type Packet struct {
	Fields   []Field
	ByteSize int
}

func NewPacket(fi []Field, bs int) *Packet {
	pck := Packet{fi, bs}
	return &pck
}
func (pck *Packet) writeFirstline() {
	fmt.Printf("|")
	byteStr := ""
	for i := 0; i < 32; i++ {
		if (i < 10) {
			byteStr = strings.Repeat(" ", pck.ByteSize/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.ByteSize-(pck.ByteSize/2-1)-1)
		} else {
			byteStr = strings.Repeat(" ", pck.ByteSize/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.ByteSize-(pck.ByteSize/2-1)-2)
		}
		fmt.Printf("%s", byteStr)
		fmt.Printf("|")
	}
	fmt.Printf("\n")
}
func (pck *Packet) writeFreeFields() {
	cnt := 0
	fmt.Printf("|")
	for _, v := range pck.Fields {
		for i := 0; i < v.Length; i++ {
			if i == v.Length-1 {
				fmt.Printf("%s", strings.Repeat(" ", pck.ByteSize)+"|")
			} else {
				fmt.Printf("%s", strings.Repeat(" ", pck.ByteSize+1))
			}
			cnt += 1
			if 32 == cnt {
				fmt.Printf("\n|")
				cnt = 0
			}
		}
	}
}
func (pck *Packet) writerFields() {
	cnt := 0
	fmt.Printf("|")
	fieldStr := ""
	for _, v := range pck.Fields {
		fieldStr = strings.Repeat(" ", (v.Length*(pck.ByteSize+1)-len(v.Name))/2) + v.Name + strings.Repeat(" ", (v.Length*(pck.ByteSize+1) - (v.Length*(pck.ByteSize+1)-len(v.Name))/2))
		for i := 0; i < (v.Length * (pck.ByteSize + 1)); i += (pck.ByteSize + 1) {
			if (i + pck.ByteSize + 1) == (v.Length * (pck.ByteSize + 1)) {
				fmt.Printf(fieldStr[i : i+pck.ByteSize])
				fmt.Printf("|")
			} else {
				fmt.Printf(fieldStr[i : i+pck.ByteSize+1])
			}
			cnt++
			if cnt == 32 {
				fmt.Printf("\n|")
				cnt = 0
			}
		}
	}
	fmt.Printf("\b")
}
func (pck *Packet) Show() {
	pck.writeFirstline()
	pck.writerFields()
}