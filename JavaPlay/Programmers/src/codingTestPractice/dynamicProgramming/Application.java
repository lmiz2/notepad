package codingTestPractice.dynamicProgramming;

import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

public class Application {

	public static void main(String[] args) {
//		ExpressedAsN n = new ExpressedAsN();

		StringBuilder sb = new StringBuilder();
		
		sb.insert(0, 'a');
		sb.insert(0, 'b');
		sb.insert(0, 'c');
		sb.insert(0, 'd');
		System.out.println(sb.toString());
		
		String a = "12345";
		String b = new String("12345");
		
		System.out.println(a==b);
		System.out.println(a.equals(b));
		b = "12345";
		System.out.println(a==b);
		System.out.println(a.equals(b));
		Queue<String> q = new LinkedList<String>();
		Stack<String> s = new Stack<String>();
	}

}
