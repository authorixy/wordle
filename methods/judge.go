package methods

type JudgeInfo int8

const (
	InWordButIncorrectPosition JudgeInfo = 1
	InCorrectCompletely		   JudgeInfo = 2
	Correct					   JudgeInfo = 3
)

func Judge(input, answer []byte) (bool, []JudgeInfo) {
	set := make(map[byte]struct{}, len(input))
	for _, c := range answer {
		set[c] = struct{}{}
	}
	res := make([]JudgeInfo, len(input))
	judgeResult := true
	for idx, c := range input {
		if c == answer[idx] {
			res[idx] = Correct
		} else if _, ok := set[c]; ok {
			res[idx] = InWordButIncorrectPosition
			judgeResult = false
		} else {
			res[idx] = InCorrectCompletely
			judgeResult = false
		}
	}
	return judgeResult, res
}

