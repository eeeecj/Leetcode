package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MergeSort([]int{7, 5, 6, 4}))
}

//比较类排序：通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn)，因此也称为非线性时间比较类排序。

//1、冒泡法(n^2)
//比较相邻的元素。如果第一个比第二个大，就交换他们两个。
//对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
//针对所有的元素重复以上的步骤，除了最后一个。
//持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

//2、快速排序（nlogn）
//从数列中挑出一个元素，称为 “基准”（pivot）;
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
//在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
func QuickSort(nums []int) []int {
	left, right := 0, len(nums)-1
	if len(nums) == 0 || left == right {
		return nums
	}
	//设定基准
	privot := nums[left]
	//保证基准元素左边比他小，右边比他大
	for left < right {
		//如果右边元素小于基准，则与左边元素start交换位置，此时的start为基准元素
		for left < right && nums[right] > privot {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		//如果左边元素小于基准，则与右边元素end交换位置，此时end为基准元素
		for left < right && nums[left] < privot {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		}
	}
	//递归调用，左边一定是left+1，保证每次递归都会减少数组长度，否则若left一直为0，则出现溢出情况。
	return append(QuickSort(nums[:left]), append([]int{nums[left]}, QuickSort(nums[left+1:])...)...)
}

//3、插入排序（n^2）
//将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
//从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
func InsertSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		for j := i - 1; j >= 0; j-- {
			//判断是否为大于nums[i],大于则向后挪动，否则就加入值
			if temp >= nums[j] {
				nums[j+1] = temp
				break
			}
			nums[j+1] = nums[j]
			nums[j] = temp
		}
	}
	return nums
}

//4、希尔排序
//选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；
//按增量序列个数 k，对序列进行 k 趟排序；
//每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
//仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。

func ShellSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	//将数组分为gap份，[i,i+gap,...]为一份
	gap := len(nums) / 2
	for gap > 0 {
		//判断gap中每份的大小，并进行挪动
		for i := gap; i < len(nums); i++ {
			j := i
			for j-gap >= 0 && nums[j] < nums[j-gap] {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
				j -= gap
			}
		}
		gap = gap / 2
	}
	return nums
}

//5、选择排序（O(n^2)）
//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
//再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//重复第二步，直到所有元素均排序完毕。

func SelectSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums)-1; i++ {
		min := i
		//每次选择最小的交换位置
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				nums[min], nums[j] = nums[j], nums[min]
			}
		}
	}
	return nums
}

//6、堆排序
//堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。
//堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
//将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
//将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
//由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，
//得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。

func HeapSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	nums = buildHeap(nums)
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		nums = append(heapify(nums[:i], 0), nums[i:]...)
	}
	return nums
}

//构造大顶堆
func buildHeap(nums []int) []int {
	//堆顶左节点跟右节点刚好相等
	for i := len(nums) / 2; i >= 0; i-- {
		nums = heapify(nums, i)
	}
	return nums
}

//判断节点是否大于子节点
func heapify(nums []int, i int) []int {
	left := i*2 + 1
	right := i*2 + 2
	large := i
	if left < len(nums) && nums[large] < nums[left] {
		large = left
	}
	if right < len(nums) && nums[large] < nums[right] {
		large = right
	}
	//交换元素
	if large != i {
		nums[large], nums[i] = nums[i], nums[large]
		//交换元素后可能不满足大顶堆，重新构建
		nums = heapify(nums, large)
	}
	return nums
}

//7、归并排序
//申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
//设定两个指针，最初位置分别为两个已经排序序列的起始位置；
//比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
//重复步骤 3 直到某一指针达到序列尾；
//将另一序列剩下的所有元素直接复制到合并序列尾。
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	return merge(MergeSort(nums[:mid]), MergeSort(nums[mid:]))
}
func merge(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	res := []int{}
	for len(nums1) != 0 && len(nums2) != 0 {
		if nums1[0] > nums2[0] {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		} else {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		}
	}
	if len(nums1) == 0 && len(nums2) != 0 {
		res = append(res, nums2...)
	}
	if len(nums2) == 0 && len(nums1) != 0 {
		res = append(res, nums1...)
	}
	return res
}

