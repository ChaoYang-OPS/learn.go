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
	fmt.Println(mathScore, mathScore2)
	mathScore2 = mathScore
	fmt.Println(mathScore2)
	fmt.Println(getScoreOfStudent(""))
	vg := &voteGame{
		students: []*student{
			&student{
				name: fmt.Sprintf("%d", 1),
			},
			&student{
				name: fmt.Sprintf("%d", 2),
			},
			&student{
				name: fmt.Sprintf("%d", 3),
			},
			&student{
				name: fmt.Sprintf("%d", 4),
			},
			&student{
				name: fmt.Sprintf("%d", 5),
			},
		},
	}
	leader := vg.goRun()
	fmt.Println(leader)
	leader.Distribute()

	var stdTF = &student{
		name: "TF",
	}
	var ldTF Leader = Leader(*stdTF)
	ldTF.Distribute()

	bytesTest1 := []byte{}
	var str1 string = string(bytesTest1)
	fmt.Println(str1)
}

func getScoreOfStudent(name string) (Math, Chinese, English) {
	return 0, 0, 0
}

type voteGame struct {
	students []*student
}

// 嵌套对象定义（继承）方式来定义班长
//type Leader struct {
//	student
//}
func (l *Leader) Distribute() {
	fmt.Println("send homework")
}

// 使用类型重定义
type Leader student

type FooTestFuncRedefine func()
type FooTestFuncRedefine1 []string // 类型可以是任意的

func (f *FooTestFuncRedefine) test111() {

}

func (g *voteGame) goRun() *Leader {
	// todo ....
	for _, item := range g.students {
		randInt := rand.Int()
		if randInt%2 == 0 {
			item.voteA(g.students[randInt%len(g.students)])
		} else {
			item.voteA(g.students[randInt%len(g.students)]) // todo 用随机数代替
		}
	}
	maxScore := -1
	maxScoreIndex := -1
	for i, item := range g.students {
		if maxScore < item.agree {
			maxScore = item.agree
			maxScoreIndex = i
		}
	}
	if maxScoreIndex >= 0 { // 判断是否大于0，因为没有学生，那么index就是默认值-1
		return (*Leader)(g.students[maxScoreIndex]) // 强制转换
	}
	return nil
}

//type Leader = student
type student struct {
	name     string
	agree    int
	disagree int
}

func (std *student) voteA(target *student) {
	target.agree++
}

func (std *student) voteD(target *student) {
	target.disagree++
}
