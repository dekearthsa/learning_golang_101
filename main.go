package main

import (
	"fmt"
	"gobasic/controller" // ไฟล์ที่เราเรียกจะมาจาก folder ฝั่งของ controll // 
)

func main() {
	fmt.Println("------------")
	fmt.Println("hello world")
	fmt.Println("------------")
	//  เรียกใช้ function ใน main //
	variableLearning()
	operatorMath()
	operatorCompare()

	calGrade := calculateScore(100)
	fmt.Println("grade =>  ", calGrade)
	mapMethodfunc()
	loopingFunction()
	typeOfNumber, value := retrunTwoValueInFunction(30, 25)
	fmt.Println("Type of number =", typeOfNumber, "sumation => ", value)

	numberInVaridic := summation(10, 20, 30, 40, 22, 33, 44)
	fmt.Println("variadic function sumation => ", numberInVaridic)
	storeArrayProduct()
	pushProductInArray()
	loopOutOfArrayProduct()
	fmt.Println("this from another package",controller.Add(23, 22))
}

// การกำหนดตัวแปรต่าง ๆ //
func variableLearning() {
	var name string = "variable name"
	var age int = 13
	var score float64 = 32.4
	var isPass bool = true
	// สามารถใช้ type inferance  := //
	nameInferance := "name earth"
	ageInferance := 12
	scoreInferance := 33.4
	isPassInferace := true
	// เปลี่ยน ค่าของ name เป็นค่าอื่น เราสามาถใช้ const เเทน //
	const isName string = "this variable will not change"

	fmt.Println("this is variable => " + name)
	fmt.Println("user age is =>", age)
	fmt.Println("score user is => ", score)
	fmt.Println("is pass => ", isPass)
	fmt.Println("name inferace =>", nameInferance)
	fmt.Println("age inferace =>", ageInferance)
	fmt.Println("score inferace =>", scoreInferance)
	fmt.Println("is pass inferace =>", isPassInferace)
	name = "Convert variable name to Earth"

	fmt.Println(name)
	fmt.Println(isName)

}

// การคำนวณ​เลข //
func operatorMath() {
	var result float64 = 0
	const numberX float64 = 10
	const numberY float64 = 5
	numberA, numberB := 11, 23
	result = numberY + numberY
	fmt.Println("the result is ", result)
	result = numberX - numberY
	fmt.Println("the result is ", result)
	result = numberX - numberY
	fmt.Println("the result is ", result)
	result = numberX / numberY
	fmt.Println("the result is ", result)
	result = numberX * numberY
	fmt.Println("the result is ", result)
	fmt.Println(numberA + numberB)
}

// การใช้ตัวเปรียบเทียบ //
func operatorCompare() {
	const stringX string = "Hello"
	const stringY string = "Hello"
	const stringZ string = "Word"

	const numberX float64 = 12.4
	const numberY float64 = 19.2

	fmt.Println("this is const X == Y", (stringX == stringY))
	fmt.Println("this is const X == Y", (stringX == stringZ))
	fmt.Println("this is const X == Y", (stringX == stringZ))

	fmt.Println("this is const number X == Y", (numberX >= numberY))
	fmt.Println("this is const number X == Y", (numberX <= numberY))
}

// การส่งตัวแปรเข้ามาใน function if else if else //
func calculateScore(score float64) string {
	if score > 90 {
		const grade string = "A"
		return grade
	} else if score <= 90 && score > 80 {
		const grade string = "B"
		return grade
	} else if score <= 80 && score > 70 {
		const grade string = "C"
		return grade
	} else if score <= 70 && score > 60 {
		const grade string = "D"
		return grade
	} else if score <= 60 && score > 50 {
		const grade string = "C"
		return grade
	} else {
		const grade string = "F"
		return grade
	}
}

// การใช้งาน map methods //
func mapMethodfunc() {

	arrayOfName := [3]string{"Aa", "Bb", "Cc"}
	dictOfArrayName := map[string]string{"th": "Thailand", "us": "United States", "uk": "United Kingdom"}
	emptyArrayOfDict := map[string]string{}

	fmt.Println("name of array => ", arrayOfName)
	fmt.Println("name of array 0 => ", arrayOfName[0])
	fmt.Println("name of array 1 => ", arrayOfName[1])

	fmt.Println("name of dict th => ", dictOfArrayName["th"])
	fmt.Println("name of dict us => ", dictOfArrayName["us"])
	fmt.Println("name of dict uk => ", dictOfArrayName["uk"])

	emptyArrayOfDict["eat"] = "กิน"
	emptyArrayOfDict["hello"] = "สวัสดี"
	emptyArrayOfDict["drink"] = "ดื่ม"

	fmt.Println("empty of dict eat => ", emptyArrayOfDict["eat"])
	fmt.Println("empty of dict hello => ", emptyArrayOfDict["hello"])
	fmt.Println("empty of dict drink => ", emptyArrayOfDict["drink"])

}

