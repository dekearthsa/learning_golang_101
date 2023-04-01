// เมื่อเราสร้าง folder เก็บไฟล์ให้เรียกใช้ package ตามชื่อ folder
package controller

// หากต้องการให้ค่าของ function add สามารถเรียกใช้ด้านนอกได้เราจำเป็นต้องกำหนดให้ ตัวหน้าเป็น พิมพ์ใหญ่ //
func Add(numberX, numberY int) int {
	return numberX + numberY
}

func PopOut(numberX, numberY int) int {
	return numberX - numberY
}
