<?php

namespace App\Days\Day7;

class u16
{
    const BITS = 0xFFFF;

    public int $value;

    public function __construct(int $num)
    {
        $this->value = $num & self::BITS;
    }

    public static function new(int $num = 0): u16
    {
        return new self($num);
    }

    public function and(u16 $other): u16
    {
        return self::new($this->value & $other->value);
    }

    public function or(u16 $other): u16
    {
        return self::new($this->value | $other->value);
    }

    public function lshift(u16 $other): u16
    {
        return self::new($this->value << $other->value);
    }

    public function rshift(u16 $other): u16
    {
        return self::new($this->value >> $other->value);
    }

    public function not(): u16
    {
        return self::new(~ $this->value);
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}
