package code75

/*
 Running Sum of 1d Array
配列numsが与えられたとき配列のランニングサムを
runningSum[i] = sum(nums[0]...nums[i])と定義します．
numsの総和を返します。

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
*/

func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

/*
logic
このコードは、配列 nums の各要素を前から順に走査し、その要素の値に前の要素の値を加える処理を行っています。
まず、for文で配列 nums の要素を走査しています。for文では、変数 i を 1 から配列の要素数より小さい値まで、1ずつ増やしながら繰り返し処理を行います。
次に、nums[i] += nums[i-1] という文があります。これは、配列 nums の i 番目の要素に、i-1 番目の要素を加える処理を行っています。
例えば、配列 nums が [1, 2, 3, 4] の場合、最初のfor文では、i は 1 から 3 まで、3回繰り返し処理されます。
そして、各繰り返しで nums[i] += nums[i-1] が実行されるため、次のように配列 nums の各要素が更新されます。
1.最初の繰り返し: nums[1] = nums[1] + nums[0] = 2 + 1 = 3
2.2番目の繰り返し: nums[2] = nums[2] + nums[1] = 3 + 3 = 6
3.3番目の繰り返し: nums[3] = nums[3] + nums[2] = 4 + 6 = 10
最終的に、配列 nums は [1, 3, 6, 10] になります。
最後に、return nums とあるので、配列 nums を関数の戻り値として返すことになります。
*/
