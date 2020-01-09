# Endpoint Documentation
| Verb          | Route                                   |
| ------------- |:---------------------------------------:|
| BASE URL      | quicky.dev/api                          |



| Verb          | Route                                   | Quick Description                                                                        |
| ------------- |:---------------------------------------:| --------------------------------:                                                        |
| GET           | < OPERATING_SYSTEM >/generic                         | setup script sent as response.                                              |
| GET           | < OPERATING_SYSTEM >/availableItems                  | sends struct of supported items for download.                               |
| POST          | < OPERATING_SYSTEM >/dynamic                         | takes custom list of software from user that generates correct setuptscript.|
| GET           | < OPERATING_SYSTEM >/scripts/:uuid                   | takes in uuid and sends user the file to the install via CL.                |

## Example request

`https://www.quicky.dev/macos/generic`

## Example response

