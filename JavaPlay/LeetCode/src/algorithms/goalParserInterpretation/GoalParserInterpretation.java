package algorithms.goalParserInterpretation;

class GoalParserInterpretation {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution s = new Solution();
		System.out.println(s.interpret("G()(al)"));
	}

}
class Solution {
    public String interpret(String command) {
    	StringBuilder sb = new StringBuilder();
    	for(int i = 0; i < command.length(); i++) {
    		if(command.charAt(i) == '(') {
    			i++;
    			if(command.charAt(i) == ')') {
    				sb.append('o');
    			}else{
    				i += 2;
    				sb.append("al");
    			}
    		}else if(command.charAt(i) == 'G') {
    			sb.append('G');
    		}
    	}
		return sb.toString();
    }
}