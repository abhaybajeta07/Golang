package encrypt

func Nimbus(str string) string{
	encryptedStr := ""
	for _, c =: range str {
		asciiCode := int(c)
		charachter := string(asciiCode + 3)
		encryptedStr += charachter
	}
	return encryptedStr
}