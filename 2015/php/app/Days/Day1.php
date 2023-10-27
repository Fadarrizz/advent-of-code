<?php

namespace App\Days;

class Day1 extends Day
{
    public function part1(): string
    {
        $floor = 0;
        foreach (str_split($this->getInput()) as $i => $c) {
            if ($c === '(') {
                $floor++;
            } else {
                $floor--;
            }
        }

        return $floor;
    }

    public function part2(): string
    {
        $floor = 0;
        foreach (str_split($this->getInput()) as $i => $c) {
            if ($c === '(') {
                $floor++;
            } else {
                $floor--;
            }

            if ($floor === -1) {
                return $i+1;
            }
        }

        return false;
    }
}