//8、计数排序
//花O(n)的时间扫描一下整个序列 A，获取最小值 min 和最大值 max
//开辟一块新的空间创建新的数组 B，长度为 ( max - min + 1)
//数组 B 中 index 的元素记录的值是 A 中某元素出现的次数
//最后输出目标整数序列，具体的逻辑是遍历数组 B，输出相应元素以及对应的个数

func CountSort(nums []int) []int {
	min, max := getMaxMin(nums)
	bitmap := make([]int, max-min+1)
	for _, v := range nums {
		bitmap[v-min]++
	}
	index := 0
	for i := 0; i < len(bitmap); i++ {
		for bitmap[i] > 0 {
			nums[index] = i + 1
			bitmap[i]--
			index++
		}
	}
	return nums
}
func getMaxMin(nums []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for _, v := range nums {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max
}

//9、桶排序
//桶排序是计数排序的升级版。这个是利用了函数的映射关系，是否高效就在于这个映射函数的确定。
//所以为了使桶排序更加高效，我们要保证做到以下两点：
//1. 在额外空间充足的情况下，尽量增大桶的数量
//2. 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
//设置固定数量的空桶。
//把数据放到对应的桶中。
//对每个不为空的桶中数据进行排序。
//拼接不为空的桶中数据，得到结果
//最后，对于桶中元素的排序，选择何种比较排序算法对于性能的影响至关重要。

func BucketSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	bitmap := make([][]int, len(nums))
	_, max := getMaxMin(nums)
	for _, v := range nums {
		k := v * (len(nums) - 1) / max
		bitmap[k] = append(bitmap[k], v)
	}
	tempos := 0
	for i := 0; i < len(nums); i++ {
		bitlen := len(bitmap[i])
		if bitlen > 0 {
			bitmap[i] = InsertSort(bitmap[i])
			copy(nums[tempos:], bitmap[i])
			tempos += bitlen
		}
	}
	return nums
}

//10、基数排序
//基数排序与桶排序、计数排序都用到了桶的概念，但对桶的使用方法上有明显差异：
//基数排序：根据键值的每位数字来分配桶；
//计数排序：每个桶只存储单一键值；
//桶排序：每个桶存储一定范围的数值；
//基数排序按取数方向分为两种：从左取每个数列上的数，为最高位优先（Most Significant Digit first, MSD）；
//从右取每个数列上的数，为最低位优先（Least Significant Digit first, LSD）。下列以LSD为例。
//
//基数排序步骤：
//将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零
//从最低位开始，依次进行一次排序
//从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列

func RadixSort(nums []int) []int {
	key := maxDigists(nums)
	count := make([]int, 10)
	radix := 1
	temp := make([]int, len(nums))
	for i := 0; i < key; i++ {
		for j := 0; j < len(count); j++ {
			count[j] = 0
		}
		for j := 0; j < len(nums); j++ {
			k := nums[j] / radix % 10
			count[k]++
		}
		for j := 1; j < len(count); j++ {
			count[j] = count[j-1] + count[j]
		}
		//逆序，将后面较大的先取出，使其排放在后面。
		for j := len(nums) - 1; j >= 0; j-- {
			k := nums[j] / radix % 10
			temp[count[k]-1] = nums[j]
			count[k]--
		}
		for j := 0; j < len(nums); j++ {
			nums[j] = temp[j]
		}
		radix *= 10
	}
	return nums
}

func maxDigists(nums []int) int {
	res := 0
	count := 1
	for i := 0; i < len(nums); i++ {
		for count < nums[i] {
			count *= 10
			res++
		}
	}
	return res
}
