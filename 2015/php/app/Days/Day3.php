<?php

namespace App\Days;

class Day3 extends Day
{
    public function part1(): mixed
    {
        $chars = str_split(trim($this->getInput()));

        $visited = ['0,0' => true];
        $x = 0; $y = 0;
        foreach ($chars as $c) {
            match ($c) {
                '>' => $x++,
                '<' => $x--,
                '^' => $y--,
                'v' => $y++,
            };

            $visited["$x,$y"] = true;
        }

        return count($visited);
    }

    public function part2(): mixed
    {
        $chars = str_split(trim($this->getInput()));

        $visited = ['0,0' => true];
        $sx = 0; $sy = 0; $rx = 0; $ry = 0;
        foreach ($chars as $i => $c) {
            if ($i % 2 === 0) {
                match ($c) {
                    '>' => $sx++,
                    '<' => $sx--,
                    '^' => $sy--,
                    'v' => $sy++,
                };
                $visited["$sx,$sy"] = true;
            } else {
                match ($c) {
                    '>' => $rx++,
                    '<' => $rx--,
                    '^' => $ry--,
                    'v' => $ry++,
                };
                $visited["$rx,$ry"] = true;
            }
        }

        return count($visited);
    }
}
