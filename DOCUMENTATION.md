# Endpoint Documentation

| Verb          | Route                                   | Quick Description                                                             |
| ------------- |:---------------------------------------:| --------------------------------:                                             |
| BASE URL      | quicky.dev/api                          |                                                                               |
| GET           | /generic                                | setup script sent as response.                                                |
| GET           | /availableItems                         | sends struct of supported items for download.                                 |
| POST          | /dynamic                                | takes custom list of software from user that generates correct setupt script. |
| GET           | /scripts/:uuid                          | takes in uuid and sends user the file to the install via CL.                  |

