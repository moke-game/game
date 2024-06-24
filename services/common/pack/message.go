package pack

type Message struct {
	ID      uint32 // Message ID
	DataLen uint32 // Message length
	Data    []byte // Message content

}
