package com.mwb;

public class Main {

    public static void main(String[] args) {
	// write your code here

        // size of int 32 bytes
        int a = 5/3;

        // size of float 32 bytes
        float b = 5f / 3f;

        // size of double 64 bytes
        double c = 5d / 3d;

        System.out.println("int: " + a);
        System.out.println("float: " + b);
        System.out.println("double: " + c);

        double p = 2d;
        double kg = p * 0.45359237d;

        System.out.println("pound: " + p + ", kg: " + kg);

        // language level should be 8 to use '_'
        c = 3_000_000.141_592_37d;

        System.out.println("c: " + c);
    }
}
