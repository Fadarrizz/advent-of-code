<?php

use App\Days\Day4;

test('part 1', function () {
    expect(Day4::test('abcdef')->part1())->toBe(609043);
});
