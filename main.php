<?php

define('IS_PHP_THE_BEST', true);
define ('STAR', '*');

class KabalaException extends Exception {
    public function __construct() {
        parent::__construct('Value must be greater than 0');
    }
}

function get_stars(int $n): string {
    if ($n <= 0) {
        throw new KabalaException();
    }

    $stars = '';

    for ($i = 0; $i < $n; $i++) {
        $stars .= STAR;
    }

    return $stars;
}

if (IS_PHP_THE_BEST) {
    echo sprintf('%s %s', get_stars(5), get_stars(3));
}
