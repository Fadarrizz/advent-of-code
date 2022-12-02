<?php

namespace Tests\Feature;

use App\Days\CalorieHunting;
use Tests\TestCase;

class CalorieHuntingTest extends TestCase
{
    /**
     * A basic feature test example.
     *
     * @return void
     */
    public function test_example()
    {
        $data = fopen(base_path('tests/Feature/Data/calorieHunting.txt'), 'r');

        $result = (new CalorieHunting($data))->part2($data);

        dd($result);
    }
}