// for loop //
func loopingFunction() {
	numberArray := []int{10, 20, 30, 40, 50, 40, 30, 40}
	stringArray := []string{"Aa", "Bb", "Cc", "Dd", "Ee", "Ff", "Gg"}
	arrayWithKey := map[string]string{"k1": "AA", "k2": "BB", "k3": "DD", "k4": "EE"}
	for i := 0; i < len(numberArray); i++ {
		fmt.Println("loop el int => ", numberArray[i])
	}

	for i := 0; i < len(stringArray); i++ {
		fmt.Println("loop el string =>", stringArray[i])
	}

	// สำหรับการ loop ข้อมูลโดยข้อมูลเป็น key and value //
	for key, value := range arrayWithKey {
		println("this is key => ", key, ", this is value => ", value)
	}

	// กรณีไม่ใช้ key ให้ใช้ เครื่องหมาย _ เเทนการสร้างตัวแปร //
	for _, value := range arrayWithKey {
		println("ignore key => ", value)
	}

}

// การ คืนค่ากลับมามากกว่า 1 ค่า //
func retrunTwoValueInFunction(numberX, numberY int) (string, int) { //  (numberX, numberY int) ค่า input, (string, int) ระบุว่าค่าส่งกลับมี type เป็นอะไร
	total := numberX + numberY
	typeNumber := ""

	if total%2 == 0 {
		typeNumber = "even"
	} else {
		typeNumber = "odd"
	}

	return typeNumber, total // return กลับ ต้อง return กลับ 2 ค่า โดยตำแหน่งต้องตรงกับ (string, int) ที่เรา set ไว้บน function
}

// การสร้าง function ที่รับค่า parameter แบบ ไม่จำกัดจำนวน
// การ คืนค่ากลับมามากกว่า 1 ค่า //
func summation(numberX ...int) int {
	total := 0

	for i := 0; i < len(numberX); i++ {
		total = total + numberX[i]
	}

	// หรือ จะใช้ รูปแบบนี้สำหรับการ วน loop ก็ได้
	// for _, value := range numberX {
	// 	total = total + value
	// }

	return total
}

// การใช้งาน struct หรือ โครงสร้างข้อมูล จะใช้ก็ต่อเมื่อเราต้องการเก็บข้อมูลต่างชนิดใน array จำเป็นต้องใช้ struct //
type Product struct {
	name        string
	price       float64
	category    string
	isPromotion bool
	discount    int
}

func storeArrayProduct() {
	myproduct := Product{name: "pen", price: 12.0, category: "consumtion", isPromotion: false, discount: 12}
	fmt.Println("value in myproduct", myproduct)
	fmt.Println("call only name in myproduct => ", myproduct.name)
	fmt.Println("call only price in myproduct => ", myproduct.price)
}

// กรณีที่เราจะใส่ค่าของข้อมูลที่เรากำหนดเองเข้าไปใน array สามารถทำได้ตามตัวอย่างต่อไปนี้ //
func pushProductInArray() {
	arrayOfProduct := []Product{} // ประกาศค่าของ array ให้เก็บข้อมูลรูปแบบ ของ struct product ที่เราได้กำหนดไว้
	arrayOfProduct = append(
		arrayOfProduct,
		Product{name: "pen", price: 12.0, category: "consumtion", isPromotion: false, discount: 12})

	arrayOfProduct = append(
		arrayOfProduct,
		Product{name: "remover", price: 52.0, category: "consumtion", isPromotion: true, discount: 13})

	fmt.Println("we push data in product array", arrayOfProduct)
}

// เราจะ loop ข้อมูลที่อยู่ใน array ของ product  //
func loopOutOfArrayProduct() {
	arrayOfProduct := []Product{} // ประกาศค่าของ array ให้เก็บข้อมูลรูปแบบ ของ struct product ที่เราได้กำหนดไว้
	arrayOfProduct = append(
		arrayOfProduct,
		Product{name: "pen", price: 12.0, category: "consumtion", isPromotion: false, discount: 12})

	arrayOfProduct = append(
		arrayOfProduct,
		Product{name: "remover", price: 52.0, category: "consumtion", isPromotion: true, discount: 13})

	for _, value := range arrayOfProduct {
		fmt.Println("productname: ", value.name, ", product price: ", value.price)
	}

}
