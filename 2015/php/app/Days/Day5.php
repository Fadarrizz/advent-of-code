<?php

namespace App\Days;

class Day5 extends Day
{
    public function part1(): mixed
    {
        $strings = explode("\n", trim($this->getInput()));

        $niceStrings = array_filter($strings, fn (string $string) => $this->isNiceString($string));

        return count($niceStrings);
    }

    public function part2(): mixed
    {
        $strings = explode("\n", trim($this->getInput()));

        $niceStrings = array_filter($strings, fn (string $string) => $this->isNiceString2($string));

        return count($niceStrings);
    }

    private function isNiceString(string $string): bool
    {
        if (strlen($string) - strlen(str_replace(['a','e','i','o','u'], '', $string)) < 3) {
            return false;
        }

        if (!$this->twoInARow($string, 1)) {
            return false;
        }

        if (str_replace(['ab','cd','pq','xy'], '', $string) !== $string) {
            return false;
        }

        return true;
    }

    private function isNiceString2(string $string): bool
    {
        if (!$this->twoPair($string)) {
            return false;
        }

        if (!$this->twoInARow($string, 2)) {
            return false;
        }

        return true;
    }

    private function twoInARow(string $string, int $offset): bool
    {
        $len = strlen($string);
        for ($i = 0; $i < $len; $i++) {
            if (isset($string[$i + $offset]) && $string[$i] === $string[$i + $offset]) {
                return true;
            };
        }

        return false;
    }

    private function twoPair(string $string): bool
    {
        $len = strlen($string);
        for ($i = 2; $i < $len; $i++) {
            $a = $string[$i - 2] . $string[$i - 1];

            for ($j = $i + 1; $j < $len; $j++) {
                $b = $string[$j - 1] . $string[$j];

                if ($a === $b) {
                    return true;
                }
            }
        }

        return false;
    }
}
