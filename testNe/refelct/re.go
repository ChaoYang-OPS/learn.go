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

func valueIsEmpty() {
	var ifc interface{}
	v := reflect.ValueOf(ifc)
	fmt.Printf("v 持有真实的值 %t  \n", v.IsValid()) // false

	ifc = 8
	v = reflect.ValueOf(ifc)
	fmt.Printf("v 持有真实的值 %t\n", v.IsValid()) // true

	var user *User = nil
	v = reflect.ValueOf(user)
	if v.IsValid() {
		// v 持有值是 nil, true
		// 调用 v.IsNil() 之前必须保证 v.IsValid() True
		fmt.Printf("v 持有值是 nil, %t\n", v.IsNil())
	} else {
		fmt.Printf("v 没有有值\n")
	}
	var u User
	v = reflect.ValueOf(u)
	if v.IsValid() {
		// 调用 v.IsZero() 之前必须保证 v.IsValid() True
		// v 持有值是默认值, true
		fmt.Printf("v 持有值是默认值, %t\n", v.IsZero())
	} else {
		fmt.Printf("v 没有有值\n")
	}
}

// 可寻址
func addressable() {
	v1 := reflect.ValueOf(1)
	var x int
	v2 := reflect.ValueOf(x)
	v3 := reflect.ValueOf(&x)
	v4 := v3.Elem()
	fmt.Printf("v1 is addressable %t\n", v1.CanAddr()) // false
	fmt.Printf("v2 is addressable %t\n", v2.CanAddr()) // false
	fmt.Printf("v3 is addressable %t\n", v3.CanAddr()) // false
	fmt.Printf("v4 is addressable %t\n", v4.CanAddr()) // true

	slice := make([]int, 3, 5)
	v5 := reflect.ValueOf(slice)
	v6 := v5.Index(0)
	// 切片的value不可寻址
	// 切片中某个元素的value可寻址
	fmt.Printf("v6 is addressable %t\n", v6.CanAddr()) // true
	fmt.Printf("v5 is addressable %t\n", v5.CanAddr()) // false
	mp := make(map[int]bool, 5)
	v7 := reflect.ValueOf(mp)
	fmt.Printf("v7 is addressable %t\n", v7.CanAddr()) // false
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

func changeValue() {
	var i int = 10
	fmt.Printf("address of i is %p\n", &i) // 0xc0000b2010
	iValue := reflect.ValueOf(&i)
	if iValue.CanAddr() {
		iValue.SetInt(8)
		fmt.Printf("iValue i=%d\n", i)
	}
	iValue2 := iValue.Elem()
	if iValue2.CanAddr() {
		iValue2.SetInt(8)
		fmt.Printf("iValue2 i=%d\n", i)        // iValue2 i=8
		fmt.Printf("address of i is %p\n", &i) // 0xc0000b2010
	}
	// 通过反射修改值，需要传指针,如果要修改成功前提必须是可寻值
	var s string = "hello"
	sValue := reflect.ValueOf(&s)
	sValue.Elem().SetString("go")
	fmt.Println(s) // go

	user := User{
		Name: "TF",
		Age:  19,
		Addr: "",
	}
	userValue := reflect.ValueOf(&user)
	//fmt.Println(userValue.CanAddr()) // false
	// 未导出成员，不可以CanSet
	userValue.Elem().FieldByName("Age").SetInt(29) // FieldByName 返回的Field是可寻址的
	fmt.Println(user.Age)                          // 29
}
func changeSlice() {
	users := make([]*User, 3, 5)
	users[0] = &User{
		Name: "TF",
		Age:  20,
		Addr: "",
	}
	sliceValue := reflect.ValueOf(users)
	if sliceValue.Len() > 0 {
		sliceValue.Index(0).Elem().FieldByName("Age").SetInt(29)
		fmt.Println(users[0].Age) // 29
	}
	sliceValue.Index(1).Set(reflect.ValueOf(&User{
		Name: "SRE",
		Age:  10,
		Addr: "",
	}))
	fmt.Println(users[1].Age) // 10
	sliceValue2 := reflect.ValueOf(&users)
	// 修改成4,容量只能小于初始的容量,
	sliceValue2.Elem().SetCap(4)
	fmt.Println(cap(users)) // 4
	sliceValue2.Elem().SetLen(2)
	fmt.Println(len(users)) // 2
}
func changeMap() {
	u1 := &User{
		Name: "T1",
		Age:  20,
		Addr: "",
	}
	u2 := &User{
		Name: "T2",
		Age:  20,
		Addr: "",
	}
	userMap := make(map[int]*User, 5)
	userMap[1] = u1

	mapValue := reflect.ValueOf(userMap)
	mapType := mapValue.Type()
	// 获得map key的type
	keyType := mapType.Key()
	// 获得map value的type
	valueType := mapType.Elem()
	// type of key int type of value *main.User
	fmt.Printf("type of key %v type of value %v\n", keyType, valueType)
	// set key value
	mapValue.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf(u2))
	uValue := mapValue.MapIndex(reflect.ValueOf(1))
	uValue.Elem().FieldByName("Age").SetInt(50)
	for k, v := range userMap {
		fmt.Printf("%d %s %d\n", k, v.Name, v.Age)
	}
}

func callFunction() {
	valueFunc := reflect.ValueOf(Add)
	typeFunc := valueFunc.Type()
	argInNum := typeFunc.NumIn()
	args := make([]reflect.Value, argInNum)
	for i := 0; i < argInNum; i++ {
		if typeFunc.In(i).Kind() == reflect.Int {
			args[i] = reflect.ValueOf(i)
		}
	}
	results := valueFunc.Call(args)
	if typeFunc.Out(0).Kind() == reflect.Int {
		i := results[0].Interface().(int)
		fmt.Printf("sum is %d\n", i)
	}
}

func callMethod() {
	u1 := &User{
		Name: "T1",
		Age:  20,
		Addr: "",
	}
	uValue := reflect.ValueOf(u1)
	// MethodByName 不能获得未导出的方法
	getNameMetod := uValue.MethodByName("GetName")
	results := getNameMetod.Call([]reflect.Value{})
	i := results[0].Interface().(string)
	fmt.Printf("%s\n", i) // T1
}

func newStruct() {
	t := reflect.TypeOf(User{})
	value := reflect.New(t)
	value.Elem().FieldByName("Name").SetString("SRE")
	user := value.Interface().(*User)
	fmt.Println(user.Name) // SRE
}

func newSlice() {
	var slices []User
	sliceType := reflect.TypeOf(slices)
	sliceValue := reflect.MakeSlice(sliceType, 3, 5)
	sliceValue.Index(0).Set(reflect.ValueOf(User{
		Name: "FS",
		Age:  30,
		Addr: "",
	}))
	users := sliceValue.Interface().([]User)
	for _, u := range users {
		fmt.Println(u.Name)
	}
	fmt.Println(users[0].Age)
}

func newMap() {
	var userMap map[int]*User
	mapType := reflect.TypeOf(userMap)
	mapValue := reflect.MakeMapWithSize(mapType, 5)
	u2 := &User{
		Name: "T2",
		Age:  20,
		Addr: "",
	}
	mapValue.SetMapIndex(reflect.ValueOf(2), reflect.ValueOf(u2))
	mp := mapValue.Interface().(map[int]*User)
	for k, v := range mp {
		// 2 T2
		fmt.Printf("%d %s\n", k, v.Name)
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
	//valueOtherConversion()
	//valueIsEmpty()
	//addressable()
	////changeValue()
	//changeSlice()
	//changeMap()
	//callFunction()
	callMethod()
	newStruct()
	newMap()
}
