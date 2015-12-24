<?php
$servername = "localhost";
$username = "root";
$password = "kamal";
$dbname = "kaf";

if(!(mysqli_connect($servername,$username,$password))) {
    die("Connection failed: " . $conn->connect_error);
} 

$sql = "SELECT id, festname FROM fests WHERE month = 'October'";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
    // output data of each row
    while($row = $result->fetch_assoc()) {
        echo "id: " . $row["id"]. " - Name: " . $row["firstname"] "<br>";
    }
} else {
    echo "0 results";
}
$conn->close();
?>