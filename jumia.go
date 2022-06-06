package main

import "sort"

//1
//select id , name from customer order by name desc, id

//2
//select max (travel_count.tour) from (select (select count(*)
//as num from countries c where c.min_size <= f.family_size)
//as tour from families f) as travel_count

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

func priceCheck(products []string, productPrices []float32, productSold []string, soldPrice []float32) int32 {
	original

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
