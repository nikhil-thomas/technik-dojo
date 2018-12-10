import java.util.Objects;

public class Developer {

    private String firstName;
    private String lastName;
    private String github;

    public Developer(String firstName, String lastName, String github) {
        this.firstName = firstName;
        this.lastName = lastName;
        this.github = github;
    }

    public String getFirstName() {
        return firstName;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getGithub() {
        return github;
    }

    public void setGithub(String github) {
        this.github = github;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Developer developer = (Developer) o;
        return Objects.equals(firstName, developer.firstName) &&
                Objects.equals(lastName, developer.lastName) &&
                Objects.equals(github, developer.github);
    }

    @Override
    public int hashCode() {
        return Objects.hash(firstName, lastName, github);
    }

    @Override
    public String toString() {
        return "Developer{" +
                "firstName='" + firstName + '\'' +
                ", lastName='" + lastName + '\'' +
                ", github='" + github + '\'' +
                '}';
    }

    public void printFullName() {
        System.out.println(firstName + " " + lastName);
    }
}
