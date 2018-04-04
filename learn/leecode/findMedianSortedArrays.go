package main

import (
	"fmt"
	"math"
)
//难的一匹
func main(){
	var num1  = []int{1,1}
	var num2  = []int{1,2}
	fmt.Println(findMedianSortedArrays(num1, num2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len_short := len(nums1)
	len_long := len(nums2)



	if len_short > len_long{
		return findMedianSortedArrays(nums2,nums1)
	}

	var i, j, imin, imax = 0,0,0,len_short
	var m_left, m_right float64= 0, 0
	var half = (len_long+len_short+1)/2   //i + j == len_short - i +len_long  - j (or: m - i + n - j + 1)

	for imin<=imax{
		i = (imax + imin)/2
		j = half - i

		/*
		<a> (j == 0 or i == m or B[j-1] <= A[i]) and
		(i == 0 or j = n or A[i-1] <= B[j])
		Means i is perfect, we can stop searching.

		<b> j > 0 and i < m and B[j - 1] > A[i]
			Means i is too small, we must increase it.

		<c> i > 0 and j < n and A[i - 1] > B[j]
		Means i is too big, we must decrease it.
		*/
		if i>0  && nums1[i-1]>nums2[j]{  //i大了  i变小
			imax -= 1
		} else if  i<len_short && nums2[j-1]>nums1[i]{ //j大了 i变大
			imin += 1
		}else {
			if i == 0 {
				m_left = float64(nums2[j-1])
			}else if j == 0 {
				m_left = float64(nums1[i-1])
			}else {
				m_left = math.Max(float64(nums1[i-1]),float64(nums2[j-1]))
			}
			break
		}

	}



	if (len_short+len_long)%2==1{
		return m_left
	}


	if i == len_short{
		m_right = float64(nums2[j])
	}else if j == len_long{
		m_right = float64(nums1[i])
	}else{
		m_right = math.Min(float64(nums1[i]),float64(nums2[j]))
	}

	return  (m_left+m_right) / 2


}


