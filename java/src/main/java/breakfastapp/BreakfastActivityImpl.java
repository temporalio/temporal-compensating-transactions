package breakfastapp;

public class BreakfastActivityImpl implements BreakfastActivity {

    @Override
    public void getBowl() {
        System.out.println("Getting bowl");
    }

    @Override
    public void putBowlAway() {
        System.out.println("Putting bowl away");
    }

    @Override
    public void addCereal() {
        System.out.println("Adding cereal");
    }

    @Override
    public void putCerealBackInBox() {
        System.out.println("Put cereal back in box");
    }

    @Override
    public void addMilk() {
        System.out.println("Adding milk");
    }
}
