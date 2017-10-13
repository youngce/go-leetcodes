package solutions

func max(x int ,y int) int {
	if x>y{
		return x
	}
	return y
}
func maxSubArray(nums []int) int {
	//var maxCurr, maxGlobal int
	maxCurr,maxGlobal:=nums[0],nums[0]

	for i:=1;i<len(nums) ;i++  {
		maxCurr=max(nums[i],maxCurr+nums[i])
		if maxCurr>maxGlobal{
			maxGlobal=maxCurr
		}
	}
	return maxGlobal



}
