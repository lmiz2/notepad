package algorithms.ipo;

import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;

public class Ipo {
    public static void main(String[] args) {
        Solution s = new Solution();
//        System.out.println(s.findMaximizedCapital(2,2,new int[]{1,2,3},new int[]{0,1,2}));
        System.out.println(s.findMaximizedCapital(2,0,new int[]{1,2,3},new int[]{0,1,1}));
    }
}
class Solution {
    public int findMaximizedCapital(int k, int W, int[] Profits, int[] Capital) {
        AtomicInteger sum = new AtomicInteger(W);
        List<Project> list = new LinkedList<>();
        for(int i = 0 ; i < Profits.length ; i++){
            list.add(new Project(Capital[i],Profits[i]));
        }
        Collections.sort(list);
        while(k-- > 0){
            Optional<Project> op = list.stream().filter(project -> sum.get() >= project.getCapital()).findFirst();
            op.ifPresentOrElse(
                    project -> {
                        sum.addAndGet(project.getProfit());
                        list.remove(project);
                        },
                    () -> {});
        }
        return sum.get();
    }
}
class Project implements Comparable<Project>{
    private int capital;
    private int profit;
//    private int efficiency;

    public Project(int capital, int profit) {
        this.capital = capital;
        this.profit = profit;
//        this.efficiency = profit - capital;
    }

    public int getCapital() {
        return capital;
    }

    public int getProfit() {
        return profit;
    }

//    public int getEfficiency() {
//        return efficiency;
//    }

    @Override
    public String toString() {
        return "Project{" +
                "capital=" + capital +
                ", profit=" + profit +
//                ", efficiency=" + efficiency +
                '}';
    }

    @Override
    public int compareTo(Project o) {
//        if(this.efficiency != o.efficiency){
//            return o.efficiency-this.efficiency;
//        }
        if(this.profit != o.profit){
            return o.profit-this.profit;
        }
        return this.capital-o.capital;
    }
}