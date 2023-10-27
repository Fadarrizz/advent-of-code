<?php

namespace Tests\Feature;

use App\Days\Day7;

test('part 1', function () {
    $input = <<<EOD
123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i
EOD;
    $d = Day7::test($input);
    $circuit = $d->parseCircuit();

    expect($d->emulate($circuit, 'x')->value)->toBe(123);
    expect($d->emulate($circuit, 'y')->value)->toBe(456);
    expect($d->emulate($circuit, 'd')->value)->toBe(72);
    expect($d->emulate($circuit, 'e')->value)->toBe(507);
    expect($d->emulate($circuit, 'f')->value)->toBe(492);
    expect($d->emulate($circuit, 'g')->value)->toBe(114);
    expect($d->emulate($circuit, 'h')->value)->toBe(65412);
    expect($d->emulate($circuit, 'i')->value)->toBe(65079);
});

test('real', function () {
    Day7::new()->part1();
});

test('part 2', function () {
    expect(Day7::test('qjhvhtzxzqqjkmpb')->part2())->toBe(1);
    expect(Day7::test('xxyxx')->part2())->toBe(1);
    expect(Day7::test('uurcxstgmygtbstg')->part2())->toBe(0);
    expect(Day7::test('ieodomkazucvgmuy')->part2())->toBe(0);
});
