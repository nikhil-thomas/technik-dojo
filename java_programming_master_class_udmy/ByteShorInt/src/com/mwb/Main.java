package com.mwb;

public class Main {

    public static void main(String[] args) {
	    int myValue = 10000;
        System.out.println(myValue);

        // int size is 32
        int minInt = -2_147_483_648;
        int maxInt = 2_147_483_647;

        int intResult = maxInt/2;
        System.out.println("intResult: " + intResult);

        // byte size is 8
        byte myByte = -128;
        byte myByte2 = 127;

        byte byteResult = (byte) (myByte2/2);
        System.out.println("byteResult: " + byteResult);

        // short size is 16
        short minVal = -32768;
        short maxVal = 32767;
        short shortResult = (short) (maxVal/2);
        System.out.println("shortResult: " + shortResult);

        // long size is 64
        long val = 100L;
        long minLongVal = -9_223_372_036_854_775_808L;
        long maxLongVal = 9_223_372_036_854_775_807L;

        // long can accept integer hence cast is not necessary
        long longResult = (maxLongVal/2);
        System.out.println("longResult: " + longResult);

        System.out.println("::::::::::");

        byte a = 4;
        short b = 3;
        int c = 3;
        long d = 50000 + 10 * (a + b + c);
        long e = 50000L + 10L * (a + b + c);
        short f = (short) (1000 + 10 * (a + b + c));
        System.out.println("Result: " + d + " : " + e + " : " + f);
    }
}
