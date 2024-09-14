package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"unsafe"
)

//멀티반환 함수
// 반환값을 타입으로만 선언한 경우

func Divide(a, b int) (int, bool) {

	if b == 0 {
		return 0, false
	}
	return a / b, true
}

//반환값에 변수와 타입을 넣는경우

func Divide2(a, b int) (result int, success bool) {

	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

//재귀 호출함수

func cur(n int) {

	if n == 0 {

		return
	}

	fmt.Println(n)
	cur(n - 1)
	fmt.Println("After", n)
}

func myAge() int {
	return 18
}

func main() {
	var a int = 10
	var msssage string = "hellow"
	b := "string" //선언 대입문 (var == :)
	c := 3.5

	d := a * int(c) // float타입인 c를 int로 타입변환후 연산!

	e := 1234.123
	f := 1234.123

	g := e * f

	var two int = 2
	var ten int = 10
	var fo = 1234523.1234

	fmt.Println(a, msssage, b, d, g)
	fmt.Printf("two:%d ten:%d fo: %f\n", two, ten, fo)

	// stdin := bufio.NewReader(os.Stdin)

	// var from int
	// var to int

	// n, err := fmt.Scanln(&from, &to)

	// if err != nil {
	// 	fmt.Println(err)
	// 	stdin.ReadString('\n')
	// } else {
	// 	fmt.Println(n, from, to)
	// }

	var x int8 = 4
	var y int8 = 127

	fmt.Printf("x: %d\n y:%d\n x+y: %v\n", x, y, x+y)
	fmt.Printf("x: %08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y: %08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)

	var nx = 0.1
	var ny = 0.2
	var nc = 0.3
	var nd int64 = 127

	fmt.Printf("nx: %08f\n ny: %08f\n nx+ny=nc:%v\n", nx, ny, math.Nextafter(nx+ny, nc) == nc)
	fmt.Printf("nd: %016b\n nd+1: %08b\n nd: %v\n", nd, nd+1, 127 < nd+1)

	// 멀티함수 호출

	value1, success := Divide(9, 3)
	fmt.Println(value1, success)

	value2, success := Divide(9, 0)
	fmt.Println(value2, success)

	value3, success := Divide2(4, 0)
	fmt.Println(value3, success)

	value4, success := Divide2(4, 2)
	fmt.Println(value4, success)

	//재귀함수 호출

	cur(3)

	switch age := myAge(); true {

	case age < 19:
		fmt.Println("나이가 너무 어립니다.")

	case age > 19:
		fmt.Println(("사용가능"))
	}

	for i := 0; i < 4; i++ {

		for j := 0; j < 1; j++ {
			fmt.Print(strings.Repeat(" ", 3-i) + strings.Repeat("*", i*2+1))

		}

		fmt.Println()
	}

	for i := 0; i < 7; i++ {

		for j := 0; j < 1; j++ {

			if i >= 4 {
				fmt.Print(strings.Repeat(" ", i-3) + strings.Repeat("*", i*-2+13))
			} else {
				fmt.Print(strings.Repeat(" ", 3-i) + strings.Repeat("*", i*2+1))
			}
		}

		fmt.Println()
	}

	// 배열

	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	days := [3]string{"monday", "tuesday", "wednesday"}
	fmt.Println(days)

	const Y int = 4
	ages := [Y]int{12, 23, 43}
	fmt.Println(ages)

	students := [...]int{1, 2, 3}
	fmt.Println(students)

	//배열순회

	nums2 := [8]int{6: 1, 7: 1}
	nums2[5] = 1

	fmt.Println(nums2)

	for i := 0; i < len(nums2); i++ {
		fmt.Println(nums2[i])
	}

	// range를 활용한 배열 순회
	// 인텍스와 값으 전부 가져온다.

	nums3 := [3]int{10, 11, 12}

	for idex, value := range nums3 {

		fmt.Println(idex, value)
	}

	//배열의 값만 가져오고 싶을때

	for _, value := range nums3 {

		fmt.Println(value)
	}

	/*
		<중요>
		-배열은 연속된 메모리다.

		-배열 요소 찾아가기
		 요소위치 = 배열 시작 주소 +(인덱스 x 타입크기)

		 자료구조를 선택할때는 신중하게 해야한다.
		 random access가 자주 일어나는 작업을 할땐는 여러 자료구조중 배열이 가장빠르다
		 이유는 배열은 연속된 메모리이기 때문이다.
	*/

	//배열 복사
	// 배열복사시 반드시 값과 공간의 타입크기가 일치해야한다.

	a1 := [5]int64{1, 2, 3, 4, 5}

	for i, v := range a1 {
		fmt.Printf("a1[%d] = %d\n", i, v)
	}

	b1 := [5]int64{100, 200, 300, 400, 500}

	for i, v := range b1 {
		fmt.Printf("b1[%d] = %d\n", i, v)
	}

	b1 = a1

	for i, v := range b1 {

		fmt.Printf("b1[%d] = %d\n", i, v)
	}

	// * int와 int64는 크기는 같지만 서로 다른 타입이다.

	//다중배열 / 순회

	arrays := [2][5]int{ //5개짜리 int배열이 2개 있다.
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}, // 감싸는 중괄호가 같은 라인에서 닫히지 않으면 반드시 콤마를 입력한다.
	}

	for _, array := range arrays {
		for _, v := range array {
			fmt.Print(v, " ")
		}
		fmt.Println()

	}

	//**구조체** advance

	//구조체 선언 첫번째 단계는 타입을 선언한다. type [변수명] structure{}

	//구조체 변수 초기화
	var winter Aespa // 전부 기본값으로 초기화

	//전부 기본값임으로 값을 넣어 초기화한다.
	winter.Name = "김민정"
	winter.Age = 23
	winter.State = "Korea"

	fmt.Printf("%v\n", winter)
	fmt.Printf("이름:%s 나이:%d 국적:%s\n ", winter.Name, winter.Age, winter.State)

	//미리 값을 초기화 하는법
	var karina Aespa = Aespa{"유지민", 24, "korea"}
	fmt.Printf("이름:%s 나이:%d 국적:%s\n", karina.Name, karina.Age, karina.State)

	//특정 값만 초기화 하는법
	var ning Aespa = Aespa{Name: "ning", Age: 22}
	fmt.Println(ning)

	vip := Vip{
		User{"winter", 23},
		"골드",
		1000,
	}

	fmt.Printf("이름: %s 나이: %d 등급: %s 누적사용금액: %d만원\n",
		vip.UserInfo.Name,
		vip.UserInfo.age,
		vip.VipLevel,
		vip.Price,
	)

	// 위에 방식을 embadded field 방식으로 해보자

	vip2 := Vip2{
		User{"김민정", 23},
		"플레티넘",
		2000,
	}

	// 이런게 vip2. 으로 한번에 User에 접근 할 수 있다.
	fmt.Printf("이름: %s 나이: %d 등급: %s 누적사용금액: %d만원\n",
		vip2.Name,
		vip2.age,
		vip2.VipLevel,
		vip2.Price,
	)

	//** 메모리 패딩을 줄이려면 8바이트보다 작은 필드는 8바이트 크기를 고려해서 배치하자.

	//포인터
	//- 메모리 주소를 값으로 갖는 타입

	v1 := 10

	var p *int // 이 코드의 의미는 "p라는 변수는 포인터 변수로서 int값을 가지고, 다른 변수의 메모리 주소를 값으로 받을수있다."를 의미한다.
	p = &v1

	fmt.Printf("p가 참조하고 있는 v1의 메모리주소: %p\n", p)
	fmt.Printf("p가 참조하고 있는 v1의 메모리 공간에 저장된 값: %d\n", *p)

	*p = 20

	fmt.Println(*p)
	fmt.Println(v1)

	var kim Winter

	changeValue(&kim)
	fmt.Printf("이름: %s 나이: %d\n", kim.Name, kim.Age)

	//*** 구조체를 포인터 변수로 선언하면 포인터 변수 메모리에 구조체의 메모리 주소값을 복사함으로 메모리가 훨씬 이득이다
	// 기본적으로 메모리의 주소값은 64비트 컴퓨터에서는 8바이트, 32비트 컴퓨터에서는 4바이트이다.
	// 만약 포인터 변수로 지정하지 않고 복사를하면 구조체의 메모리 주소값이 아닌 메모리 전체의 값을 복사한다

	var kimMin *Aespa = &Aespa{"winter", 23, "ko"}

	fmt.Printf("이름: %s 나이:%d 국적:%s\n", kimMin.Name, kimMin.Age, kimMin.State)

}

//----------------------------------------------------------------------------

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func changeValue(arg *Winter) {
	arg.Name = "김민정"
	arg.Age = 23

	//-----------------------------------------------------------------------------

	str := "amyon 아이"

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])

	// } // 한글은 따로 공간에 저장됨으로 순회가 되도 이상한 값이 나옴

	// 많이 사용하는 영어,숫자는 utf-8(1~4바이트)에서 1바이트 공간에 저장된다.
	// 20%가 어떠한 것의 80%를 차지한다는 논리 즉 영어,숫자가 전체의 80%를 차지 함으로 1바이트 공간에 저장하고 나머지는 나머지 공간에 저장

	//한글까지 순회하는법
	// []rune을 사용한다.  [8]string 이건 길이 8을 가지는 문자열 배열이하는 뜻인데 []이런식으로 비워두면 슬라이스라고해서 동적배열을 의미한다.
	// rune은 int32의 별칭이다.

	// arr := []rune(str)

	// for i := 0; i < len(arr); i++ {

	// 	fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	// }

	// 더 간단한 방식은 배열에서 배운 range를 사용한다.

	for i, v := range str {

		fmt.Printf("타입:%T 값:%d 문자값:%c 인덱스:%d\n", v, v, v, i)
	}

	str2 := "Hello"

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("Data:%d Len:%d\n", stringHeader1.Data, stringHeader1.Len)

	fmt.Println(toUpper1(str2))
	fmt.Println(toUpper2(str2))
}

