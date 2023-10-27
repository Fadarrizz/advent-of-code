<?php

namespace App\Days;

class Day4 extends Day
{
    public function part1(): mixed
    {
        $key = trim($this->getInput());
        $num = 0;
        $hash = md5("$key$num");

        while (!str_starts_with($hash, '00000')) {
            $num++;
            $hash = md5("$key$num");
        }

        return $num;
    }

    public function part2(): mixed
    {
        $key = trim($this->getInput());
        $num = 0;
        $hash = md5("$key$num");

        while (!str_starts_with($hash, '000000')) {
            $num++;
            $hash = md5("$key$num");
        }

        return $num;
    }
}
