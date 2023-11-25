<h1>Download the Firebase Configuration File</h1>
Go to your Firebase console.
Select the desired project.
Navigate to "Project settings" > "Service accounts" > "Firebase Admin SDK".
Click on "Generate new private key file" and save the json file.

<h1>Download the Executable</h1>
Go to the Releases section of the GitHub repository and download the latest version of the "user" executable for Linux.

Set the Environment Variables

Set the necessary environment variables for the project. Open a terminal and enter the following commands:

```
export SECRETS_FILE_PATH="/path/to/your/firebase/secrets.json"
export PROJECT_ID="your_project_id"
```

Replace /path/to/your/firebase/secrets.json with the actual path of your Firebase secrets file and your_project_id with your Firebase project ID.

<h1>Application Execution</h1>
Grant Execution Permissions

From the directory where you downloaded the "user" executable, make it executable with the command:

```
chmod +x user
```

<h1>Run the Executable</h1>
Start the application by executing:

```
./user
```
<h1>Verify</h1>
Once the application is started, it should be accessible via a browser or any HTTP client at the address <a href="http://localhost:50001/swagger/index.html#">http://localhost:50001/swagger/index.html#</a>.
