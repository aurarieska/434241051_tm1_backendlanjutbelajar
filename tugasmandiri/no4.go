package main

import "fmt"

type Student struct {
	ID int			`json: "id"`
	Name string		`json: "nama"`
	Grade float64	`json: "kelas"`
	IsActive bool	`json: "status"`
}
var siswa1 = Student{
	ID : 1,
	Name : "Aura",
	Grade : 100,
	IsActive : false,
}

func GetInfo() string {
	info := fmt.Sprintln(siswa1.ID, siswa1.Name, siswa1.Grade, siswa1.IsActive)
	return info
}

func UpdateGrade(grade float64){
	siswa1.Grade = grade
}

func activate(){
	siswa1.IsActive = true
}

func deactive(){
	siswa1.IsActive = false
}