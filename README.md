
  # Gator

  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

  ## Description

  An RSS feed aggregator, built using Go.

  ## Table of Contents
  * [Installation](#installation)
  * [Usage](#usage)
  * [License](#license)
  * [Contribution](#contribution)
  * [Tests](#tests)
  * [Questions](#questions)
  
  ## Installation

  You will need Postgres and Go installed to build. Run go install github.com/bmkersey/Go-Gator@latest or to build locally clone the repo and from the root run go build

  ## Usage

Run the application with one of the following commands:

```bash
go run main.go <command> ##if you built
gator ##if you installed
```

### Available Commands

| Command      | Description                                    |
|--------------|------------------------------------------------|
| `login`      | Log in to your account(requires username arg)  |
| `register`   | Register a new user (requires username arg)    |
| `reset`      | Reset entire database (use caution)            |
| `users`      | View a list of all users                       |
| `agg`        | Aggregate and display data (takes duration arg)|
| `addfeed`    | Add a new RSS feed (requires login)            |
| `feeds`      | Show available feeds                           |
| `follow`     | Follow a feed (requires login)                 |
| `following`  | List followed feeds (requires login)           |
| `unfollow`   | Unfollow a feed (requires login)               |
| `browse`     | Browse all recent posts (requires login)       |

  ## License

  This project uses the The MIT License.  
  Please visit [https://opensource.org/licenses/MIT](https://opensource.org/licenses/MIT) to learn more.
  

  ## Contribution

  Sole contributor
  
  ## Tests 

  No testing package
  
  ## Questions
  Questions? Comments? Concerns? Feel free to reach out!  
  Email: bmkersey@gmail.com  
  GitHub: [bmkersey](https://github.com/bmkersey)  
  
  