<?php
/**
 * Simple CSV Upload Form for batch processing images with remove-bg
 */

require __DIR__.'/vendor/autoload.php';

use Mtownsend\RemoveBg\RemoveBg;

// your API key from remove.bg
$apikey = '';
// the folder where the results should be stored, relativ to this file
$results = './results/';

//CSV Settings
//column index for the url in the csv file, starting at 0
$urlColIndex = 0;
$delimiter = ';';

$error = '';
$resultData = array();
//process the uploaded file
if (isset($_POST['upload'])) {
    if (isset($_FILES['file']) && $_FILES['file']['error'] === 0 ) {
        $removebg = new RemoveBg($apikey);
        $csvFilename =  $_FILES['file']['tmp_name'];
        $f = fopen($csvFilename, 'r');
        while ( ($data = fgetcsv($f, 4096 , $delimiter) ) !== false) {
            $url = $data[$urlColIndex];
            $newFileName = pathinfo( parse_url($url, PHP_URL_PATH), PATHINFO_FILENAME) . ".png";
            if ($newFileName == '') {
                $newFileName = uniqid('',false).'.png';
            }
            try {
                $result = $removebg->url($url)->save($results.$newFileName);
                if ($result) {
                    $resultData[$url] = $newFileName;
                }
            } catch ( Exception $e) {
                $error = $e->getMessage();
            }
        }
    } else {
        $error = 'No file specified or upload error.';
    }
}

?>
<html>
    <head>
        <title>Simple CSV processing | remove.bg</title>
        <style>
            label{display:block}
            .error {color:red}
            .content {margin-left: auto; margin-right: auto; width:40em;text-align: center}
        </style>
    </head>
    <body>
        <div class="content">
            <h1>Simple CSV processing</h1>
            <form method="POST" enctype="multipart/form-data">
                <p>
                    <label for="file">Please choose a .csv file:</label>
                    <input id="file" type="file" name="file">
                </p>
                <p>
                    <input  name="upload" type="submit" value="Upload and process">
                </p>
            </form>

            <div class="result">
                <?php
                    if ($error != '') {
                        echo '<div class="error">'.$error.'</div>';
                    }
                    if (count($resultData)) {
                        echo '<h2>Results</h2>';
                        foreach ($resultData as $org => $file) {
                            echo '<p><a href="'.$org.'" target="_blank">Original</a> &raquo; <a href="'.$results.$file.'" target="_blank">Result</a></p>';
                        }
                    }
                ?>
            </div>
        </div>
    </body>
</html>
