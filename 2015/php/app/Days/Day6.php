<?php

namespace App\Days;

use PDO;

class Day6 extends Day
{
    public function part1(): mixed
    {
        $instructions = explode("\n", trim($this->getInput()));

        $lamps = [];
        for ($i = 0; $i < 1000; $i++) {
            $lamps[$i] = array_fill(0, 1000, 0);
        }

        foreach ($instructions as $instruction) {
            $words = explode(" ", $instruction);

            if ($words[0] === 'toggle') {
                $start = explode(',', $words[1]);
                $end = explode(',', $words[3]);

                for ($i = $start[0]; $i <= $end[0]; $i++) {
                    for ($j = $start[1]; $j <= $end[1]; $j++) {
                        $lamps[$i][$j] += 2;
                    }
                }

                continue;
            }

            $start = explode(',', $words[2]);
            $end = explode(',', $words[4]);

            for ($i = $start[0]; $i <= $end[0]; $i++) {
                for ($j = $start[1]; $j <= $end[1]; $j++) {
                    if ($words[1] === 'on') {
                        $lamps[$i][$j] += 1;
                    } else {
                        $lamps[$i][$j] = max(0, $lamps[$i][$j] - 1);
                    }
                }
            }
        }

        $count = 0;
        for ($i = 0; $i < 1000; $i++) {
            for ($j = 0; $j < 1000; $j++) {
                $count += $lamps[$i][$j];
            }
        }

        return $count;
    }

    public function part2(): mixed
    {
        return '';
    }
}
