package util

func ReadBinary(fileName string) [][]bool {
	lines := ReadLines(fileName)

	res := make([][]bool, len(lines))

	for i, l := range lines {
		res[i] = make([]bool, len(l))
		for j, c := range l {
			if c == '1' {
				res[i][j] = true
			}
		}
	}
	return res
}
