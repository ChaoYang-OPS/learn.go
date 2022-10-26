package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name,omitempty"` //16 字节
	Age  int    `json:"age,omitempty"`
	Addr string `json:"addr"`
}

func (u User) GetName() string {
	return u.Name
}
func (u User) setName() string {
	return u.Name
}
func RefType() {
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("Hello world")
	typeT := reflect.TypeOf(true)
	fmt.Println(typeI.String())
	fmt.Println(typeS.String())
	fmt.Println(typeI.Kind())
	fmt.Println(typeS.Kind())
	fmt.Println(typeT)
	u1 := User{}
	useType := reflect.TypeOf(u1)
	fmt.Println(useType)
	fmt.Println(useType.Kind() == reflect.Struct) // struct
	fmt.Println(useType.String())
	u2 := &User{}
	use2Type := reflect.TypeOf(u2)
	fmt.Println(use2Type)
	fmt.Println(use2Type.Kind() == reflect.Ptr) //  ptr
	fmt.Println(use2Type.String())

	typeUser3 := use2Type.Elem() // 解析指针
	fmt.Println(typeUser3)
	fmt.Println(typeUser3.Kind() == reflect.Struct) //   struct
	fmt.Println(typeUser3.String())

	fmt.Println(useType.Name())
	fmt.Println(useType.PkgPath())
	fmt.Println(useType.Size()) // 24
}

func getField() {
	u1 := User{}
	typeUser := reflect.TypeOf(u1)
	fileNum := typeUser.NumField()
	for i := 0; i < fileNum; i++ {
		field := typeUser.Field(i)
		fmt.Println(i, field.Name,
			field.Anonymous, field.Tag.Get("json"),
			field.Offset, field.Type,
			field.IsExported())
	}
	if nameFiled, ok := typeUser.FieldByName("Name"); ok {
		fmt.Println("+++++++++++", nameFiled.IsExported())
	}
}

func Add(a, b int) int {
	return a + b
}
func getFunc() {
	typeFunc := reflect.TypeOf(Add)
	fmt.Printf("Kind %v\n", typeFunc.Kind() == reflect.Func)
	argInNum := typeFunc.NumIn()   // 输入参数的个数
	argOutNum := typeFunc.NumOut() // 输出参数的个数
	for i := 0; i < argInNum; i++ {
		argType := typeFunc.In(i)
		fmt.Println("IN", argType)
	}
	for i := 0; i < argOutNum; i++ {
		argType := typeFunc.Out(i) // 输出参数
		fmt.Println("OUT", argType)
	}
}

type Plane struct {
	Name string
}
type Bird struct {
	N string
}
type Frog struct {
	N string
}

func getStructRelation() {
	plan := Plane{}
	bird := Bird{}
	frog := Frog{}
	planType := reflect.TypeOf(plan)
	birdType := reflect.TypeOf(bird)
	frogType := reflect.TypeOf(frog)
	_ = frogType

	fmt.Println(birdType.AssignableTo(planType))  // false
	fmt.Println(birdType.ConvertibleTo(planType)) // false
}
func memolyAligin() {
	type C struct {
		sex    bool
		age    int64
		weight uint16
	}
	t := reflect.TypeOf(C{})
	fmt.Println(t.Size())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%s %d\n", field.Name, field.Offset)
	}

}

func getMethod() {
	typeUser := reflect.TypeOf(User{})
	methodNum := typeUser.NumMethod() // 未导出方法，不包含在内
	for i := 0; i < methodNum; i++ {
		method := typeUser.Method(i)
		fmt.Printf("name: %s type %s export %t\n", method.Name,
			method.Type, method.IsExported())
	}
}

func valueOtherConversion() {
	iValue := reflect.ValueOf(1)
	sValue := reflect.ValueOf("Hello")
	userValue := reflect.ValueOf(User{
		Name: "TF",
		Age:  10,
		Addr: "AD",
	})
	fmt.Println(iValue)
	fmt.Println(sValue)
	fmt.Println(userValue)
	// value converse to type
	iType := iValue.Type()
	sType := sValue.Type()
	userType := userValue.Type()
	fmt.Println(iValue.Kind() == iType.Kind(), iType.Kind() == reflect.Int)
	fmt.Println(sValue.Kind() == sType.Kind(), sType.Kind() == reflect.String)
	fmt.Println(userValue.Kind() == userType.Kind(), userType.Kind(), userType.Kind() == reflect.Struct)

	userValue2 := &User{}
	u2 := reflect.ValueOf(userValue2).Elem() // 解析指针
	fmt.Println(u2.Kind() == reflect.Struct)
	fmt.Printf("orgin value is %d   %d\n", iValue.Interface().(int), iValue.Int())
	fmt.Printf("orgin svalue is %s   %s\n", sValue.Interface().(string), sValue.String())

	user := userValue.Interface().(User)
	fmt.Println(user.Age, user.Name)
}
func main() {
	//RefType()
	//getField()
	//memolyAligin()
	//getMethod()
	//getFunc()
	//getStructRelation()
	valueOtherConversion()
}
