package main

import "fmt"

func testKuangKuang(p []string){
	p = append(p, "asdf")
}
func testKuangKuangStar(p *[]string){
	*p = append(*p, "asdf")
}

func testMap(p map[string]string){
	p["asdf"] = "jlkk"
}

func main() {
	var kuangkuang []string
	fmt.Println(kuangkuang)

	testKuangKuang(kuangkuang)
	fmt.Println(kuangkuang)

	testKuangKuangStar(&kuangkuang)
	fmt.Println(kuangkuang)

	m := make(map[string]string)
	fmt.Println(m)

	testMap(m)
	fmt.Println(m)

}
