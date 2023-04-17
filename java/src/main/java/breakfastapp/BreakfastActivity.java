package breakfastapp;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

@ActivityInterface
public interface BreakfastActivity {
    @ActivityMethod
    void getBowl();

    @ActivityMethod
    void putBowlAway();

    @ActivityMethod
    void addCereal();

    @ActivityMethod
    void putCerealBackInBox();

    @ActivityMethod
    void addMilk();
}