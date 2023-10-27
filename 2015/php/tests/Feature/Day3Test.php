<?php

use App\Days\Day3;

test('part 2', function () {
    expect(Day3::test('^v')->part2())->toBe(3);
    expect(Day3::test('^>v<')->part2())->toBe(3);
    expect(Day3::test('^v^v^v^v^v')->part2())->toBe(11);
});
