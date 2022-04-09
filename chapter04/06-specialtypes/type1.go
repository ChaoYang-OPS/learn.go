package main

import (
	"fmt"
)

type Math = int
type English = int
type Chinese = int
func main() {
	var mathScore int = 100
	var mathScore2 Math = 100
	fmt.Println(mathScore,mathScore2)
	mathScore2 = mathScore
	fmt.Println(mathScore2)
	fmt.Println(getScoreOfStudent(""))
}

func getScoreOfStudent(name string) (Math,Chinese,English) {
	return 0,0,0
}


type voteGame struct {
	students []*student
}

func (g *voteGame) goRun() *leader {
	// todo ....
	for _,item := range g.students{
		item.voteA(g.students[0])  // todo 用随机数代替
	}
	maxScore := -1
	maxScoreIndex := -1
	for i,item := range  g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 {  // 判断是否大于0，因为没有学生，那么index就是默认值-1
		return g.students[maxScoreIndex]
	}
	return nil
}

type leader = student
type student struct {
	agree int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}