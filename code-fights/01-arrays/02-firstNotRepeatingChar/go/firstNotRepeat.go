package main

func main() {

}

func firstNotRepeatingCharacter(s string) string {
	fnrc := '_'
	chars := map[byte]int{}
	fi := len(s)

	for i := range s {
		if _, ok := chars[s[i]]; ok {
			chars[s[i]] = -1
		} else {
			chars[s[i]] = i
		}
	}
	for k, v := range chars {
		if v != -1 && v < fi {
			fnrc = rune(k)
			fi = v
		}
	}
	return string(fnrc)
}
