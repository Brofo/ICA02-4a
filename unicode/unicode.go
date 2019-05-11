package unicode

const islandsk = "\x6e\x6f\x72\xc3\xb0\x75\x72\x6f\x67\x73\x75\xc3\xb0\x75\x72"
const japansk = "\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97"

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	if language == "jp" {
		return japansk
	} else if language == "is" {
		return islandsk
	} else {
		return "Error: velg jp (japansk) eller is (islandsk)"
	}
}
