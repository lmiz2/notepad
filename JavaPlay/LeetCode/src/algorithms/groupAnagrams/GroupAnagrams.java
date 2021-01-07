package algorithms.groupAnagrams;

import java.util.*;

public class GroupAnagrams {
    public static void main(String[] args) {
        Solution s = new Solution();
        String[] strs = new String[]{"eat","tea","tan","ate","nat","bat","az","bz","za"};
//        String[] strs = new String[]{""};
        System.out.println(s.groupAnagrams(strs));
    }

}
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> rs = new LinkedList<>();
        Map<AnagramString,List<String>> map = new HashMap<>();
        for(String str : strs){
            AnagramString astr = new AnagramString(str);
            if(map.containsKey(astr)){
                map.get(astr).add(str);
            }else{
                List<String> l = new LinkedList<>();
                l.add(str);
                map.put(astr,l);
            }
        }

        for(AnagramString key : map.keySet()){
            rs.add(map.get(key));
        }

        return rs;
    }
}
class AnagramString {
    final static int ALPHABET_START_INDEX_LOWER = 97;
    final static int ALPHABET_END_INDEX_LOWER = 122;
    final static int ALPHABET_COUNT_FOR_LOWER = ALPHABET_END_INDEX_LOWER-ALPHABET_START_INDEX_LOWER+1;
    final static int[] PRIME_NUMBERS = new int[]{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101};
    private String str;
//    private int[] charCounts;
    private long uniqueHash;
    private int hashcode;

    public AnagramString(String str) {
        this.str = str;
//        this.charCounts = new int[ALPHABET_COUNT_FOR_LOWER];
        this.hashcode = 0;
        this.uniqueHash = 1;
        for (char c : this.str.toCharArray()) {
//            charCounts[c-ALPHABET_START_INDEX_LOWER]++;
            this.uniqueHash *= PRIME_NUMBERS[c-ALPHABET_START_INDEX_LOWER];
            this.hashcode += c;
        }
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        AnagramString that = (AnagramString) o;
//        return hashcode == that.hashcode && Arrays.equals(charCounts, that.charCounts);
        return hashcode == that.hashcode && this.uniqueHash == that.uniqueHash;
    }

    @Override
    public int hashCode() {
        return this.hashcode;
    }

}