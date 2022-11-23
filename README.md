# Udacity Golang CRM Backend project.

Implementation in Golang of the CRM backend project following the requirements detailed in [Proyect Rubrik](https://review.udacity.com/#!/rubrics/4856/view "Proyect Rubrik")

The project is implemented in the main.go file
The dependencies are the following:

*	"encoding/json"
*	"fmt"
*	"github.com/gorilla/mux"
*	"log"
*	"net/http"
*	"time"
  
Install the indicated modules and run the program using the command line go run main.go

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/01_20220125_%202022-11-23_14.42.22.png)

Every time the project starts, it generates five predefined customers.

The endpoints have been created following the requirements, and are the following:

Getting all customers through a the /customers path

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/02_20220125_%202022-11-23_14.45.11.png)

Getting a single customer through a /customers/{id} path

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/03_20220125_%202022-11-23_14.45.28.png)

Creating a customer through a /customers path

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/04_20220125_%202022-11-23_14.47.40.png)

Define in Postman a request body with all the fields except the Id field.

For the creation of a customer, the generation of the id that it must have has been automated.
A timestamp created using the time module will be used as id

Updating a customer through a /customers/{id} path

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/05_20220125_%202022-11-23_14.48.01.png)

Define in Postman a request body with all the fields except the Id field.

Deleting a customer through a /customers/{id} path

![picture alt](https://github.com/erfederuiz/udagoback/blob/main/images/06_20220125_%202022-11-23_14.48.28.png)













