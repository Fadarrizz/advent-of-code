<?php

namespace App\Days;

use App\Days\Day7\Instruction;
use App\Days\Day7\u16;

class Day7 extends Day
{
    public array $circuit;

    public function part1(): mixed
    {
        $this->circuit = $this->parseCircuit();

        return $this->emulate('a');
    }

    /**
     * @return array<string, Instruction>
     */
    public function parseCircuit(): array
    {
        $circuit = [];
        foreach (explode("\n", trim($this->getInput())) as $line) {
            $instruction = Instruction::parse($line);
            $circuit[$instruction->wire()] = $instruction;
        }

        return $circuit;
    }

    public function emulate(mixed $wire): ?u16
    {
        if (is_null($wire)) {
            return null;
        }

        if ($wire instanceof u16) {
            return $wire;
        }

        $instruction = $this->circuit[$wire];

        if ($instruction instanceof u16) {
            return $instruction;
        }

        /* @var Instruction $instruction */
        if (is_null($instruction->command())) {
            $this->circuit[$wire] = $this->emulate($instruction->firstArg());
        } else {
            $this->circuit[$wire] = $this->applyCommand(
                $instruction->command(),
                $this->emulate($instruction->firstArg()),
                $this->emulate($instruction->secondArg())
            );
        }

        return $this->circuit[$wire];
    }

    private function applyCommand(string $command, u16 $a, ?u16 $b): u16
    {
        return match ($command) {
            "AND" => $a->and($b),
            "OR" => $a->or($b),
            "LSHIFT" => $a->lshift($b),
            "RSHIFT" => $a->rshift($b),
            "NOT" => $a->not(),
        };
    }

    public function part2(): mixed
    {
        $this->circuit = $this->parseCircuit();
        $this->circuit['b'] = u16::new(46065);

        return $this->emulate('a');
    }
}
