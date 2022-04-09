package main

import (
	"fmt"
	"math/rand"
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
	vg := &voteGame{
		students: []*student{
			&student{
				name: fmt.Sprintf("%d",1),
			},
			&student{
				name: fmt.Sprintf("%d",2),
			},
			&student{
				name: fmt.Sprintf("%d",3),
			},
			&student{
				name: fmt.Sprintf("%d",4),
			},
			&student{
				name: fmt.Sprintf("%d",5),
			},
			&student{
				name: fmt.Sprintf("%d",6),
			},
			&student{
				name: fmt.Sprintf("%d",7),
			},
		},
	}
	leader := vg.goRun()
	fmt.Println(leader)
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
		rnadInt := rand.Int()
		if rnadInt %2 == 0 {
			item.voteA(g.students[rnadInt%len(g.students)])
		} else {
			item.voteA(g.students[rnadInt%len(g.students)])  // todo 用随机数代替
		}
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
	name string
	agree int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}