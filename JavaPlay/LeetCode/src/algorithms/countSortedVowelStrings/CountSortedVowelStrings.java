package algorithms.countSortedVowelStrings;

public class CountSortedVowelStrings {
    public static void main(String[] args) {
        System.out.println(new Solution().countVowelStrings(0));
    }
}
class Solution {
    final static int CHAR_ASCII_START_INDEX = 97;
    final static char[] VOWELS = new char[]{'a','e','i','o','u'};
    public int countVowelStrings(int n) {
        if(n==1) {
            return oneVowelStrings();
        }
        int[][] mem = new int[26][51];
        int t = 0;
        for (char vowel : VOWELS) {
            t += recursive(vowel,n-1,mem);
        }
        return t;
    }
    public int recursive(char firstChar, int leftedN, int[][] mem){
        if(mem[firstChar-CHAR_ASCII_START_INDEX][leftedN] > 0){
            return mem[firstChar-CHAR_ASCII_START_INDEX][leftedN];
        }
        if(leftedN == 1){
            mem[firstChar-CHAR_ASCII_START_INDEX][leftedN] = twoVowelStrings(firstChar,mem);
            return mem[firstChar-CHAR_ASCII_START_INDEX][leftedN];
        }else{
            int t = 0;
            for (char vowel : VOWELS) {
                if(firstChar <= vowel){
                    t += recursive(vowel,leftedN-1,mem);
                }
            }
            mem[firstChar-CHAR_ASCII_START_INDEX][leftedN] = t;
            return t;
        }
    }
    public int twoVowelStrings(char firstChar, int[][] mem){
        if(mem[firstChar-CHAR_ASCII_START_INDEX][1] > 0){
            return mem[firstChar-CHAR_ASCII_START_INDEX][1];
        }else{
            switch (firstChar){
                case 'a':
                    mem[firstChar-CHAR_ASCII_START_INDEX][1] = 5;
                    return 5;
                case 'e':
                    mem[firstChar-CHAR_ASCII_START_INDEX][1] = 4;
                    return 4;
                case 'i':
                    mem[firstChar-CHAR_ASCII_START_INDEX][1] = 3;
                    return 3;
                case 'o':
                    mem[firstChar-CHAR_ASCII_START_INDEX][1] = 2;
                    return 2;
                case 'u':
                    mem[firstChar-CHAR_ASCII_START_INDEX][1] = 1;
                    return 1;
            }
        }
        return -1;
    }
    public int oneVowelStrings(){
        return 5;
    }
}