# nibbs-transfer-status-checker

This project is simple API that checks if the default NIP service provider is active. 
It's feeding from the NIP live success / failure rate service and returns an indicator if the provider has a failure rate of less than a specified percentage.
The maximum allowed rate is currently set to 1.7%, this rate can be changed in the .env file located in the root directory.

It's a Golang project, it was developed using Echo (https://echo.labstack.com), a high performance, extensible, minimalist Go web framework.

<b>Installation</b>
<p><i>Requirements:</i></p>
Install Docker

<p>Step 1: Clone the project, natigate to the root directory on your terminal and run <b><i>docker build -t nibbs-transfer-status-checker .</i></b> to build the docker image.</p>
<p>Step 2: Run <b><i>docker images</i></b> to see the list of all docker images on the machine, the image of this project should have been created.</p>
<p>Step 3. Run <b><i>docker run -it -p 8090:8090 nibbs-transfer-status-checker</i></b> to run the generated image in a container. The image runs on port 8090 of the host machine. </p>

<p>Lunch your browser or any REST client application and point to <b><i>http://localhost:8090/api/v1/transfer-status</i></b> to check real-time result. </p>
