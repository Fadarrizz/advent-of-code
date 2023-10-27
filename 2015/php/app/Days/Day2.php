<?php

namespace App\Days;

class Day2 extends Day
{
    public function part1(): mixed
    {
        $lines = explode("\n", trim($this->getInput()));

        $result = 0;
        foreach ($lines as $line) {
            [$l, $w, $h] = explode("x", $line);

            $lw = $l * $w;
            $wh = $w * $h;
            $hl = $h * $l;

            $result += (2*$lw) + (2*$wh) + (2*$hl) + min($lw, $wh, $hl);
        }

        return $result;
    }

    public function part2(): mixed
    {
        $lines = explode("\n", trim($this->getInput()));

        $result = 0;
        foreach ($lines as $line) {
            [$l, $w, $h] = explode("x", $line);

            $ribbon = $l + $l + $w + $w;
            $bow = $l * $w * $h;

            $result += $ribbon + $bow;
        }

        return $result;
    }
}
