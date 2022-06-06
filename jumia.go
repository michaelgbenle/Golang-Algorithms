package main

import "sort"

//1) A retail company maintains the data of its customers in the CUSTOMER table.
//Write a query to print the IDs and NAMEs of the Customers, sorted by Customer.Name
//in descending order. if two or more customers have the same CUSTOMER.NAME, then
//sort these by CUSTOMER.ID in ascending order
//select id , name from customer order by name desc, id

//2) A travel and tour company has two tables relating to customers
//each tour offers a discount if a minimum number of people book at the same time.
//Write a query to print the maximum number of  discounted tours any one family in the families table can choose from.
//SOLUTION
//select max (travel_count.tour) from (select (select count(*)
//as num from countries c where c.min_size <= f.family_size)
//as tour from families f) as travel_count

//3
//Given an array of long integers 'arr' and a number 'num'
//iterate through the elements in the arr and
//double the value of num whenever an element equals num.
//find the maximum possible value of num,
//Knowing that arr can be reordered before iteration.
func doubleSizeArr(arr []int64, b int64) int64 {
	num := b
	sort.SliceStable(arr, func(a, b int) bool {
		return arr[a] < arr[b]
	})

	for _, a := range arr {
		if a == num {
			num = num + 2
		}
	}
	return num
}

//there is a shop with an old style cash register.
//Rather than scanning items and
//the price from the database
//the price of each item is typed in manually
//the method sometimes lead to errors
//given a list of items and their correct prices
//compare the prices to those
//entered when each item was sold
//determine the number of errors in selling prices
func priceCheck(products []string, productPrices []float32, productSold []string, soldPrice []float32) int32 {
	originalProducts := make(map[string]float32)
	var errors int32
	for i, product := range products {
		originalProducts[product] = productPrices[i]
	}
	for i, product := range productSold {
		originalPrice := originalProducts[product]
		if soldPrice[i] != originalPrice {
			errors++
		}
	}
	return errors
}

// select distinct p.Name, c.Name from schedule s left join Professor p on
//p.id = s.professor_id left join Course c on c.id = s.course_id inner join
//Department d on d.id = p.department_id

func closedPaths(number int32) int32 {
	var sum int32
	for number > 0 {
		if number%10 == 0 || number%10 == 4 || number%10 == 6 || number%10 == 9 {
			sum += 1
		}
		number /= 10
	}
	return sum
}
