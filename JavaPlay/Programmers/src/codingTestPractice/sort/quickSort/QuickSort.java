package codingTestPractice.sort.quickSort;

import java.util.Arrays;

public class QuickSort {
	
	public static void main(String[] args) {
		int[] arr = new int[] {1,3,7,4,4,8,4,6,5,0,2,6,23};
		quickSort(arr, 0, arr.length-1);
		System.out.println(Arrays.toString(arr));
	}

	
	public static void quickSort(int[] arr, int start, int end) {
		int rightPartition = divide(arr,start,end);
		if(start < rightPartition-1) {
			quickSort(arr, start, rightPartition-1);
		}
		if(rightPartition < end) {
			quickSort(arr, rightPartition, end);
		}
	}
	public static int divide(int[] arr, int start, int end) {
		int pivot = arr[(start+end)/2];
		while(start <= end) {
			while(arr[start] < pivot) start++;
			while(arr[end] > pivot) end--;
			if(start <= end) {
				swap(arr,start,end);
				start++;
				end--;
			}
		}
		return start;
	}
	
	public static void swap(int[] arr, int a, int b) {
		int tmp = arr[a];
		arr[a] = arr[b];
		arr[b] = tmp;
	}
}
