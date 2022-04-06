package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	abck := "abc(defghij)kl"
	bbck := "abc(def（gh)ij)kl"
	cbck := "abc(defgh)i(jkl)"

	listA, _ := obtainCtt(abck)
	for i, value := range listA {
		fmt.Println("i:", i, " val:", value)
	}
	listB, _ := obtainCtt(bbck)
	for i, value := range listB {
		fmt.Println("i:", i, " val:", value)
	}
	listC, _ := obtainCtt(cbck)
	for i, value := range listC {
		fmt.Println("i:", i, " val:", value)
	}

}

func leftBracketIdx(ctt string) (idx int, extra int) {
	idx = -1
	extra = -1
	firstBraket := strings.Index(ctt, "(")
	if firstBraket >= 0 {
		idx = firstBraket
		extra = 1
	}
	//中文字符下标需要加3
	firstBraketZ := strings.Index(ctt, "（")
	if firstBraketZ >= 0 && firstBraketZ < firstBraket {
		idx = firstBraketZ
		extra = 3
	}
	return
}

func rightBracketIdx(ctt string) (idx, extra int) {
	idx = -1
	firstBraket := strings.Index(ctt, ")")
	if firstBraket >= 0 {
		idx = firstBraket
		extra = 1
	}
	//中文字符下标需要加3
	firstBraketZ := strings.Index(ctt, "）")
	if firstBraketZ >= 0 && firstBraketZ < firstBraket {
		idx = firstBraketZ
		extra = 3
	}
	return
}

func obtainLeft(ctt string) (leftList, leftExtraList []int) {
	for {
		leftIdx, extra := leftBracketIdx(ctt)
		if leftIdx < 0 {
			break
		}
		leftList = append(leftList, leftIdx)
		leftExtraList = append(leftExtraList, extra)
		ctt = ctt[leftIdx+extra:]
	}
	return
}
func obtainRight(ctt string) (rightList, rightExtraList []int) {
	for {
		rightIdx, extra := rightBracketIdx(ctt)
		if rightIdx < 0 {
			break
		}
		rightList = append(rightList, rightIdx)
		rightExtraList = append(rightExtraList, extra)
		ctt = ctt[rightIdx+extra:]
	}
	return
}

func obtainCtt(ctt string) (cttList []string, err error) {
	leftList, LeftExtraList := obtainLeft(ctt)
	rightList, rightExtraList := obtainRight(ctt)
	if len(leftList) != len(rightList) {
		return nil, errors.New("括号不对等")
	}
	if len(leftList) == 0 {
		return nil, err
	}
	leftList = append(leftList, -1)
	rightList = append(rightList, -1)

	idxList := make([]int, 0)
	idxExtraList := make([]int, 0)
	// todo 括号数量不对等不处理
	// todo 取出i然后对比
	leftIdx := 0
	rightIdx := 0
	for leftIdx < len(leftList) {
		tmpIdx := -1
		tmpExtraIdx := 0
		for rightIdx < len(rightList) {
			if leftList[leftIdx] >= 0 && rightList[rightIdx] >= 0 {
				if leftList[leftIdx] < rightList[rightIdx] {
					tmpIdx = leftList[leftIdx]
					tmpExtraIdx = LeftExtraList[leftIdx]
					leftIdx++
				} else {
					tmpIdx = rightList[rightIdx]
					tmpExtraIdx = rightExtraList[rightIdx]
					rightIdx++
				}
			} else if leftList[leftIdx] >= 0 {
				tmpIdx = leftList[leftIdx]
				tmpExtraIdx = LeftExtraList[leftIdx]
				leftIdx++
			} else if rightList[rightIdx] >= 0 {
				tmpIdx = rightList[rightIdx]
				tmpExtraIdx = rightExtraList[rightIdx]
				rightIdx++
			} else {
				tmpIdx = -1
				rightIdx++
				leftIdx++
			}

			if tmpIdx < 0 {
				continue
			}
			idxList = append(idxList, tmpIdx)
			idxExtraList = append(idxExtraList, tmpExtraIdx)
		}
	}

	//recordIdx := idxList[0]
	//recordIdx := idxList[0]
	for i := 1; i < len(idxList); i++ {
		cttList = append(cttList, ctt[idxList[i-1]+idxExtraList[i-1]:idxList[i]])
	}
	return
}

func obtainCtt3(ctt string) (cttList []string, err error) {
	recordRightIdx, _ := rightBracketIdx(ctt)
	if recordRightIdx < 0 {
		return cttList, errors.New("right bracket is not exist")
	}

	leftIdx := -1
	tmpCtt := ctt
	for {
		left, extra := leftBracketIdx(ctt)
		if left < 0 {
			break
		}
		if leftIdx < 0 {
			leftIdx = left + extra
			continue
		}
		if left > recordRightIdx {
			break
		}
		cttList = append(cttList, tmpCtt[leftIdx:left])
		leftIdx = left + extra
	}
	if leftIdx == -1 {
		return cttList, errors.New("left bracket is not exist")
	}

	right, _ := obtainRight(ctt)
	if len(right) == 0 {
		return cttList, errors.New("right bracket is not exist")
	}

	//cttList = append(cttList, ctt[leftBracket:rightBracket])
	//ctt = ctt[splitIdx:]

	return
}

func obtainCtt2(ctt string) (cttList []string, err error) {
	for {
		//中文字符下标需要加3
		firstBraket := strings.Index(ctt, "(")
		firstBraketZ := strings.Index(ctt, "（")
		firstRight := strings.Index(ctt, ")")
		firstRightZ := strings.Index(ctt, "）")
		//if firstRight < 0 && firstRightZ < 0 {
		//	return cttList, errors.New("right bracket is not exist")
		//}

		leftBracket := -1
		if firstBraket >= 0 {
			if firstBraket < firstBraketZ {
				leftBracket = firstBraket + 1
			} else {
				leftBracket = firstBraketZ + 3
			}
		} else if firstBraketZ >= 0 {
			leftBracket = firstBraketZ + 3
		}
		if leftBracket == -1 {

		}

		rightBracket := 0
		splitIdx := 0
		if firstRight >= 0 {
			if firstRight < firstRightZ {
				rightBracket = firstRight
				splitIdx = firstRight + 1
			} else if firstRightZ >= 0 {
				rightBracket = firstRightZ
				splitIdx = firstRightZ + 3
			}
		} else {
			rightBracket = firstRightZ
			splitIdx = firstRightZ + 3
		}
		cttList = append(cttList, ctt[leftBracket:rightBracket])
		ctt = ctt[splitIdx:]
	}

	return
}
