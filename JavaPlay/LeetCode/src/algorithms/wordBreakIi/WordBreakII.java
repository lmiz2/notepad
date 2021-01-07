package algorithms.wordBreakIi;

import java.util.*;
import java.util.concurrent.atomic.AtomicReference;

public class WordBreakII {
    public static void main(String[] args) {
        String s = "pineapplepenapple";
//        String s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
        List<String> wordDict = List.of("apple", "pen", "applepen", "pine", "pineapple");
//        List<String> wordDict = List.of("a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa");
        System.out.println(new Solution().wordBreak(s,wordDict));
//        System.out.println(s.substring("apple".length()-1,s.length()));
    }
}
// 1. Only DFS : 10~11ms
//class Solution {
//    public List<String> wordBreak(String s, List<String> wordDict) {
//        boolean[] check = new boolean[26];
//        wordDict.forEach(s1 -> {
//            for (int i = 0 ; i < s1.length() ; i++){
//                check[s1.charAt(i)-97] = true;
//            }
//        });
//        for(int i = 0; i < s.length(); i++){
//            if(!check[s.charAt(i)-97]){
//                return List.of();
//            }
//        }
//        return recursive(s,wordDict,new LinkedList<>(), new StringBuilder());
//    }
//    public List<String> recursive(String s, List<String> wordDict, List<String> list, StringBuilder sb){
//        if(s.length() == 0){
//            list.add(sb.toString());
//        }else{
//            wordDict.stream().filter(word -> s.indexOf(word) == 0).forEach(word ->{
//                System.out.println(sb.toString());
//                StringBuilder tsb = new StringBuilder(sb);
//                if(tsb.length() != 0){
//                    tsb.append(" ");
//                }
//                tsb.append(word);
//                recursive(s.substring(word.length(),s.length()),wordDict,list,tsb);
//            });
//        }
//        return list;
//    }
//}

// 2. DFS With Caching : 19ms ... why??
class Solution {
    public List<String> wordBreak(String s, List<String> wordDict) {
        boolean[] check = new boolean[26];
        wordDict.forEach(s1 -> {
            for (int i = 0 ; i < s1.length() ; i++){
                check[s1.charAt(i)-97] = true;
            }
        });
        for(int i = 0; i < s.length(); i++){
            if(!check[s.charAt(i)-97]){
                return List.of();
            }
        }
        return recursive(s,wordDict, new HashMap<>());
    }
    public List<String> recursive(String s, List<String> wordDict, Map<String,List<String>> mem){
        List<String> rs = new LinkedList<>();
        if(mem.containsKey(s)){
            return mem.get(s);
        }else{
            wordDict.stream().filter(word -> s.indexOf(word) == 0).forEach(word ->{
                if(word.equals(s)){
                    rs.add(word);
                }else{
                    List<String> list = recursive(s.substring(word.length()),wordDict,mem);
                    list.forEach(s1 -> rs.add(word+" "+s1));
                }
            });

        }
        mem.put(s,rs);
        return rs;
    }
}

