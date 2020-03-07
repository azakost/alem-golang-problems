package main

import "fmt"

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

func IsAnagram(str1, str2 string) bool {
	t := byte(0)
	for _, x := range str1 + str2 {
		if x != ' ' {
			t ^= byte(x)
		}
	}
	if t == 0 {
		return true
	}
	return false
}

// func IsAnagram(str1, str2 string) bool {
// 	a, b := clean(str1), clean(str2)
// 	if a == b {
// 		return true
// 	}
// 	return false
// }

// func clean(s string) string {
// 	s = sort(s)
// 	if s[0] == ' ' {
// 		return clean(s[1:])
// 	}
// 	return s
// }

// func sort(s string) string {
// 	l := len(s)
// 	ar := make([]string, l)
// 	for i := 0; i < l; i++ {
// 		ar[i] = s[i : i+1]
// 	}

// 	for x := 0; x < len(s); x++ {
// 		for y := x + 1; y < len(s); y++ {
// 			if ar[x] > ar[y] {
// 				tmp := ar[x]
// 				ar[x] = ar[y]
// 				ar[y] = tmp
// 			}
// 		}
// 	}
// 	txt := ""
// 	for _, a := range ar {
// 		txt += a
// 	}
// 	return txt
// }
