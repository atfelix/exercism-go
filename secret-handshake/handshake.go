package secret

const (
	wink = 1 << iota
	doubleBlink
	closeYourEyes
	jump
	reverse
)

var secretHandshakeMap = map[uint]string{
	wink: "wink",
	doubleBlink: "double blink",
	closeYourEyes: "close your eyes",
	jump: "jump",
}

func Handshake(code uint) []string {
	handshake := []string{}
	for i := uint(1); i < reverse; i <<= 1 {
		if code & i != 0 {
			handshake = append(handshake, secretHandshakeMap[i])
		}
	}
	if code & reverse != 0 {
		for i, j := 0, len(handshake) - 1; i < j; i, j = i + 1, j - 1 {
			handshake[i], handshake[j] = handshake[j], handshake[i]
		}
	}
	return handshake
}