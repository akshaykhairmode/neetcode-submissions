type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	builder := strings.Builder{}
	for _, v := range strs {
		length := len(v)
		builder.WriteString(fmt.Sprintf("%d#", length))
		builder.WriteString(v)
	}
	return builder.String()
}

func getJump(encoded string, currIndex int) (string, int) {
	jump := ""
	for currIndex < len(encoded) {
		if encoded[currIndex] == '#' {
			return jump, currIndex
		}
		jump += string(encoded[currIndex])
		currIndex++
	}
	return jump, currIndex
}

func (s *Solution) Decode(encoded string) []string {
	outArr := []string{}
	currIndex := 0
	for currIndex < len(encoded) {
		jump, updatedIndex := getJump(encoded, currIndex)
		jumpInt, err := strconv.Atoi(jump)
		if err != nil {
			break
		}

		// jump index is at '#', string starts at updatedIndex + 1
		start := updatedIndex + 1
		end := start + jumpInt
		
		extractedValue := encoded[start:end]
		outArr = append(outArr, extractedValue)
		currIndex = end
	}
	return outArr
}