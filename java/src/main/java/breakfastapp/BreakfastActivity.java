package breakfastapp;

import io.temporal.activity.ActivityInterface;
import io.temporal.activity.ActivityMethod;

@ActivityInterface
public interface BreakfastActivity {
    @ActivityMethod
    void getBowl();

    @ActivityMethod
    void putBowlAwayIfPresent();

    @ActivityMethod
    void addCereal();

    @ActivityMethod
    void putCerealBackInBoxIfPresent();

    @ActivityMethod
    void addMilk();
}