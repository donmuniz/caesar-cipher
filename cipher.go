package main

type cipher struct {
	pass      string
	startChar rune
	endChar   rune
}

func getShiftSideValue(decode bool) rune {
	operation := rune(1)
	if decode {
		operation = -1
	}

	return operation
}

func process(msg, pass string, startChar, endChar rune, decode bool) string {
	totalChar := endChar - startChar + 1
	processedMsg := ""
	iteractions := 0

	for _, char := range msg {
		passIndex := iteractions % len(pass)
		var decodedChar rune

		if char > endChar || char < startChar {
			decodedChar = char
		} else {
			msgShift := char - startChar
			passShift := rune(pass[passIndex]) - startChar
			shiftSide := getShiftSideValue(decode)

			shiftValue := msgShift + passShift*shiftSide
			decodedChar = (shiftValue+totalChar)%totalChar + startChar

			iteractions += 1
		}

		processedMsg += string(decodedChar)
	}

	return processedMsg
}

func (c cipher) Encode(msg string) string {
	return process(msg, c.pass, c.startChar, c.endChar, false)
}

func (c cipher) Decode(msg string) string {
	return process(msg, c.pass, c.startChar, c.endChar, true)
}
