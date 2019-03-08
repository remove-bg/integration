<?php
require __DIR__ . '/vendor/autoload.php';

use Mtownsend\RemoveBg\RemoveBg;

$apikey = "YOUR_API_KEY";
$removebg = new RemoveBg($apikey);

$originals = "./originals/";
$results = "./results/";

if($handle = opendir($originals)) {
    while(false !== ($file = readdir($handle))) {
        if ('.' === $file) continue;
        if ('..' === $file) continue;

        $original = $originals . $file;
        $result = $results . pathinfo($file, PATHINFO_FILENAME) . ".png";

        echo "Processing " . $original . " ... ";

        $removebg->file($original)->save($result);

        echo "saved to " . $result . ".\n";
    }
    closedir($handle);
}
