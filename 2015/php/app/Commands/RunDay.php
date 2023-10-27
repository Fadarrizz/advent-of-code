<?php

namespace App\Commands;

use Illuminate\Console\Scheduling\Schedule;
use LaravelZero\Framework\Commands\Command;

class RunDay extends Command
{
    /**
     * The signature of the command.
     *
     * @var string
     */
    protected $signature = 'r {day}';

    /**
     * The description of the command.
     *
     * @var string
     */
    protected $description = 'Run a day';

    /**
     * Execute the console command.
     */
    public function handle(): void
    {
        $dayNr = $this->argument('day');
        $class = "App\Days\Day$dayNr";
        (new $class)->run();
    }
}
