# Simple CSV Upload Form for batch processing images with remove-bg

## Install

clone the repository

``` https://github.com/remove-bg/integration.git ```

go to the php/batch-csv-urls folder
 
install dependencies  

``` composer install ```

## Configuration 

Grab a free API key from https://www.remove.bg/api

Open the index.php file and adjust the value at the top of the file.

Make sure your result folder is writable by the webserver.


``` 
// your API key from remove.bg
$apikey = '';
// the folder where the results should be stored, relativ to this file
$results = './results/';

//CSV Settings
//column index for the url, starting at 0
$urlColIndex = 0;
$delimiter = ';';
``` 

## Usage 
go to your browser, upload a csv file and start processing. 

If everything works fine you should see links to the original and the processed images. 


