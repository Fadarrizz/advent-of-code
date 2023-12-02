<?php

namespace Tests\Feature;

use App\Days\Day7;
use App\Days\Day8;

test('part 1', function () {
    $input = <<<'EOD'
    ""
    "abc"
    "aaa\"aaa"
    "\x27"
    EOD;

    expect(Day8::test($input)->part1())->toBe(12);
});

test('part 2', function () {
    expect(Day7::test('qjhvhtzxzqqjkmpb')->part2())->toBe(1);
    expect(Day7::test('xxyxx')->part2())->toBe(1);
    expect(Day7::test('uurcxstgmygtbstg')->part2())->toBe(0);
    expect(Day7::test('ieodomkazucvgmuy')->part2())->toBe(0);
});
