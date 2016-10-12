package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var smpl1 *sample1

type sample1 struct {
	Key1 string
	Key2 string
}

var smpl2 *sample2

type sample2 struct {
	Key1 []string
}

var smpl3 *sample3

type sample3 struct {
	SKey1 SKey1
}

type SKey1 struct {
	Key1a string
	Key1b string
}

func main() {
	doSample1()
	doSample2()
	doSample3()
}

func doSample1() {
	log.Println("----- doSample1 -----")
	file, _ := ioutil.ReadFile("sample1.json")
	err := json.Unmarshal(file, &smpl1)
	if err == nil {
		log.Println(smpl1)
		log.Println(smpl1.Key1)
		log.Println(smpl1.Key2)
	} else {
		log.Println(err)
	}
}

func doSample2() {
	log.Println("----- doSample2 -----")
	file, _ := ioutil.ReadFile("sample2.json")
	err := json.Unmarshal(file, &smpl2)
	if err == nil {
		log.Println(smpl2)
		log.Println(smpl2.Key1)
	} else {
		log.Println(err)
	}
}

func doSample3() {
	log.Println("----- doSample3 -----")
	file, _ := ioutil.ReadFile("sample3.json")
	err := json.Unmarshal(file, &smpl3)
	if err == nil {
		log.Println(smpl3)
		log.Println(smpl3.SKey1)
		log.Println(smpl3.SKey1.Key1a)
		log.Println(smpl3.SKey1.Key1b)
	} else {
		log.Println(err)
	}
}
