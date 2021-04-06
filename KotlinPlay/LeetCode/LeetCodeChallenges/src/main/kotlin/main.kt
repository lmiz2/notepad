import `minimum-deletion-cost-to-avoid-repeating-letters`.Solution

fun main(args: Array<String>) {
    var s : Solution = Solution()
//    var arr  = intArrayOf(1,2,3,4,1,2)
//    var str = "aabaaa"
    var arr  = intArrayOf(1,2,3,4,1)
    var str = "aabaa"
    println(s.minCost(str, arr))

}