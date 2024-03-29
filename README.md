# shkep

shkep is a tool to draw packet structure on CLI

## Usage

### Example: TCP packet

#### code

```
package main

import "github.com/subaruf/shkep/packet"

func main() {
  tcpFields := []packet.Field{
    packet.Field{16, "src port"},
    packet.Field{16, "dst port"},
    packet.Field{32, "seq number"},
    packet.Field{32, "ack number"},
    packet.Field{4, "offset"},
    packet.Field{6, "unused"},
    packet.Field{6, "ctrl flag"},
    packet.Field{16, "window size"},
    packet.Field{16, "checksum"},
    packet.Field{16, "urgent pointer"},
  }
  byteWidthForRendering := 4
  tcpPacket := packet.NewPacket(tcpFields, byteWidthForRendering)
  tcpPacket.Show()
}

```

#### Result

```
|0  |1  |2  |3  |4  |5  |6  |7  |8  |9  |10 |11 |12 |13 |14 |15 |16 |17 |18 |19 |20 |21 |22 |23 |24 |25 |26 |27 |28 |29 |30 |31 |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
|                            src port                           |                            dst port                           |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
|                                                           seq number                                                          |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
|                                                           ack number                                                          |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
|     offset    |         unused        |       ctrl flag       |                          window size                          |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
|                            checksum                           |                         urgent pointer                        |
|---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---|
```
