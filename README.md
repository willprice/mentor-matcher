# mentor-matcher
A website for mentees to find mentors and mentors to offer their wisdom to others

https://mentor-matcher.fly.dev/

## Local Installation

This project uses GNU Make to automate the setup process. Ensure that you have a `.env` file in the 
root directory of the project. See `.env.example` for an example of the required environment 
variables.

To build the binaries, run the following command:

```bash
make
```

Alternatively, you can use Docker compose to run the project locally. You'll still need the `.env` 
file. To do so, run the following command:

```bash
make docker-up
```

To stop the Docker containers, run the following command:

```bash
make docker-down
```

