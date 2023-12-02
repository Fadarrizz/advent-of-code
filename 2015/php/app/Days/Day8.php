<?php

namespace App\Days;

class Day8 extends Day
{
    public function part1(): mixed
    {
        $count = 0;
        foreach (explode("\n", $this->getInput()) as $line) {
            $trimmed = trim($line);

            $evaluatedString = stripcslashes(trim($trimmed, '"'));

            dump(strlen($trimmed) . ' - ' . strlen($evaluatedString));
            $count += strlen($trimmed) - strlen($evaluatedString);
        }

        return $count;
    }

    public function part2(): mixed
    {
        return "";
    }
}
