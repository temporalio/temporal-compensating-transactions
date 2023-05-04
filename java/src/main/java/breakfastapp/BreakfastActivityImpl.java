package breakfastapp;

public class BreakfastActivityImpl implements BreakfastActivity {

    @Override
    public void getBowl() {
        System.out.println("Getting bowl");
    }

    @Override
    public void putBowlAwayIfPresent() {
        System.out.println("Putting bowl away if bowl is out");
    }

    @Override
    public void addCereal() {
        System.out.println("Adding cereal");
    }

    @Override
    public void putCerealBackInBoxIfPresent() {
        System.out.println("Put cereal back in box if there is cereal");
    }

    @Override
    public void addMilk() {
        System.out.println("Adding milk");
    }
}
