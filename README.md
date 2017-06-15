# Create-Suspend-Terminate-Go-Scripts
Create, Suspend and Terminate Go Scripts

How to Call the Go Program from the PHP Frontend

php -r '$myObj->id = 2;$myObj->client_id = 15880;$myObj->device_id = 11064;$myObj->domainname = "dosattack.net--";$myObj->ipaddress = "151.101.192.253";$myObj->domainport = 443;$myObj->sitename = "dosattacknet";$myObj->status = "Created";$myObj->updated_at = "4/4/2017  11:05:31 AM";$myObj->enableSsl = 0;$myObj->enableGzip = 0;$myObj->sslKey = "dosattacknet.key";$myObj->sslCrt = "dosattacknet.crt";$myObj->wcRedirect = 0;$myObj->action = "create";$myJSON = json_encode($myObj);echo "\n";$output = shell_exec("./cdn_create " .escapeshellarg( $myJSON));echo $output;'

This is a sample call to the Go Program.
Please note the 'action' parameter (action = "create"). This defines the flow of the program. The acceptable values for this parameter are "create", "suspend", "unsuspend", "terminate"
The suspended files are moved to 'suspend' directory appended with the .suspend extension.
The terminate files are moved to 'terminate' directory appended with the .terminate extension - these are maintained for backtracking purposes - if a need arises in the future to investigate the subscriptions, etc. then, these files can be used for reporting purposes.
The label (cdn_create) indicates the file name/Go program to call. This name will be finalized and duly communicated to the frontend team.

If additional parameters will be added to this JSON or there are any changes to the input data types, they must be communicated explicitly.

It is advised that a full upload path be maintained for the  SSL certificate files as an externally configurable parameter at the frontend and be passed on to the backend via the JSON. The same for the DNS Zone file. The reason being that any application-wide configuration ought to be maintained and controlled centrally.

The Go Program will return either a "success" or an "error" upon completion.
After completion, the frontend must prompt the customer : "Please create CNAME DNS record at your domain registrar and point the CNAME to deviceid.domain.extension.pokecdn.net.  Please allow 24-48 hours for the DNS to update."

Assumptions:

•	The parameter values are duly validated by the frontend for each input field - ex. ipaddress for valid IP values, domainname for valid domain names, etc.
•	Any changes to the input format are duly communicated by the frontend team to the backend
•	After program execution, the listener will automatically propagate/rsync the newly generated files to all the edge nodes
