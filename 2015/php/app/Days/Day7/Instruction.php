<?php

namespace App\Days\Day7;

class Instruction
{
    /**
     * @param array<int,string|u16> $arguments
     */
    public function __construct(
        private ?string $command,
        private array $arguments,
        private string $wire,
    ) {}

    public static function parse(string $str): Instruction
    {
        preg_match("/[A-Z]+/", $str, $command);
        preg_match_all("/[a-z0-9]+/", $str, $args);
        $wire = array_pop($args[0]);

        $args = array_map(function ($arg) {
            return is_numeric($arg) ? u16::new((int) $arg) : $arg;
        }, $args[0]);

        return new self($command[0] ?? null, $args, $wire);
    }

    public function command(): ?string
    {
        return $this->command;
    }

    public function firstArg(): string|u16
    {
        return $this->arguments[0];
    }

    public function secondArg(): string|u16|null
    {
        return $this->arguments[1] ?? null;
    }

    public function wire(): string
    {
        return $this->wire;
    }
}
