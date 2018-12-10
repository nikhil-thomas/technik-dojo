public class Greeting {
    public static void main(String[] args) {
        greet("Alex");

        Developer developer = new Developer("Alex", "Dan", "abc123");

        developer.printFullName();
    }

    public static void greet(String name) {
        System.out.println("Hello, " + name);
    }
}