//----------------------------------------------------------------------------------

func toUpper1(str string) string {

	var res string

	for _, v := range str {

		if v >= 'a' && v <= 'z' {
			res += string('A' + (v - 'a'))
		} else {
			res += string(v)
		}
	}
	return res
}

func toUpper2(str string) string {
	var builder strings.Builder

	for _, v := range str {

		if v >= 'a' && v <= 'z' {

			builder.WriteRune('A' + (v - 'a'))
		} else {
			builder.WriteRune(v)
		}
	}

	return builder.String()
}

//구조체 타입선언

type Aespa struct {
	Name  string
	Age   int
	State string
}

type Winter struct {
	Name string
	Age  int
}

// 구조체를 필드로 같는 구조체
type User struct {
	Name string
	age  int
}

type Vip struct {
	UserInfo User
	VipLevel string
	Price    int
}

// 위에 방식을 embadded field 방식으로 해보자

type Vip2 struct {
	User     // 이렇게 필드값에 구조체 자체를 넣어준다.
	VipLevel string
	Price    int
}

/*
rune -int32
byte - unit8

go는 최강타입임으로 연산시 타입이 정확하게 일치해야한다.
타입이 다르다면 타입변환을 해야한다.


서식
%d - 정수타입
%f -실수타임
\n- 줄바꿈

*/
