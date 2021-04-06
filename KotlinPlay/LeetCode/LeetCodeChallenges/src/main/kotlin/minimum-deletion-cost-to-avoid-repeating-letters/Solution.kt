package `minimum-deletion-cost-to-avoid-repeating-letters`

class Solution {
    fun minCost(s: String, cost: IntArray): Int {
        val list = ArrayList<Pair<Char, Int>>(cost.size)
        for (i in cost.indices) {
            list.add(Pair(s[i],cost[i]))
        }
        return conquer(divide(list),cost)
    }

    fun divide(charAndCosts: List<Pair<Char, Int>>): List<Pair<Int, Int>> = run {
        val rs = ArrayList<Pair<Int, Int>>()
        var left = 0
        var right = 0
        var before = '-'
        for( (i,d) in charAndCosts.withIndex() ){
            if(before == '-'){ // 처음
                before = d.first
            }else if(before != d.first) { // 다른 캐릭터 나오면
                if(left != right) rs.add(Pair(left,right))

                before = d.first
                left = i
                right = i
            }else{ // 같은 캐릭터
                right++
            }
        }
        if(left != right) rs.add(Pair(left,right))
        return rs
    }

    fun conquer(arr: List<Pair<Int, Int>>, costs: IntArray) : Int = run {
        var total = 0
        for(d in arr){
            var sumCost = 0
            var maxCost = 0
            for(i in d.first..d.second){
                sumCost += costs[i]
                maxCost = if(maxCost <= 0){
                    costs[i]
                }else{
                    maxCost.coerceAtLeast(costs[i])
                }
            }
            total += (sumCost - maxCost)
        }
        return total
    }


}