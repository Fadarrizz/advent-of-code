<?php

namespace Tests\Feature;

use App\Days\Day6;

test('part 1', function () {
    expect(Day6::test('turn on 0,0 through 999,999')->part1())->toBe(1000*1000);
    expect(Day6::test("turn on 0,0 through 999,999\nturn off 499,499 through 500,500")->part1())->toBe(1000*1000-4);
    expect(Day6::test("turn on 0,0 through 0,1\ntoggle 0,0 through 0,1")->part1())->toBe(0);
});

test('part 2', function () {
    expect(Day6::test('qjhvhtzxzqqjkmpb')->part2())->toBe(1);
    expect(Day6::test('xxyxx')->part2())->toBe(1);
    expect(Day6::test('uurcxstgmygtbstg')->part2())->toBe(0);
    expect(Day6::test('ieodomkazucvgmuy')->part2())->toBe(0);
});
