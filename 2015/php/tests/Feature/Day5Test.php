<?php

namespace Tests\Feature;

use App\Days\Day5;

test('part 1', function () {
    expect(Day5::test('ugknbfddgicrmopn')->part1())->toBe(1);
    expect(Day5::test('aaa')->part1())->toBe(1);
    expect(Day5::test('jchzalrnumimnmhp')->part1())->toBe(0);
    expect(Day5::test('haegwjzuvuyypxyu')->part1())->toBe(0);
    expect(Day5::test('dvszwmarrgswjxmb')->part1())->toBe(0);
});

test('part 2', function () {
    expect(Day5::test('qjhvhtzxzqqjkmpb')->part2())->toBe(1);
    expect(Day5::test('xxyxx')->part2())->toBe(1);
    expect(Day5::test('uurcxstgmygtbstg')->part2())->toBe(0);
    expect(Day5::test('ieodomkazucvgmuy')->part2())->toBe(0);
});
