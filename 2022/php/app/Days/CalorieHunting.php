<?php

namespace App\Days;

class CalorieHunting
{
    private array $elfs;

    public function __construct($data)
    {
        $this->elfs = $this->sumCaloriesPerElf($data);
    }

    public function part1(): int
    {
        return max($this->elfs);
    }

    public function part2(): int
    {
        $_elfs = $this->elfs;

        rsort($_elfs);

        return $_elfs[0] + $_elfs[1] + $_elfs[2];
    }

    private function sumCaloriesPerElf($data): array
    {
        $elfs = [];
        $calories = [];
        while (! feof($data)) {
            $line = fgets($data);

            if ($line === "\n") {
                $elfs[] = array_reduce($calories, fn ($carry, $calorie) => $carry += $calorie, 0);
                $calories = [];
                continue;
            }

            $calories[] = $line;
        }

        fclose($data);

        return $elfs;
    }
}
