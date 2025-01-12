package main

import "fmt"

func main() {
	/*
		map crud

			create map via map() function
				variable := make(map[keyType]valueType)

				variable := map[keyType]valueType{
					k1:v1,
					k2:v2,
				}

			update map
				variable[key1] = value1

			get map value
				v1, ok := variable[key1]
				if ok {
					code block
				}

			delete map data
				delete(map1,key1)


		map loop
			for k, v := range variable {
				code block
			}

		map slice
			variable := make([]map[ktype]vtype, len, cap)

		slice map
			variable := make(map[ktype][]slicetype)

		map is reference type
			use copy() function to copy map

	*/
	userInfo1 := make(map[string]string)
	userInfo1["username"] = "Larry"
	userInfo1["sex"] = "male"
	userInfo1["hobby"] = "golang"

	fmt.Println(userInfo1)
	for k, v := range userInfo1 {
		fmt.Println(k, ":", v)
	}

	delete(userInfo1, "hobby")
	fmt.Println(userInfo1)

	userInfos := make([]map[string]string, 3, 3)
	userInfos[0] = userInfo1

	userInfo2 := map[string]string{
		"userName": "pitiao",
		"sex":      "male",
		"hobby":    "play",
	}

	userInfos[1] = userInfo2

	fmt.Println(userInfos)

	mapSlice := make(map[string][]string)
	mapSlice["hobbies"] = []string{
		"game",
		"eating",
		"game",
	}
	mapSlice["names"] = []string{
		"xiaowang",
		"xiaohong",
	}
	fmt.Println(mapSlice)

}
