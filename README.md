# Burns Anxiety Inventory Console Application

## Overview

The **Burns Anxiety Inventory Console Application** is a simple Golang program designed to assess and record anxiety symptoms in individuals. It allows users to read and respond to a series of questions, indicating on a scale from 0 to 3 how much they have experienced each symptom over the past week. The scale is as follows:

- 0: Not at all
- 1: A little
- 2: Partially
- 3: Very much

The collected responses are then saved into a PostgreSQL database for later analysis and tracking of anxiety levels.

## Requirements

Before you can run this application, make sure you have the following prerequisites installed on your system:

- **Go (Golang)**: You can download and install Go from the official website: https://golang.org/dl/

- **PostgreSQL**: Install and set up PostgreSQL on your system. You'll need to create a database for this application and configure the connection details in the code.

## Usage

1. Clone this repository to your local machine:
   ```
   git clone https://github.com/your-username/burns-anxiety-inventory.git
   ```

2. Navigate to the project directory:
   ```
   cd burns-anxiety-inventory
   ```

3. Configure the PostgreSQL database connection in the `main.go` file.

4. Build and run the application:
   ```
   go run main.go
   ```

5. Follow the prompts to answer each question in the inventory. Use the scale (0 to 3) to rate your anxiety level for each symptom.

6. Once you've completed the inventory, the results will be saved into the PostgreSQL database for later analysis.

## Code Structure

- `main.go`: The main application file. It contains the code to interact with the user, collect responses, and store them in the PostgreSQL database.

- `WorryingFeelings/worryingFeelings.go`: Asks questions about the level of worrying feelings of the user and receives their answers.

- `WorryingThoughts/worryingThoughts.go`: Asks questions about the level of worrying thoughts of the user and receives their answers.

- `CheckInput/checkInput.go`: Imports the necessary libraries for reading user input and converting it to an integer.

- `SymptomsOfPhysicalDiscomfort/symptomsOfPhysicalDiscomfort.go`: Imports the necessary libraries for printing text to the console, getting user input, and converting strings to integers.

- `PrintLastResults/printLastResults.go`: Imports the necessary libraries for connecting to a PostgreSQL database, querying data, and formatting dates.


## Contributing

Contributions are welcome! If you have ideas for improvements or new features, please submit issues or pull requests.

## License

This project is open-source and available under the GPL-3.0 license. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

---

Assess and track your anxiety levels easily with the Burns Anxiety Inventory Console Application. Your responses are securely stored in a PostgreSQL database for your future reference and analysis. Enjoy using this simple yet effective tool! 😌📊
