<?php

namespace App\Days;

abstract class Day
{
    private bool $isTest = false;
    private ?string $input = null;

    public static function new(): static
    {
        return new static();
    }

    public static function test(string $input = null): static
    {
        $instance = new static();
        $instance->isTest = true;
        $instance->input = $input;

        return $instance;
    }

    public function run(): void
    {
        printf("Part 1: %s\nPart 2: %s", $this->part1(), $this->part2());
    }

    public function getInput(): string
    {
        if (!is_null($this->input)) {
            return $this->input;
        }

        $name = class_basename($this);
        $base = ($this->isTest) ? "tests" : "app/Days";

        return file_get_contents(base_path("$base/inputs/$name"));
    }

    abstract public function part1(): mixed;
    abstract public function part2(): mixed;
}
