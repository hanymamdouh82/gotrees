// Copyright 2023 Hany Mamdouh. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package gotrees

type Person struct {
	Name string
	Age  int
	Boss string
}

var (
	developer1 = Node[Person]{
		Id: "5",
		Data: Person{
			Name: "Zaher",
			Age:  25,
		},
	}

	developer2 = Node[Person]{
		Id: "4",
		Data: Person{
			Name: "Amr",
			Age:  24,
		},
		Children: []*Node[Person]{&developer5},
	}

	developer4 = Node[Person]{
		Id: "44",
		Data: Person{
			Name: "Jebril",
			Age:  32,
		},
	}

	developer5 = Node[Person]{
		Id: "5",
		Data: Person{
			Name: "Adham",
			Age:  12,
		},
	}

	developer3 = Node[Person]{
		Id: "3",
		Data: Person{
			Name: "Doaa",
			Age:  37,
		},
	}

	teamleader1 = Node[Person]{
		Id: "2",
		Data: Person{
			Name: "Mezo",
			Age:  40,
		},
		Children: []*Node[Person]{&developer1, &developer2, &developer4},
	}

	teamleader2 = Node[Person]{
		Id: "1",
		Data: Person{
			Name: "Hager",
			Age:  38,
		},
		Children: []*Node[Person]{&developer3},
	}

	boss = Node[Person]{
		Id: "0",
		Data: Person{
			Name: "Hany",
			Age:  41,
		},
		Children: []*Node[Person]{&teamleader1, &teamleader2},
	}

	rawData = []Person{
		{
			Name: "Amr",
			Age:  24,
			Boss: "Mezo",
		},
		{
			Name: "Zaher",
			Age:  25,
			Boss: "Mezo",
		},
		{
			Name: "Jebril",
			Age:  31,
			Boss: "Mezo",
		},
		{
			Name: "Doaa",
			Age:  37,
			Boss: "Hager",
		},
		{
			Name: "Hager",
			Age:  38,
			Boss: "Hany",
		},
		{
			Name: "Mezo",
			Age:  40,
			Boss: "Hany",
		},
		{
			Name: "Hany",
			Age:  41,
		},
	}
)
