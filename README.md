# VetHeal

## Animal Disease Diagnostic Support for Owners

### Akdeniz University - Institute of Sciences - Computer Science Engineering Department - Senior Project Design

Live version can be found in: https://www.vetheal.app for frontend.

Usually, an animal signals its owner when his health is at risk or he has a disease, exhibits unusual behavior or changes his body. However, owners who do not have in-depth knowledge of the pet's disease will be able to enter these symptoms into the application intended to be developed in this project, to see what the ailments in their pets are and their risk ratios.

The system works primarily by selecting an animal to determine which type of pet to diagnose. Next, the system provides symptom categories. The system will be able to ask the user whether there are other symptoms associated with these symptoms or questions about these symptoms. At the end of the questions, the system presents the priority of this problem, an overview of this problem, symptom assessments, and recommendations to the user.

In this study, after the data was normalized and all database normalization applications were made, the back-end development was done using the test-driven development method. React framework, which is very popular, is used in the front-end. In this research, We’ve lived through the TDD cycle. We followed the test-code-refactor (or red-green-green) cycle all the way for the back-end.

## 1. Introduction

### 1.1. Problem of Research

In recent days, dogs, cats, rabbits etc. are not only animals that live around humans, but almost family members that living in the same house, have constant contact with their owners. There are pet stores and veterinarians to take care of this new type of non-human family members. However, people connected to their pets may experience emotional losses due to the disease or sudden death of their pets, and there is also an increased risk of transmission of microorganisms between humans and pets.

Usually, an animal gives a signal to its owner when his health is at risk or he has a disease, exhibiting unusual behavior or changing his body. However, owners who do not have in-depth knowledge of the disease of the pet tend to neglect such symptoms. They rely only on regular checks of their pets by veterinarians.

### 1.2. Specific Aims

The health of these pets, which makes our lives better, is quite important. Animals with any disease cannot tell us the symptoms of this disease by speaking like humans. But these animals often reveal the presence of these diseases through changes in their body and behavior, as well as their disease or discomfort. For this reason, pet owners should regularly observe changes in their pets and consult a veterinarian when they notice these changes.
Many animal owners underestimate the veterinary control of their pets, which should be carried out routinely, and postpone it as long as they can. Sometimes even wait until their pet is sick, and then a check-up is carried out. Pet owners often put on hold in this way, as they think their pets become stressed and scared, anxious and aggressive when they go to the veterinary.

At the same time, rising veterinary costs may prevent some pet owners from accepting the necessary medical care. Thanks to this application, we can prevent both the negative impact of pet owners financially and psychologically, and we can find out the condition of the pet, health, the degree of risk of this disease if it is sick and whether veterinary intervention is required, without affecting the psychology of our pets.

Sometimes we cannot reach veterinarians all the time and quickly. At the same time, some pet owners may be reluctant to consult the veterinarian in smaller problems, as they cannot find veterinary expenses and time. At the same time, pet owners find more general and negative responses when they search the internet for symptoms in their pets. This causes concern to pet owners.

The motivation for this research is to develop a animal health pre- diagnosis/monitoring system with deterministic method that can be easily controlled by animal owners without having deep knowledge of computer technology or the diseases of their pets. The system recommended in this article requires symptoms that the pet owner finds only from their pets, after which the system responds to some possible diseases that pets may have with a calculated rate of trust and the risk priority of the disease.

## 2. Material & Method

During the implementation of this study, the data set was carefully prepared and different software development techniques were tried. At the same time, software languages and libraries that have become very popular lately and are very efficient in terms of performance have been used.

### 2.1. Test Driven Development

The back-end part of the application has been developed with test driven development approach.Test Driven Development (TDD) is a software development approach in which test cases are developed to determine and validate what the code will do. In simple terms, test cases are created and tested for each functionality first, and if the test fails, new code is written to pass the test and make the code simple and bug-free.

### 2.2. Data

We have collected data from vets4pets, which was in raw format with nested JSON tree for Questions and Results data. Question data format had response field that represents choice option for the question, and those response object had data format called Actionable that represents either the next question or the response. Animal data was in list form that contains name, id and animal type. Symptom data was also in list format that contains id, animal id as reference, name, and question id as reference.

### 2.3. System & Database Design

After normalizing the data and transforming it into the appropriate structure, the database tables and structure are adjusted according to the 3NF normalization rules.

## 3. Languages, Frameworks & Tools

### 3.1. Front-End

The frontend, developed in the NodeJS environment, ensures readability and sustainability of the source code with ReactJS technology, increasing the speed of user access to the application with the Code Splitting method provided by the Webpack module bundler, and user-friendly and easy use of the application with the Material Design principle.

### 3.2. Back-End

In this study, modern frameworks and programming languages are used. Golang, a very popular and new language, was preferred for the back-end. Go, also known as Golang, is an open source, compiled and statically typed programming language designed by Google. It is designed to be simple, high-performance, readable and efficient.

This project has 3 layers :

- Repository Layer
- Service Layer
- Controller Layer

#### Project Tree

### Project Tree

```
├── controller/
│   ├── animal_controller_test.go
│   └── animal_controller.go
├── service/
│   ├── animal_service_test.go
│   └── animal_service.go
├── repository/
│   ├── animal_repository_test.go
│   └── animal_repository.go
├── mock/
│   ├── mock_animalrepository.go
│   └── mock_animalservice.go
├── server/
│   └── server.go
├── main.go
├── go.mod/
└── ...
```

### Installation

#### Install Dependencies

```
go mod download
```

### Build

```
go build -o /backend
```

### Run project

```
go run .
```

### Run unit tests

```
go test ./...
```

### Run unit tests & create coverage HTML file

```
go test -coverprofile cover.out and then
go tool cover -html=cover.out
```

## Dependencies

All dependencies can be found on package.json file. Also you can check the list:

- [Go](https://github.com/golang/go)
- [Go Mock](github.com/golang/mock v1.6.0)
- [Pact Go](github.com/pact-foundation/pact-go)
- [Testify](github.com/stretchr/testify)

### 3.3. Database

Postgres was preferred because of it’s open-source, allows you to store large and sophisticated data safely. It helps developers to build the most complex applications, run administrative tasks and create integral environments.

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Nida Dinç - niddinc@gmail.com
