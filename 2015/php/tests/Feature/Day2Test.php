<?php

use App\Days\Day2;

test('part 1', function () {
    $d2 = Day2::test();

    expect($d2->part1())->toBe(58+43);
});

test('part 2', function () {
    $d2 = Day2::test();

    expect($d2->part2())->toBe(34+14);
});
