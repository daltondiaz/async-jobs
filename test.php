<?php

$arg = $argv[1];
$sec = random_int(1,5);
sleep($sec);
echo "item ".$arg." test ".$sec. " s";
